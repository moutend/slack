package api

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/moutend/slack/internal/models"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Client represents API client.
type Client struct {
	botToken  string
	userToken string

	bot  *slack.Client
	user *slack.Client

	debug *log.Logger

	KeepFetchingMessages FetchMessagesConditionFunc
	KeepFetchingReplies  FetchMessagesConditionFunc
}

// New returns prepared API client.
func New(botToken, userToken string) *Client {
	return &Client{
		bot:   slack.New(botToken),
		user:  slack.New(userToken),
		debug: log.New(ioutil.Discard, "debug: ", 0),
	}
}

// SetLogger sets logger for debugging.
func (c *Client) SetLogger(w io.Writer) {
	if w != nil {
		c.debug = log.New(w, "debug: ", 0)

		c.debug.Println("SetLogger: start debugging")

		return
	}

	c.debug = log.New(ioutil.Discard, "debug:", 0)
}

// Whoami returns login user identities.
func (a *Client) Whoami() (botName, botID, userName, userID string, err error) {
	var (
		a1, a2 *slack.AuthTestResponse
		u1, u2 *url.URL
	)

	a1, err = a.bot.AuthTest()

	if err != nil {
		err = fmt.Errorf("client: failed to authorize: %w", err)

		return
	}

	a2, err = a.user.AuthTest()

	if err != nil {
		err = fmt.Errorf("client: failed to authorize: %w", err)

		return
	}

	u1, err = url.Parse(a1.URL)

	if err != nil {
		err = fmt.Errorf("client: failed to parse URL: %w", err)

		return
	}

	u2, err = url.Parse(a2.URL)

	if err != nil {
		err = fmt.Errorf("client: failed to parse URL: %w", err)

		return
	}
	if u1.Host != u2.Host {
		err = fmt.Errorf("client: bot token or user token is invalid (bot=%q,user=%q)", u1.Host, u2.Host)

		return
	}

	botName = a1.User
	botID = a1.UserID

	userName = a2.User
	userID = a2.UserID

	return
}

