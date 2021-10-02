package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/moutend/slack/internal/models"
	"github.com/slack-go/slack"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	discard = log.New(io.Discard, "debug: ", 0)
)

// FetchMessagesConditionFunc is a function which indicates continuation. Fetching is performed while this function returns true.
type FetchMessagesConditionFunc func(fetchedMessageCount int, messages []slack.Message) (keepFetching bool)

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
		debug: discard,
	}
}

// SetLogger sets logger for debugging.
func (c *Client) SetLogger(w io.Writer) {
	if w != nil {
		c.debug = log.New(w, "debug: ", 0)
	}

	c.debug.Println("SetLogger: start debugging")
}

// Whoami returns login user identities.
func (c *Client) Whoami() (botName, botID, userName, userID string, err error) {
	var (
		a1, a2 *slack.AuthTestResponse
		u1, u2 *url.URL
	)

	a1, err = c.bot.AuthTest()

	if err != nil {
		err = fmt.Errorf("api: failed to authorize: %w", err)

		return
	}

	a2, err = c.user.AuthTest()

	if err != nil {
		err = fmt.Errorf("api: failed to authorize: %w", err)

		return
	}

	u1, err = url.Parse(a1.URL)

	if err != nil {
		err = fmt.Errorf("api: failed to parse URL: %w", err)

		return
	}

	u2, err = url.Parse(a2.URL)

	if err != nil {
		err = fmt.Errorf("api: failed to parse URL: %w", err)

		return
	}
	if u1.Host != u2.Host {
		err = fmt.Errorf("api: bot token or user token is invalid (bot=%q,user=%q)", u1.Host, u2.Host)

		return
	}

	botName = a1.User
	botID = a1.UserID

	userName = a2.User
	userID = a2.UserID

	return
}

// UpsertUser updates or inserts given user into database.
func (c *Client) UpsertUser(ctx context.Context, tx boil.ContextTransactor, user slack.User) error {
	if c.debug != discard {
		data, err := json.Marshal(user)

		if err != nil {
			return err
		}

		c.debug.Printf("UpsertUser: raw data: %q\n", string(data))
	}

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
	u.TZ = user.TZ
	u.TZLabel = user.TZLabel
	u.TZOffset = int64(user.TZOffset)
	// u.Profile
	u.IsBot = user.IsBot
	u.IsAdmin = user.IsAdmin
	u.IsOwner = user.IsOwner
	u.IsPrimaryOwner = user.IsPrimaryOwner
	u.IsRestricted = user.IsRestricted
	u.IsUltraRestricted = user.IsUltraRestricted
	u.IsStranger = user.IsStranger
	u.IsAppUser = user.IsAppUser
	u.IsInvitedUser = user.IsInvitedUser
	u.Has2fa = user.Has2FA
	u.HasFiles = user.HasFiles
	u.Presence = user.Presence
	u.Locale = user.Locale
	// user.Updated
	// user.Enterprise

	if err == sql.ErrNoRows {
		err = u.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = u.Update(ctx, tx, boil.Infer())
	}

	return err
}

// FetchUsers fetches and saves users.
func (c *Client) FetchUsers(ctx context.Context, tx boil.ContextTransactor) error {
	defer c.debug.Println("FetchUsers: done")

	users, err := c.bot.GetUsersContext(ctx)

	if err != nil {
		return fmt.Errorf("api: failed to fetch users: %w", err)
	}
	for _, user := range users {
		if err := c.UpsertUser(ctx, tx, user); err != nil {
			return fmt.Errorf("api: failed to save user: %w", err)
		}
	}

	return nil
}

