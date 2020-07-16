package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/slack-go/slack"
)

// APIClient is wrapper of slack.Client.
type APIClient struct {
	botToken  string
	userToken string

	bot  *slack.Client
	user *slack.Client

	channelsCacheDuration      time.Duration
	conversationsCacheDuration time.Duration
	replyCacheDuration         time.Duration
	usersCacheDuration         time.Duration

	rootCachePath         string
	infoCachePath         string
	conversationCachePath string
	replyCachePath        string

	debug *log.Logger
}

// Whoami returns login user identities.
func (a *APIClient) Whoami() (botName, botID, userName, userID string, err error) {
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

// SetLogger sets logger for debugging.
func (a *APIClient) SetLogger(l *log.Logger) {
	if a == nil {
		return
	}
	if l == nil {
		return
	}

	a.debug = l
}

// GetAllUsersContext retrieves all users.
func (a *APIClient) GetAllUsersContext(ctx context.Context) ([]slack.User, error) {
	if a == nil {
		return nil, nil
	}

	usersCachePath := filepath.Join(a.infoCachePath, "users", "info.json")
	os.MkdirAll(filepath.Dir(usersCachePath), 0755)

	var cache struct {
		ExpiredAt time.Time    `json:"expired_at"`
		Users     []slack.User `json:"channels"`
	}

	if data, err := ioutil.ReadFile(usersCachePath); err == nil {
		a.debug.Printf("GetAllUsersContext: found %s", usersCachePath)

		if err := json.Unmarshal(data, &cache); err != nil {
			return nil, fmt.Errorf("client: failed to parse cache: %w", err)
		}
	}
	if !cache.ExpiredAt.IsZero() && cache.ExpiredAt.After(time.Now().UTC()) {
		a.debug.Println("GetAllUsersContext: use cache")

		return cache.Users, nil
	}

	users, err := a.bot.GetUsersContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("client: failed to get users: %w", err)
	}

	cache.ExpiredAt = time.Now().UTC().Add(a.usersCacheDuration)
	cache.Users = users

	data, err := json.MarshalIndent(cache, "", "  ")

	if err != nil {
		return nil, fmt.Errorf("client: failed to build cache: %w", err)
	}
	if err := ioutil.WriteFile(usersCachePath, data, 0644); err != nil {
		return nil, fmt.Errorf("client: failed to write cache: %w", err)
	}

	return users, nil
}

// GetAllChannelsContext retrieves all channels.
func (a *APIClient) GetAllChannelsContext(ctx context.Context) ([]slack.Channel, error) {
	if a == nil {
		return nil, nil
	}

	channelsCachePath := filepath.Join(a.infoCachePath, "channels", "info.json")
	os.MkdirAll(filepath.Dir(channelsCachePath), 0755)

	var cache struct {
		ExpiredAt time.Time       `json:"expired_at"`
		Channels  []slack.Channel `json:"channels"`
	}

	if data, err := ioutil.ReadFile(channelsCachePath); err == nil {
		a.debug.Printf("GetAllChannelsContext: found %s", channelsCachePath)

		if err := json.Unmarshal(data, &cache); err != nil {
			return nil, err
		}
	}
	if !cache.ExpiredAt.IsZero() && cache.ExpiredAt.After(time.Now().UTC()) {
		a.debug.Println("GetAllChannelsContext: use cache")

		return cache.Channels, nil
	}

	a.debug.Println("GetAllChannelsContext: fetch channels")

	const maxTry = 100

	channels := []slack.Channel{}
	parameter := &slack.GetConversationsParameters{
		Limit: 100,
		Types: []string{"public_channel", "private_channel", "mpim", "im"},
	}

	for i := 0; i < maxTry; i++ {
		cs, nextCursor, err := a.user.GetConversationsContext(ctx, parameter)

		if err != nil {
			a.debug.Printf("GetAllChannelsContext: failed to get channels: %w", err)

			break
		}

		channels = append(channels, cs...)
		parameter.Cursor = nextCursor

		if parameter.Cursor == "" {
			break
		}

		a.debug.Println("GetAllChannelsContext: sleep 50 ms because of rate/ limit")
		time.Sleep(50 * time.Millisecond)
	}

	cache.ExpiredAt = time.Now().UTC().Add(a.channelsCacheDuration)
	cache.Channels = channels

	data, err := json.MarshalIndent(cache, "", "  ")

	if err != nil {
		return channels, fmt.Errorf("client: failed to build cache: %w", err)
	}
	if err := ioutil.WriteFile(channelsCachePath, data, 0644); err != nil {
		return nil, fmt.Errorf("client: failed to write cache: %w", err)
	}

	return channels, nil
}

