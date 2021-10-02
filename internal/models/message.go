// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Message is an object representing the database table.
type Message struct {
	ClientMSGID      string    `boil:"client_msg_id" json:"client_msg_id" toml:"client_msg_id" yaml:"client_msg_id"`
	Type             string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	Channel          string    `boil:"channel" json:"channel" toml:"channel" yaml:"channel"`
	User             string    `boil:"user" json:"user" toml:"user" yaml:"user"`
	Text             string    `boil:"text" json:"text" toml:"text" yaml:"text"`
	Timestamp        string    `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	ThreadTimestamp  string    `boil:"thread_timestamp" json:"thread_timestamp" toml:"thread_timestamp" yaml:"thread_timestamp"`
	IsStarred        bool      `boil:"is_starred" json:"is_starred" toml:"is_starred" yaml:"is_starred"`
	LastRead         string    `boil:"last_read" json:"last_read" toml:"last_read" yaml:"last_read"`
	Subscribed       bool      `boil:"subscribed" json:"subscribed" toml:"subscribed" yaml:"subscribed"`
	UnreadCount      int64     `boil:"unread_count" json:"unread_count" toml:"unread_count" yaml:"unread_count"`
	SubType          string    `boil:"sub_type" json:"sub_type" toml:"sub_type" yaml:"sub_type"`
	Hidden           bool      `boil:"hidden" json:"hidden" toml:"hidden" yaml:"hidden"`
	DeletedTimestamp string    `boil:"deleted_timestamp" json:"deleted_timestamp" toml:"deleted_timestamp" yaml:"deleted_timestamp"`
	EventTimestamp   string    `boil:"event_timestamp" json:"event_timestamp" toml:"event_timestamp" yaml:"event_timestamp"`
	BotID            string    `boil:"bot_id" json:"bot_id" toml:"bot_id" yaml:"bot_id"`
	UserName         string    `boil:"user_name" json:"user_name" toml:"user_name" yaml:"user_name"`
	Inviter          string    `boil:"inviter" json:"inviter" toml:"inviter" yaml:"inviter"`
	Topic            string    `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Purpose          string    `boil:"purpose" json:"purpose" toml:"purpose" yaml:"purpose"`
	Name             string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	OldName          string    `boil:"old_name" json:"old_name" toml:"old_name" yaml:"old_name"`
	ReplyCount       int64     `boil:"reply_count" json:"reply_count" toml:"reply_count" yaml:"reply_count"`
	ParentUserID     string    `boil:"parent_user_id" json:"parent_user_id" toml:"parent_user_id" yaml:"parent_user_id"`
	Upload           bool      `boil:"upload" json:"upload" toml:"upload" yaml:"upload"`
	ItemType         string    `boil:"item_type" json:"item_type" toml:"item_type" yaml:"item_type"`
	ReplyTo          int64     `boil:"reply_to" json:"reply_to" toml:"reply_to" yaml:"reply_to"`
	Team             string    `boil:"team" json:"team" toml:"team" yaml:"team"`
	ResponseType     string    `boil:"response_type" json:"response_type" toml:"response_type" yaml:"response_type"`
	ReplaceOriginal  bool      `boil:"replace_original" json:"replace_original" toml:"replace_original" yaml:"replace_original"`
	DeleteOriginal   bool      `boil:"delete_original" json:"delete_original" toml:"delete_original" yaml:"delete_original"`
	CreatedAt        time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *messageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageColumns = struct {
	ClientMSGID      string
	Type             string
	Channel          string
	User             string
	Text             string
	Timestamp        string
	ThreadTimestamp  string
	IsStarred        string
	LastRead         string
	Subscribed       string
	UnreadCount      string
	SubType          string
	Hidden           string
	DeletedTimestamp string
	EventTimestamp   string
	BotID            string
	UserName         string
	Inviter          string
	Topic            string
	Purpose          string
	Name             string
	OldName          string
	ReplyCount       string
	ParentUserID     string
	Upload           string
	ItemType         string
	ReplyTo          string
	Team             string
	ResponseType     string
	ReplaceOriginal  string
	DeleteOriginal   string
	CreatedAt        string
}{
	ClientMSGID:      "client_msg_id",
	Type:             "type",
	Channel:          "channel",
	User:             "user",
	Text:             "text",
	Timestamp:        "timestamp",
	ThreadTimestamp:  "thread_timestamp",
	IsStarred:        "is_starred",
	LastRead:         "last_read",
	Subscribed:       "subscribed",
	UnreadCount:      "unread_count",
	SubType:          "sub_type",
	Hidden:           "hidden",
	DeletedTimestamp: "deleted_timestamp",
	EventTimestamp:   "event_timestamp",
	BotID:            "bot_id",
	UserName:         "user_name",
	Inviter:          "inviter",
	Topic:            "topic",
	Purpose:          "purpose",
	Name:             "name",
	OldName:          "old_name",
	ReplyCount:       "reply_count",
	ParentUserID:     "parent_user_id",
	Upload:           "upload",
	ItemType:         "item_type",
	ReplyTo:          "reply_to",
	Team:             "team",
	ResponseType:     "response_type",
	ReplaceOriginal:  "replace_original",
	DeleteOriginal:   "delete_original",
	CreatedAt:        "created_at",
}

// Generated where

var MessageWhere = struct {
	ClientMSGID      whereHelperstring
	Type             whereHelperstring
	Channel          whereHelperstring
	User             whereHelperstring
	Text             whereHelperstring
	Timestamp        whereHelperstring
	ThreadTimestamp  whereHelperstring
	IsStarred        whereHelperbool
	LastRead         whereHelperstring
	Subscribed       whereHelperbool
	UnreadCount      whereHelperint64
	SubType          whereHelperstring
	Hidden           whereHelperbool
	DeletedTimestamp whereHelperstring
	EventTimestamp   whereHelperstring
	BotID            whereHelperstring
	UserName         whereHelperstring
	Inviter          whereHelperstring
	Topic            whereHelperstring
	Purpose          whereHelperstring
	Name             whereHelperstring
	OldName          whereHelperstring
	ReplyCount       whereHelperint64
	ParentUserID     whereHelperstring
	Upload           whereHelperbool
	ItemType         whereHelperstring
	ReplyTo          whereHelperint64
	Team             whereHelperstring
	ResponseType     whereHelperstring
	ReplaceOriginal  whereHelperbool
	DeleteOriginal   whereHelperbool
	CreatedAt        whereHelpertime_Time
}{
	ClientMSGID:      whereHelperstring{field: "\"message\".\"client_msg_id\""},
	Type:             whereHelperstring{field: "\"message\".\"type\""},
	Channel:          whereHelperstring{field: "\"message\".\"channel\""},
	User:             whereHelperstring{field: "\"message\".\"user\""},
	Text:             whereHelperstring{field: "\"message\".\"text\""},
	Timestamp:        whereHelperstring{field: "\"message\".\"timestamp\""},
	ThreadTimestamp:  whereHelperstring{field: "\"message\".\"thread_timestamp\""},
	IsStarred:        whereHelperbool{field: "\"message\".\"is_starred\""},
	LastRead:         whereHelperstring{field: "\"message\".\"last_read\""},
	Subscribed:       whereHelperbool{field: "\"message\".\"subscribed\""},
	UnreadCount:      whereHelperint64{field: "\"message\".\"unread_count\""},
	SubType:          whereHelperstring{field: "\"message\".\"sub_type\""},
	Hidden:           whereHelperbool{field: "\"message\".\"hidden\""},
	DeletedTimestamp: whereHelperstring{field: "\"message\".\"deleted_timestamp\""},
	EventTimestamp:   whereHelperstring{field: "\"message\".\"event_timestamp\""},
	BotID:            whereHelperstring{field: "\"message\".\"bot_id\""},
	UserName:         whereHelperstring{field: "\"message\".\"user_name\""},
	Inviter:          whereHelperstring{field: "\"message\".\"inviter\""},
	Topic:            whereHelperstring{field: "\"message\".\"topic\""},
	Purpose:          whereHelperstring{field: "\"message\".\"purpose\""},
	Name:             whereHelperstring{field: "\"message\".\"name\""},
	OldName:          whereHelperstring{field: "\"message\".\"old_name\""},
	ReplyCount:       whereHelperint64{field: "\"message\".\"reply_count\""},
	ParentUserID:     whereHelperstring{field: "\"message\".\"parent_user_id\""},
	Upload:           whereHelperbool{field: "\"message\".\"upload\""},
	ItemType:         whereHelperstring{field: "\"message\".\"item_type\""},
	ReplyTo:          whereHelperint64{field: "\"message\".\"reply_to\""},
	Team:             whereHelperstring{field: "\"message\".\"team\""},
	ResponseType:     whereHelperstring{field: "\"message\".\"response_type\""},
	ReplaceOriginal:  whereHelperbool{field: "\"message\".\"replace_original\""},
	DeleteOriginal:   whereHelperbool{field: "\"message\".\"delete_original\""},
	CreatedAt:        whereHelpertime_Time{field: "\"message\".\"created_at\""},
}

// MessageRels is where relationship names are stored.
var MessageRels = struct {
}{}

// messageR is where relationships are stored.
type messageR struct {
}

// NewStruct creates a new relationship struct
func (*messageR) NewStruct() *messageR {
	return &messageR{}
}

// messageL is where Load methods for each relationship are stored.
type messageL struct{}

var (
	messageAllColumns            = []string{"client_msg_id", "type", "channel", "user", "text", "timestamp", "thread_timestamp", "is_starred", "last_read", "subscribed", "unread_count", "sub_type", "hidden", "deleted_timestamp", "event_timestamp", "bot_id", "user_name", "inviter", "topic", "purpose", "name", "old_name", "reply_count", "parent_user_id", "upload", "item_type", "reply_to", "team", "response_type", "replace_original", "delete_original", "created_at"}
	messageColumnsWithoutDefault = []string{"client_msg_id", "type", "channel", "user", "text", "timestamp", "thread_timestamp", "is_starred", "last_read", "subscribed", "unread_count", "sub_type", "hidden", "deleted_timestamp", "event_timestamp", "bot_id", "user_name", "inviter", "topic", "purpose", "name", "old_name", "reply_count", "parent_user_id", "upload", "item_type", "reply_to", "team", "response_type", "replace_original", "delete_original", "created_at"}
	messageColumnsWithDefault    = []string{}
	messagePrimaryKeyColumns     = []string{"timestamp"}
)

type (
	// MessageSlice is an alias for a slice of pointers to Message.
	// This should generally be used opposed to []Message.
	MessageSlice []*Message
	// MessageHook is the signature for custom Message hook methods
	MessageHook func(context.Context, boil.ContextExecutor, *Message) error

	messageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageType                 = reflect.TypeOf(&Message{})
	messageMapping              = queries.MakeStructMapping(messageType)
	messagePrimaryKeyMapping, _ = queries.BindMapping(messageType, messageMapping, messagePrimaryKeyColumns)
	messageInsertCacheMut       sync.RWMutex
	messageInsertCache          = make(map[string]insertCache)
	messageUpdateCacheMut       sync.RWMutex
	messageUpdateCache          = make(map[string]updateCache)
	messageUpsertCacheMut       sync.RWMutex
	messageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var messageBeforeInsertHooks []MessageHook
var messageBeforeUpdateHooks []MessageHook
var messageBeforeDeleteHooks []MessageHook
var messageBeforeUpsertHooks []MessageHook

var messageAfterInsertHooks []MessageHook
var messageAfterSelectHooks []MessageHook
var messageAfterUpdateHooks []MessageHook
var messageAfterDeleteHooks []MessageHook
var messageAfterUpsertHooks []MessageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Message) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Message) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Message) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Message) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Message) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Message) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Message) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Message) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Message) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMessageHook registers your hook function for all future operations.
func AddMessageHook(hookPoint boil.HookPoint, messageHook MessageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		messageBeforeInsertHooks = append(messageBeforeInsertHooks, messageHook)
	case boil.BeforeUpdateHook:
		messageBeforeUpdateHooks = append(messageBeforeUpdateHooks, messageHook)
	case boil.BeforeDeleteHook:
		messageBeforeDeleteHooks = append(messageBeforeDeleteHooks, messageHook)
	case boil.BeforeUpsertHook:
		messageBeforeUpsertHooks = append(messageBeforeUpsertHooks, messageHook)
	case boil.AfterInsertHook:
		messageAfterInsertHooks = append(messageAfterInsertHooks, messageHook)
	case boil.AfterSelectHook:
		messageAfterSelectHooks = append(messageAfterSelectHooks, messageHook)
	case boil.AfterUpdateHook:
		messageAfterUpdateHooks = append(messageAfterUpdateHooks, messageHook)
	case boil.AfterDeleteHook:
		messageAfterDeleteHooks = append(messageAfterDeleteHooks, messageHook)
	case boil.AfterUpsertHook:
		messageAfterUpsertHooks = append(messageAfterUpsertHooks, messageHook)
	}
}