// UpsertChannel updates or inserts given channel into database.
func (c *Client) UpsertChannel(ctx context.Context, tx boil.ContextTransactor, channel slack.Channel) error {
	if c.debug != discard {
		data, err := json.Marshal(channel)

		if err != nil {
			return err
		}

		c.debug.Printf("UpsertChannel: raw data: %q\n", string(data))
	}

	ch, err := models.FindChannel(ctx, tx, channel.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		ch = &models.Channel{}
	}

	ch.ID = channel.ID
	ch.IsOpen = channel.IsOpen
	ch.LastRead = channel.LastRead
	ch.UnreadCount = int64(channel.UnreadCount)
	ch.UnreadCountDisplay = int64(channel.UnreadCountDisplay)
	ch.IsGroup = channel.IsGroup
	ch.IsShared = channel.IsShared
	ch.IsIm = channel.IsIM
	ch.IsExtShared = channel.IsExtShared
	ch.IsOrgShared = channel.IsOrgShared
	ch.IsPendingExtShared = channel.IsPendingExtShared
	ch.IsPrivate = channel.IsPrivate
	ch.IsMpim = channel.IsMpIM
	ch.Unlinked = int64(channel.Unlinked)
	ch.NameNormalized = channel.NameNormalized
	ch.NumMembers = int64(channel.NumMembers)
	ch.Creator = channel.Creator
	ch.IsArchived = channel.IsArchived
	ch.Topic = channel.Topic.Value
	ch.Purpose = channel.Purpose.Value
	ch.IsChannel = channel.IsChannel
	ch.IsGeneral = channel.IsGeneral
	ch.IsMember = channel.IsMember
	ch.Locale = channel.Locale
	ch.Name = channel.Name
	ch.User = channel.User

	if err == sql.ErrNoRows {
		err = ch.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = ch.Update(ctx, tx, boil.Infer())
	}

	return err
}