type replyInfo struct {
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *APIClient) loadReplyInfo(channelID string) (replyInfo, error) {
	var info replyInfo

	if data, err := ioutil.ReadFile(filepath.Join(a.infoCachePath, "reply", channelID+".json")); err == nil {
		if err := json.Unmarshal(data, &info); err != nil {
			return info, err
		}
	}

	return info, nil
}

func (a *APIClient) saveReplyInfo(channelID string, info replyInfo) error {
	data, err := json.MarshalIndent(info, "", "  ")

	if err != nil {
		return err
	}

	path := filepath.Join(a.infoCachePath, "reply", channelID+".json")
	os.MkdirAll(filepath.Dir(path), 0755)

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

type replyCache map[string]struct {
	Messages map[string]slack.Message `json:"messages"`
}

func (a *APIClient) loadReplyCache(channelID string) (replyCache, error) {
	var cache replyCache

	if data, err := ioutil.ReadFile(filepath.Join(a.replyCachePath, channelID+".json")); err == nil {
		if err := json.Unmarshal(data, &cache); err != nil {
			return nil, fmt.Errorf("client: failed to parse cache: %w", err)
		}
	}
	if cache == nil {
		cache = map[string]struct {
			Messages map[string]slack.Message `json:"messages"`
		}{}
	}

	return cache, nil
}

func (a *APIClient) saveReplyCache(channelID string, cache replyCache) error {
	data, err := json.MarshalIndent(cache, "", "  ")

	if err != nil {
		return err
	}

	path := filepath.Join(a.replyCachePath, channelID+".json")
	os.MkdirAll(filepath.Dir(path), 0755)

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

// GetAllRepliesContext retrieves all messages associated with the channel.
func (a *APIClient) GetAllRepliesContext(ctx context.Context, channelID, timestamp string, forceFetch bool) ([]slack.Message, error) {
	info, err := a.loadReplyInfo(channelID)

	if err != nil {
		return nil, err
	}

	cache, err := a.loadReplyCache(channelID)

	if err != nil {
		return nil, err
	}
	if !forceFetch && !info.UpdatedAt.IsZero() && time.Now().UTC().Sub(info.UpdatedAt) < a.replyCacheDuration {
		a.debug.Println("GetAllRepliesContext: skip fetch")

		messages := []slack.Message{}

		for _, m := range cache[timestamp].Messages {
			messages = append(messages, m)
		}

		return messages, nil
	}
	if _, ok := cache[timestamp]; !ok {
		cache[timestamp] = struct {
			Messages map[string]slack.Message `json:"messages"`
		}{
			Messages: map[string]slack.Message{},
		}
	}

	a.debug.Println("GetAllRepliesContext: fetch replies")

	defer a.saveReplyCache(channelID, cache)
	defer func() {
		a.saveReplyInfo(channelID, info)
	}()

	parameter := &slack.GetConversationRepliesParameters{
		ChannelID: channelID,
		Timestamp: timestamp,
		Inclusive: false,
		Latest:    "",
		Limit:     100,
		Oldest:    "",
	}
	errorCount := 0

	for {
		msgs, hasMore, nextCursor, err := a.user.GetConversationRepliesContext(ctx, parameter)

		if err != nil {
			errorCount++

			if errorCount < 10 {
				a.debug.Println("GetAllRepliesContext: sleep because error:", err)
				time.Sleep(25 * time.Millisecond)
				continue
			} else {
				return nil, fmt.Errorf("client: failed to get messages: %w", err)
			}
		}
		for _, m := range msgs {
			cache[timestamp].Messages[m.Timestamp] = m
		}

		parameter.Cursor = nextCursor

		if parameter.Cursor == "" || !hasMore {
			break
		}
		if !info.UpdatedAt.IsZero() {
			break
		}

		a.debug.Println("GetAllRepliesContext: sleep 50 ms because of rate/ limit")
		time.Sleep(50 * time.Millisecond)
	}

	info.UpdatedAt = time.Now().UTC()

	messages := []slack.Message{}

	for _, m := range cache[timestamp].Messages {
		messages = append(messages, m)
	}

	return messages, nil
}

type conversationInfo struct {
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *APIClient) loadConversationInfo(channelID string) (conversationInfo, error) {
	var info conversationInfo

	if data, err := ioutil.ReadFile(filepath.Join(a.infoCachePath, "conversation", channelID+".json")); err == nil {
		if err := json.Unmarshal(data, &info); err != nil {
			return info, err
		}
	}

	return info, nil
}

func (a *APIClient) saveConversationInfo(channelID string, info conversationInfo) error {
	data, err := json.MarshalIndent(info, "", "  ")

	if err != nil {
		return err
	}

	path := filepath.Join(a.infoCachePath, "conversation", channelID+".json")
	os.MkdirAll(filepath.Dir(path), 0755)

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

type conversationCache map[string]slack.Message

func (a *APIClient) loadConversationCache(channelID string) (conversationCache, error) {
	var cache conversationCache

	if data, err := ioutil.ReadFile(filepath.Join(a.conversationCachePath, channelID+".json")); err == nil {
		if err := json.Unmarshal(data, &cache); err != nil {
			return cache, err
		}
	}
	if cache == nil {
		cache = conversationCache{}
	}

	return cache, nil
}

func (a *APIClient) saveConversationCache(channelID string, cache conversationCache) error {
	data, err := json.MarshalIndent(cache, "", "  ")

	if err != nil {
		return err
	}

	path := filepath.Join(a.conversationCachePath, channelID+".json")
	os.MkdirAll(filepath.Dir(path), 0755)

	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		return err
	}

	return nil
}

// GetAllConversationsContext retrieves all messages associated with the channel.
func (a *APIClient) GetAllConversationsContext(ctx context.Context, channelID string, forceFetch bool) ([]slack.Message, error) {
	info, err := a.loadConversationInfo(channelID)

	if err != nil {
		return nil, err
	}

	cache, err := a.loadConversationCache(channelID)

	if err != nil {
		return nil, err
	}
	if !forceFetch && !info.UpdatedAt.IsZero() && time.Now().UTC().Sub(info.UpdatedAt) < a.conversationsCacheDuration {
		a.debug.Println("GetAllConversationsContext: use cache")

		rc, err := a.loadReplyCache(channelID)

		if err != nil {
			return nil, err
		}

		messages := []slack.Message{}

		for _, m := range cache {
			messages = append(messages, m)
		}
		for _, ms := range rc {
			for _, m := range ms.Messages {
				messages = append(messages, m)
			}
		}

		return messages, nil
	}

	a.debug.Println("GetAllConversationsContext: fetch conversations", channelID)

	defer a.saveConversationCache(channelID, cache)
	defer func() {
		a.saveConversationInfo(channelID, info)
	}()

	timestamps := []string{}
	parameter := &slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Inclusive: false,
		Latest:    "",
		Limit:     100,
		Oldest:    "",
	}
	errorCount := 0

	for {
		result, err := a.user.GetConversationHistoryContext(ctx, parameter)

		if err != nil {
			errorCount++

			if errorCount < 10 {
				a.debug.Println("GetAllConversationsContext: sleep because error:", err)
				time.Sleep(25 * time.Millisecond)
				continue
			} else {
				return nil, fmt.Errorf("client: failed to get messages: %w", err)
			}
		}
		for _, m := range result.Messages {
			cache[m.Timestamp] = m

			if m.ReplyCount > 0 {
				timestamps = append(timestamps, m.Timestamp)
			}
		}

		parameter.Cursor = result.ResponseMetaData.NextCursor

		if parameter.Cursor == "" || !result.HasMore {
			break
		}
		if !info.UpdatedAt.IsZero() {
			break
		}

		a.debug.Println("GetAllConversationsContext: sleep 50 ms because of rate / limit")
		time.Sleep(50 * time.Millisecond)
	}

	messages := []slack.Message{}

	for _, timestamp := range timestamps {
		replies, err := a.GetAllRepliesContext(ctx, channelID, timestamp, info.UpdatedAt.IsZero())

		if err != nil {
			return nil, err
		}

		messages = append(messages, replies...)
	}
	for _, m := range cache {
		messages = append(messages, m)
	}

	info.UpdatedAt = time.Now().UTC()

	return messages, nil
}

// New returns prepared API client.
func New(botToken, userToken, cachePath string) *APIClient {
	rootCachePath := cachePath
	os.MkdirAll(rootCachePath, 0755)

	infoCachePath := filepath.Join(rootCachePath, "info")
	os.MkdirAll(infoCachePath, 0755)

	conversationCachePath := filepath.Join(rootCachePath, "conversation")
	os.MkdirAll(conversationCachePath, 0755)

	replyCachePath := filepath.Join(rootCachePath, "reply")
	os.MkdirAll(replyCachePath, 0755)

	return &APIClient{
		bot:                        slack.New(botToken),
		user:                       slack.New(userToken),
		debug:                      log.New(ioutil.Discard, "debug: ", 0),
		channelsCacheDuration:      24 * time.Hour,
		conversationsCacheDuration: 100 * time.Second,
		replyCacheDuration:         100 * time.Second,
		usersCacheDuration:         24 * time.Hour,
		rootCachePath:              rootCachePath,
		infoCachePath:              infoCachePath,
		conversationCachePath:      conversationCachePath,
		replyCachePath:             replyCachePath,
	}
}