func upsertUser(ctx context.Context, tx boil.ContextTransactor, user slack.User) error {
	u, err := models.FindUser(ctx, tx, user.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if u == nil {
		u = &models.User{}
	}

	u.ID = user.ID
	u.TeamID = user.TeamID
	u.Name = user.Name
	u.Deleted = user.Deleted
	u.Color = user.Color
	u.RealName = user.RealName

	if err == sql.ErrNoRows {
		err = u.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = u.Update(ctx, tx, boil.Infer())
	}

	return err
}

// FetchUsers fetches and saves information about users.
func (a *Client) FetchUsers(ctx context.Context, tx boil.ContextTransactor) error {
	defer a.debug.Println("FetchUsers: done")

	users, err := a.bot.GetUsersContext(ctx)

	if err != nil {
		return errors.Wrap(err, "failed to fetch users")
	}
	for _, user := range users {
		if err := upsertUser(ctx, tx, user); err != nil {
			return errors.Wrap(err, "api: failed to save user")
		}
	}

	return nil
}

func upsertChannel(ctx context.Context, tx boil.ContextTransactor, channel slack.Channel) error {
	c, err := models.FindChannel(ctx, tx, channel.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		c = &models.Channel{}
	}

	c.ID = channel.ID
	c.IsOpen = channel.IsOpen
	c.LastRead = channel.LastRead
	c.UnreadCount = int64(channel.UnreadCount)
	c.UnreadCountDisplay = int64(channel.UnreadCountDisplay)
	c.IsGroup = channel.IsGroup
	c.IsShared = channel.IsShared
	c.IsIm = channel.IsIM
	c.IsExtShared = channel.IsExtShared
	c.IsOrgShared = channel.IsOrgShared
	c.IsPendingExtShared = channel.IsPendingExtShared
	c.IsPrivate = channel.IsPrivate
	c.IsMpim = channel.IsMpIM
	c.Unlinked = int64(channel.Unlinked)
	c.NameNormalized = channel.NameNormalized
	c.NumMembers = int64(channel.NumMembers)
	c.Creator = channel.Creator
	c.IsArchived = channel.IsArchived
	c.Topic = channel.Topic.Value
	c.Purpose = channel.Purpose.Value
	c.IsChannel = channel.IsChannel
	c.IsGeneral = channel.IsGeneral
	c.IsMember = channel.IsMember
	c.Locale = channel.Locale
	c.Name = channel.Name
	c.User = channel.User

	if err == sql.ErrNoRows {
		err = c.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = c.Update(ctx, tx, boil.Infer())
	}

	return err
}

// fetchChannels fetches and saves information about channels.
func (a *Client) FetchChannels(ctx context.Context, tx boil.ContextTransactor) error {
	defer a.debug.Println("FetchChannels: done")

	parameter := &slack.GetConversationsParameters{
		Limit: 100,
		Types: []string{"public_channel", "private_channel", "mpim", "im"},
	}

	for {
		a.debug.Printf("FetchChannels: parameter: %+v\n", parameter)

		channels, nextCursor, err := a.user.GetConversationsContext(ctx, parameter)

		if err != nil {
			return errors.Wrap(err, "api: failed to fetch channels")
		}
		for _, channel := range channels {
			if err := upsertChannel(ctx, tx, channel); err != nil {
				return errors.Wrap(err, "api: failed to save channel")
			}
		}

		parameter.Cursor = nextCursor

		if parameter.Cursor == "" {
			break
		}

		a.debug.Println("FetchChannels: sleep 100 ms because of rate/ limit")
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func upsertMessage(ctx context.Context, tx boil.ContextTransactor, message slack.Message) error {
	m, err := models.FindMessage(ctx, tx, message.Timestamp)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		m = &models.Message{}
	}
	if s := strings.Split(message.Timestamp, "."); len(s) < 2 {
		return fmt.Errorf("timestamp is broken")
	} else {
		sec, err := strconv.Atoi(s[0])

		if err != nil {
			return err
		}

		nano, err := strconv.Atoi(s[1])

		if err != nil {
			return err
		}

		m.CreatedAt = time.Unix(int64(sec), int64(nano))
	}

	m.Type = message.Type
	m.Channel = message.Channel
	m.User = message.User
	m.Timestamp = message.Timestamp
	m.ThreadTimestamp = message.ThreadTimestamp
	m.Text = message.Text
	m.SubType = message.SubType
	m.BotID = message.BotID
	m.Name = message.Name
	m.ReplyCount = int64(message.ReplyCount)

	if err == sql.ErrNoRows {
		err = m.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = m.Update(ctx, tx, boil.Infer())
	}

	return err
}

// FetchMessagesConditionFunc is a function which indicates continuation. Fetching is performed while this function returns true.
type FetchMessagesConditionFunc func(fetchedMessageCount int, messages []slack.Message) (keepFetching bool)

// FetchMessages fetches and saves information about messages.
func (a *Client) FetchMessages(ctx context.Context, tx boil.ContextTransactor, channelID string) error {
	defer a.debug.Println("FetchMessages: done")

	conversationCount := 0
	replyCount := 0
	timestamps := []string{}

	parameter := &slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Inclusive: false,
		Latest:    "",
		Limit:     100,
		Oldest:    "",
	}

	for {
		a.debug.Printf("FetchMessages: fetch messages: cursor: %v\n", parameter.Cursor)

		result, err := a.user.GetConversationHistoryContext(ctx, parameter)

		if err != nil {
			return fmt.Errorf("api: failed to fetch messages: %w", err)
		}
		for _, message := range result.Messages {
			// I don't know why, but the Message.Channel field seems to be always empty.
			message.Channel = channelID

			if message.ReplyCount > 0 {
				timestamps = append(timestamps, message.Timestamp)
			}
			if err := upsertMessage(ctx, tx, message); err != nil {
				return fmt.Errorf("api: failed to save message: %s: %w", message.ClientMsgID, err)
			}
		}

		conversationCount += len(result.Messages)

		if a.KeepFetchingMessages != nil && !a.KeepFetchingMessages(conversationCount, result.Messages) {
			a.debug.Println("FetchMessages: stop fetching conversation messages")

			break
		}

		parameter.Cursor = result.ResponseMetaData.NextCursor

		if parameter.Cursor == "" || !result.HasMore {
			a.debug.Println("FetchMessages: finish fetching all messages")

			break
		}

		a.debug.Println("FetchMessages: sleep 100 ms because rate / limit")
		time.Sleep(100 * time.Millisecond)
	}
	for i, timestamp := range timestamps {
		parameter := &slack.GetConversationRepliesParameters{
			ChannelID: channelID,
			Timestamp: timestamp,
			Inclusive: false,
			Latest:    "",
			Limit:     100,
			Oldest:    "",
		}

		for {
			a.debug.Printf("FetchMessages: fetch replies (%d/%d): cursor: %v\n", i+1, len(timestamps), parameter.Cursor)

			messages, hasMore, nextCursor, err := a.user.GetConversationRepliesContext(ctx, parameter)

			if err != nil {
				return fmt.Errorf("api: failed to fetch replies: %w")
			}
			for _, message := range messages {
				// I don't know why, but the Message.Channel field seems to be always empty.
				message.Channel = channelID

				if err := upsertMessage(ctx, tx, message); err != nil {
					return fmt.Errorf("api: failed to save reply: %s: %w", message.ClientMsgID, err)
				}
			}

			replyCount += len(messages)

			if a.KeepFetchingReplies != nil && !a.KeepFetchingReplies(replyCount, messages) {
				a.debug.Println("FetchMessages: stop fetching replies")

				return nil
			}

			parameter.Cursor = nextCursor

			if parameter.Cursor == "" || !hasMore {
				a.debug.Printf("FetchMessages: finish fetching all replies (%d/%d)\n", i+1, len(timestamps))

				break
			}

			a.debug.Println("FetchMessages: sleep 100 ms because rate / limit")
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}
