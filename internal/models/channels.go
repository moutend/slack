// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Channel is an object representing the database table.
type Channel struct {
	ID                 string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	IsOpen             bool      `boil:"is_open" json:"is_open" toml:"is_open" yaml:"is_open"`
	LastRead           string    `boil:"last_read" json:"last_read" toml:"last_read" yaml:"last_read"`
	UnreadCount        int64     `boil:"unread_count" json:"unread_count" toml:"unread_count" yaml:"unread_count"`
	UnreadCountDisplay int64     `boil:"unread_count_display" json:"unread_count_display" toml:"unread_count_display" yaml:"unread_count_display"`
	IsGroup            bool      `boil:"is_group" json:"is_group" toml:"is_group" yaml:"is_group"`
	IsShared           bool      `boil:"is_shared" json:"is_shared" toml:"is_shared" yaml:"is_shared"`
	IsIm               bool      `boil:"is_im" json:"is_im" toml:"is_im" yaml:"is_im"`
	IsExtShared        bool      `boil:"is_ext_shared" json:"is_ext_shared" toml:"is_ext_shared" yaml:"is_ext_shared"`
	IsOrgShared        bool      `boil:"is_org_shared" json:"is_org_shared" toml:"is_org_shared" yaml:"is_org_shared"`
	IsPendingExtShared bool      `boil:"is_pending_ext_shared" json:"is_pending_ext_shared" toml:"is_pending_ext_shared" yaml:"is_pending_ext_shared"`
	IsPrivate          bool      `boil:"is_private" json:"is_private" toml:"is_private" yaml:"is_private"`
	IsMpim             bool      `boil:"is_mpim" json:"is_mpim" toml:"is_mpim" yaml:"is_mpim"`
	Unlinked           int64     `boil:"unlinked" json:"unlinked" toml:"unlinked" yaml:"unlinked"`
	NameNormalized     string    `boil:"name_normalized" json:"name_normalized" toml:"name_normalized" yaml:"name_normalized"`
	NumMembers         int64     `boil:"num_members" json:"num_members" toml:"num_members" yaml:"num_members"`
	User               string    `boil:"user" json:"user" toml:"user" yaml:"user"`
	Name               string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Creator            string    `boil:"creator" json:"creator" toml:"creator" yaml:"creator"`
	IsArchived         bool      `boil:"is_archived" json:"is_archived" toml:"is_archived" yaml:"is_archived"`
	Topic              string    `boil:"topic" json:"topic" toml:"topic" yaml:"topic"`
	Purpose            string    `boil:"purpose" json:"purpose" toml:"purpose" yaml:"purpose"`
	IsChannel          bool      `boil:"is_channel" json:"is_channel" toml:"is_channel" yaml:"is_channel"`
	IsGeneral          bool      `boil:"is_general" json:"is_general" toml:"is_general" yaml:"is_general"`
	IsMember           bool      `boil:"is_member" json:"is_member" toml:"is_member" yaml:"is_member"`
	Locale             string    `boil:"locale" json:"locale" toml:"locale" yaml:"locale"`
	CreatedAt          time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt          time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *channelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L channelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChannelColumns = struct {
	ID                 string
	IsOpen             string
	LastRead           string
	UnreadCount        string
	UnreadCountDisplay string
	IsGroup            string
	IsShared           string
	IsIm               string
	IsExtShared        string
	IsOrgShared        string
	IsPendingExtShared string
	IsPrivate          string
	IsMpim             string
	Unlinked           string
	NameNormalized     string
	NumMembers         string
	User               string
	Name               string
	Creator            string
	IsArchived         string
	Topic              string
	Purpose            string
	IsChannel          string
	IsGeneral          string
	IsMember           string
	Locale             string
	CreatedAt          string
	UpdatedAt          string
}{
	ID:                 "id",
	IsOpen:             "is_open",
	LastRead:           "last_read",
	UnreadCount:        "unread_count",
	UnreadCountDisplay: "unread_count_display",
	IsGroup:            "is_group",
	IsShared:           "is_shared",
	IsIm:               "is_im",
	IsExtShared:        "is_ext_shared",
	IsOrgShared:        "is_org_shared",
	IsPendingExtShared: "is_pending_ext_shared",
	IsPrivate:          "is_private",
	IsMpim:             "is_mpim",
	Unlinked:           "unlinked",
	NameNormalized:     "name_normalized",
	NumMembers:         "num_members",
	User:               "user",
	Name:               "name",
	Creator:            "creator",
	IsArchived:         "is_archived",
	Topic:              "topic",
	Purpose:            "purpose",
	IsChannel:          "is_channel",
	IsGeneral:          "is_general",
	IsMember:           "is_member",
	Locale:             "locale",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ChannelWhere = struct {
	ID                 whereHelperstring
	IsOpen             whereHelperbool
	LastRead           whereHelperstring
	UnreadCount        whereHelperint64
	UnreadCountDisplay whereHelperint64
	IsGroup            whereHelperbool
	IsShared           whereHelperbool
	IsIm               whereHelperbool
	IsExtShared        whereHelperbool
	IsOrgShared        whereHelperbool
	IsPendingExtShared whereHelperbool
	IsPrivate          whereHelperbool
	IsMpim             whereHelperbool
	Unlinked           whereHelperint64
	NameNormalized     whereHelperstring
	NumMembers         whereHelperint64
	User               whereHelperstring
	Name               whereHelperstring
	Creator            whereHelperstring
	IsArchived         whereHelperbool
	Topic              whereHelperstring
	Purpose            whereHelperstring
	IsChannel          whereHelperbool
	IsGeneral          whereHelperbool
	IsMember           whereHelperbool
	Locale             whereHelperstring
	CreatedAt          whereHelpertime_Time
	UpdatedAt          whereHelpertime_Time
}{
	ID:                 whereHelperstring{field: "\"channels\".\"id\""},
	IsOpen:             whereHelperbool{field: "\"channels\".\"is_open\""},
	LastRead:           whereHelperstring{field: "\"channels\".\"last_read\""},
	UnreadCount:        whereHelperint64{field: "\"channels\".\"unread_count\""},
	UnreadCountDisplay: whereHelperint64{field: "\"channels\".\"unread_count_display\""},
	IsGroup:            whereHelperbool{field: "\"channels\".\"is_group\""},
	IsShared:           whereHelperbool{field: "\"channels\".\"is_shared\""},
	IsIm:               whereHelperbool{field: "\"channels\".\"is_im\""},
	IsExtShared:        whereHelperbool{field: "\"channels\".\"is_ext_shared\""},
	IsOrgShared:        whereHelperbool{field: "\"channels\".\"is_org_shared\""},
	IsPendingExtShared: whereHelperbool{field: "\"channels\".\"is_pending_ext_shared\""},
	IsPrivate:          whereHelperbool{field: "\"channels\".\"is_private\""},
	IsMpim:             whereHelperbool{field: "\"channels\".\"is_mpim\""},
	Unlinked:           whereHelperint64{field: "\"channels\".\"unlinked\""},
	NameNormalized:     whereHelperstring{field: "\"channels\".\"name_normalized\""},
	NumMembers:         whereHelperint64{field: "\"channels\".\"num_members\""},
	User:               whereHelperstring{field: "\"channels\".\"user\""},
	Name:               whereHelperstring{field: "\"channels\".\"name\""},
	Creator:            whereHelperstring{field: "\"channels\".\"creator\""},
	IsArchived:         whereHelperbool{field: "\"channels\".\"is_archived\""},
	Topic:              whereHelperstring{field: "\"channels\".\"topic\""},
	Purpose:            whereHelperstring{field: "\"channels\".\"purpose\""},
	IsChannel:          whereHelperbool{field: "\"channels\".\"is_channel\""},
	IsGeneral:          whereHelperbool{field: "\"channels\".\"is_general\""},
	IsMember:           whereHelperbool{field: "\"channels\".\"is_member\""},
	Locale:             whereHelperstring{field: "\"channels\".\"locale\""},
	CreatedAt:          whereHelpertime_Time{field: "\"channels\".\"created_at\""},
	UpdatedAt:          whereHelpertime_Time{field: "\"channels\".\"updated_at\""},
}

// ChannelRels is where relationship names are stored.
var ChannelRels = struct {
}{}

// channelR is where relationships are stored.
type channelR struct {
}

// NewStruct creates a new relationship struct
func (*channelR) NewStruct() *channelR {
	return &channelR{}
}

// channelL is where Load methods for each relationship are stored.
type channelL struct{}

var (
	channelAllColumns            = []string{"id", "is_open", "last_read", "unread_count", "unread_count_display", "is_group", "is_shared", "is_im", "is_ext_shared", "is_org_shared", "is_pending_ext_shared", "is_private", "is_mpim", "unlinked", "name_normalized", "num_members", "user", "name", "creator", "is_archived", "topic", "purpose", "is_channel", "is_general", "is_member", "locale", "created_at", "updated_at"}
	channelColumnsWithoutDefault = []string{"id", "is_open", "last_read", "unread_count", "unread_count_display", "is_group", "is_shared", "is_im", "is_ext_shared", "is_org_shared", "is_pending_ext_shared", "is_private", "is_mpim", "unlinked", "name_normalized", "num_members", "user", "name", "creator", "is_archived", "topic", "purpose", "is_channel", "is_general", "is_member", "locale", "created_at", "updated_at"}
	channelColumnsWithDefault    = []string{}
	channelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChannelSlice is an alias for a slice of pointers to Channel.
	// This should generally be used opposed to []Channel.
	ChannelSlice []*Channel
	// ChannelHook is the signature for custom Channel hook methods
	ChannelHook func(context.Context, boil.ContextExecutor, *Channel) error

	channelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	channelType                 = reflect.TypeOf(&Channel{})
	channelMapping              = queries.MakeStructMapping(channelType)
	channelPrimaryKeyMapping, _ = queries.BindMapping(channelType, channelMapping, channelPrimaryKeyColumns)
	channelInsertCacheMut       sync.RWMutex
	channelInsertCache          = make(map[string]insertCache)
	channelUpdateCacheMut       sync.RWMutex
	channelUpdateCache          = make(map[string]updateCache)
	channelUpsertCacheMut       sync.RWMutex
	channelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var channelBeforeInsertHooks []ChannelHook
var channelBeforeUpdateHooks []ChannelHook
var channelBeforeDeleteHooks []ChannelHook
var channelBeforeUpsertHooks []ChannelHook

var channelAfterInsertHooks []ChannelHook
var channelAfterSelectHooks []ChannelHook
var channelAfterUpdateHooks []ChannelHook
var channelAfterDeleteHooks []ChannelHook
var channelAfterUpsertHooks []ChannelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Channel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Channel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Channel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Channel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Channel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Channel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Channel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Channel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Channel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChannelHook registers your hook function for all future operations.
func AddChannelHook(hookPoint boil.HookPoint, channelHook ChannelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		channelBeforeInsertHooks = append(channelBeforeInsertHooks, channelHook)
	case boil.BeforeUpdateHook:
		channelBeforeUpdateHooks = append(channelBeforeUpdateHooks, channelHook)
	case boil.BeforeDeleteHook:
		channelBeforeDeleteHooks = append(channelBeforeDeleteHooks, channelHook)
	case boil.BeforeUpsertHook:
		channelBeforeUpsertHooks = append(channelBeforeUpsertHooks, channelHook)
	case boil.AfterInsertHook:
		channelAfterInsertHooks = append(channelAfterInsertHooks, channelHook)
	case boil.AfterSelectHook:
		channelAfterSelectHooks = append(channelAfterSelectHooks, channelHook)
	case boil.AfterUpdateHook:
		channelAfterUpdateHooks = append(channelAfterUpdateHooks, channelHook)
	case boil.AfterDeleteHook:
		channelAfterDeleteHooks = append(channelAfterDeleteHooks, channelHook)
	case boil.AfterUpsertHook:
		channelAfterUpsertHooks = append(channelAfterUpsertHooks, channelHook)
	}
}

// One returns a single channel record from the query.
func (q channelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Channel, error) {
	o := &Channel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for channels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Channel records from the query.
func (q channelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChannelSlice, error) {
	var o []*Channel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Channel slice")
	}

	if len(channelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Channel records in the query.
func (q channelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q channelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if channels exists")
	}

	return count > 0, nil
}

// Channels retrieves all the records using an executor.
func Channels(mods ...qm.QueryMod) channelQuery {
	mods = append(mods, qm.From("\"channels\""))
	return channelQuery{NewQuery(mods...)}
}

// FindChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChannel(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Channel, error) {
	channelObj := &Channel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"channels\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, channelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from channels")
	}

	return channelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Channel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no channels provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(channelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	channelInsertCacheMut.RLock()
	cache, cached := channelInsertCache[key]
	channelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			channelAllColumns,
			channelColumnsWithDefault,
			channelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(channelType, channelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(channelType, channelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"channels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"channels\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"channels\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, channelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into channels")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for channels")
	}

CacheNoHooks:
	if !cached {
		channelInsertCacheMut.Lock()
		channelInsertCache[key] = cache
		channelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Channel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Channel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	channelUpdateCacheMut.RLock()
	cache, cached := channelUpdateCache[key]
	channelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			channelAllColumns,
			channelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, channelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(channelType, channelMapping, append(wl, channelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update channels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for channels")
	}

	if !cached {
		channelUpdateCacheMut.Lock()
		channelUpdateCache[key] = cache
		channelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q channelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for channels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChannelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in channel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all channel")
	}
	return rowsAff, nil
}

// Delete deletes a single Channel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Channel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Channel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), channelPrimaryKeyMapping)
	sql := "DELETE FROM \"channels\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for channels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q channelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no channelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for channels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChannelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(channelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from channel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for channels")
	}

	if len(channelAfterDeleteHooks) != 0 {
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
func (o *Channel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChannel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChannelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"channels\".* FROM \"channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChannelSlice")
	}

	*o = slice

	return nil
}

// ChannelExists checks if the Channel row exists.
func ChannelExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"channels\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if channels exists")
	}

	return exists, nil
}