// fetchChannels fetches and saves channels.
func (c *Client) FetchChannels(ctx context.Context, tx boil.ContextTransactor) error {
	defer c.debug.Println("FetchChannels: done")

	parameter := &slack.GetConversationsParameters{
		Limit: 100,
		Types: []string{"public_channel", "private_channel", "mpim", "im"},
	}

	for {
		c.debug.Printf("FetchChannels: cursor: %v\n", parameter.Cursor)

		channels, nextCursor, err := c.user.GetConversationsContext(ctx, parameter)

		if err != nil {
			return fmt.Errorf("api: failed to fetch channels: %w", err)
		}
		for _, channel := range channels {
			if err := c.UpsertChannel(ctx, tx, channel); err != nil {
				return fmt.Errorf("api: failed to save channel: %w", err)
			}
		}

		parameter.Cursor = nextCursor

		if parameter.Cursor == "" {
			break
		}

		c.debug.Println("FetchChannels: sleep 100 ms because of rate limit")
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

// UpsertMessage updates or inserts given message into database.
func (c *Client) UpsertMessage(ctx context.Context, tx boil.ContextTransactor, message slack.Message) error {
	if c.debug != discard {
		data, err := json.Marshal(message)

		if err != nil {
			return err
		}

		c.debug.Printf("UpsertMessage: raw data: %q\n", string(data))
	}

	m, err := models.FindMessage(ctx, tx, message.Timestamp)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		m = &models.Message{}
	}
	if s := strings.Split(message.Timestamp, "."); len(s) < 2 {
		return fmt.Errorf("timestamp is broken: %s", message.Timestamp)
	} else {
		sec, err := strconv.Atoi(s[0])

		if err != nil {
			return fmt.Errorf("failed to parse seconds part of timestamp: %s: %w", message.Timestamp, err)
		}

		nano, err := strconv.Atoi(s[1])

		if err != nil {
			return fmt.Errorf("failed to parse nano seconds part of timestamp: %s: %w", message.Timestamp, err)
		}

		m.CreatedAt = time.Unix(int64(sec), int64(nano)).UTC()
	}

	m.ClientMSGID = message.ClientMsgID
	m.Type = message.Type
	m.Channel = message.Channel
	m.User = message.User
	m.Text = message.Text
	m.Timestamp = message.Timestamp
	m.ThreadTimestamp = message.ThreadTimestamp
	m.IsStarred = message.IsStarred
	m.LastRead = message.LastRead
	m.Subscribed = message.Subscribed
	m.UnreadCount = int64(message.UnreadCount)
	m.SubType = message.SubType
	m.Hidden = message.Hidden
	m.DeletedTimestamp = message.DeletedTimestamp
	m.EventTimestamp = message.EventTimestamp
	m.BotID = message.BotID
	m.Username = message.Username
	m.Inviter = message.Inviter
	m.Topic = message.Topic
	m.Purpose = message.Purpose
	m.Name = message.Name
	m.OldName = message.OldName
	m.ReplyCount = int64(message.ReplyCount)
	m.ParentUserID = message.ParentUserId
	m.Upload = message.Upload
	m.ItemType = message.ItemType
	m.ReplyTo = int64(message.ReplyTo)
	m.Team = message.Team
	m.ResponseType = message.ResponseType
	m.ReplaceOriginal = message.ReplaceOriginal
	m.DeleteOriginal = message.DeleteOriginal

	if err == sql.ErrNoRows {
		err = m.Insert(ctx, tx, boil.Infer())
	} else {
		_, err = m.Update(ctx, tx, boil.Infer())
	}

	return err
}

// FetchMessages fetches and saves messages.
func (c *Client) FetchMessages(ctx context.Context, tx boil.ContextTransactor, channelID string) error {
	defer c.debug.Println("FetchMessages: done")

	conversationCount := 0
	replyCount := 0
	timestamps := []string{}
	retry := 25

	parameter := &slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Inclusive: false,
		Latest:    "",
		Limit:     100,
		Oldest:    "",
	}

	for {
		c.debug.Printf("FetchMessages: fetch messages: cursor: %v\n", parameter.Cursor)

		result, err := c.user.GetConversationHistoryContext(ctx, parameter)

		if err != nil && strings.Contains(err.Error(), "slack rate limit exceeded") {
			c.debug.Printf("FetchMessages: detect API rate limit: %w", err)

			if retry == 0 {
				return nil
			}

			c.debug.Printf("FetchMessages: sleep because reached API limitation: remaining retry count: %d", retry)
			time.Sleep(2 * time.Second)

			retry -= 1

			continue
		}
		if err != nil {
			return fmt.Errorf("api: failed to fetch messages: %w", err)
		}
		for _, message := range result.Messages {
			// I don't know why, but the Message.Channel field seems to be always empty.
			message.Channel = channelID

			if message.ReplyCount > 0 {
				timestamps = append(timestamps, message.Timestamp)
			}
			if err := c.UpsertMessage(ctx, tx, message); err != nil {
				return fmt.Errorf("api: failed to save message: %s: %w", message.ClientMsgID, err)
			}
		}

		conversationCount += len(result.Messages)

		if c.KeepFetchingMessages != nil && !c.KeepFetchingMessages(conversationCount, result.Messages) {
			c.debug.Println("FetchMessages: stop fetching conversation messages")

			break
		}

		parameter.Cursor = result.ResponseMetaData.NextCursor

		if parameter.Cursor == "" || !result.HasMore {
			c.debug.Println("FetchMessages: finish fetching all messages")

			break
		}

		c.debug.Println("FetchMessages: sleep 100 ms because rate limit")
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
			c.debug.Printf("FetchMessages: fetch replies (%d/%d): cursor: %v\n", i+1, len(timestamps), parameter.Cursor)

			messages, hasMore, nextCursor, err := c.user.GetConversationRepliesContext(ctx, parameter)

			if err != nil && strings.Contains(err.Error(), "slack rate limit exceeded") {
				c.debug.Printf("FetchMessages: detect API rate limit: %w", err)

				if retry == 0 {
					return nil
				}

				c.debug.Printf("FetchMessages: sleep because reached API limitation: remaining retry count: %d", retry)
				time.Sleep(2 * time.Second)

				retry -= 1

				continue
			}
			if err != nil {
				return fmt.Errorf("api: failed to fetch replies: %w", err)
			}
			for _, message := range messages {
				// I don't know why, but the Message.Channel field seems to be always empty.
				message.Channel = channelID

				if err := c.UpsertMessage(ctx, tx, message); err != nil {
					return fmt.Errorf("api: failed to save reply: %s: %w", message.ClientMsgID, err)
				}
			}

			replyCount += len(messages)

			if c.KeepFetchingReplies != nil && !c.KeepFetchingReplies(replyCount, messages) {
				c.debug.Println("FetchMessages: stop fetching replies")

				return nil
			}

			parameter.Cursor = nextCursor

			if parameter.Cursor == "" || !hasMore {
				c.debug.Printf("FetchMessages: finish fetching all replies (%d/%d)\n", i+1, len(timestamps))

				return nil
			}

			c.debug.Println("FetchMessages: sleep 100 ms because rate limit")
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}