// One returns a single message record from the query.
func (q messageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Message, error) {
	o := &Message{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for message")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Message records from the query.
func (q messageQuery) All(ctx context.Context, exec boil.ContextExecutor) (MessageSlice, error) {
	var o []*Message

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Message slice")
	}

	if len(messageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Message records in the query.
func (q messageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count message rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q messageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if message exists")
	}

	return count > 0, nil
}

// Messages retrieves all the records using an executor.
func Messages(mods ...qm.QueryMod) messageQuery {
	mods = append(mods, qm.From("\"message\""))
	return messageQuery{NewQuery(mods...)}
}

// FindMessage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessage(ctx context.Context, exec boil.ContextExecutor, timestamp string, selectCols ...string) (*Message, error) {
	messageObj := &Message{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"message\" where \"timestamp\"=?", sel,
	)

	q := queries.Raw(query, timestamp)

	err := q.Bind(ctx, exec, messageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from message")
	}

	return messageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Message) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no message provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(messageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageInsertCacheMut.RLock()
	cache, cached := messageInsertCache[key]
	messageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageAllColumns,
			messageColumnsWithDefault,
			messageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageType, messageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"message\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"message\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"message\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, messagePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into message")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Timestamp,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for message")
	}

CacheNoHooks:
	if !cached {
		messageInsertCacheMut.Lock()
		messageInsertCache[key] = cache
		messageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Message.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Message) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	messageUpdateCacheMut.RLock()
	cache, cached := messageUpdateCache[key]
	messageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update message, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"message\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, messagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageType, messageMapping, append(wl, messagePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update message row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for message")
	}

	if !cached {
		messageUpdateCacheMut.Lock()
		messageUpdateCache[key] = cache
		messageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q messageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for message")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for message")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"message\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, messagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all message")
	}
	return rowsAff, nil
}

// Delete deletes a single Message record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Message) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Message provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messagePrimaryKeyMapping)
	sql := "DELETE FROM \"message\" WHERE \"timestamp\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from message")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for message")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(messageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"message\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, messagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message")
	}

	if len(messageAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Message) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessage(ctx, exec, o.Timestamp)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"message\".* FROM \"message\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, messagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MessageSlice")
	}

	*o = slice

	return nil
}

// MessageExists checks if the Message row exists.
func MessageExists(ctx context.Context, exec boil.ContextExecutor, timestamp string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"message\" where \"timestamp\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, timestamp)
	}
	row := exec.QueryRowContext(ctx, sql, timestamp)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if message exists")
	}

	return exists, nil
}
