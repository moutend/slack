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

// User is an object representing the database table.
type User struct {
	ID                string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	TeamID            string    `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	Name              string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Deleted           bool      `boil:"deleted" json:"deleted" toml:"deleted" yaml:"deleted"`
	Color             string    `boil:"color" json:"color" toml:"color" yaml:"color"`
	RealName          string    `boil:"real_name" json:"real_name" toml:"real_name" yaml:"real_name"`
	TZ                string    `boil:"tz" json:"tz" toml:"tz" yaml:"tz"`
	TZLabel           string    `boil:"tz_label" json:"tz_label" toml:"tz_label" yaml:"tz_label"`
	TZOffset          int64     `boil:"tz_offset" json:"tz_offset" toml:"tz_offset" yaml:"tz_offset"`
	IsBot             bool      `boil:"is_bot" json:"is_bot" toml:"is_bot" yaml:"is_bot"`
	IsAdmin           bool      `boil:"is_admin" json:"is_admin" toml:"is_admin" yaml:"is_admin"`
	IsOwner           bool      `boil:"is_owner" json:"is_owner" toml:"is_owner" yaml:"is_owner"`
	IsPrimaryOwner    bool      `boil:"is_primary_owner" json:"is_primary_owner" toml:"is_primary_owner" yaml:"is_primary_owner"`
	IsRestricted      bool      `boil:"is_restricted" json:"is_restricted" toml:"is_restricted" yaml:"is_restricted"`
	IsUltraRestricted bool      `boil:"is_ultra_restricted" json:"is_ultra_restricted" toml:"is_ultra_restricted" yaml:"is_ultra_restricted"`
	IsStranger        bool      `boil:"is_stranger" json:"is_stranger" toml:"is_stranger" yaml:"is_stranger"`
	IsAppUser         bool      `boil:"is_app_user" json:"is_app_user" toml:"is_app_user" yaml:"is_app_user"`
	IsInvitedUser     bool      `boil:"is_invited_user" json:"is_invited_user" toml:"is_invited_user" yaml:"is_invited_user"`
	Has2fa            bool      `boil:"has_2fa" json:"has_2fa" toml:"has_2fa" yaml:"has_2fa"`
	HasFiles          bool      `boil:"has_files" json:"has_files" toml:"has_files" yaml:"has_files"`
	Presence          string    `boil:"presence" json:"presence" toml:"presence" yaml:"presence"`
	Locale            string    `boil:"locale" json:"locale" toml:"locale" yaml:"locale"`
	CreatedAt         time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt         time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserColumns = struct {
	ID                string
	TeamID            string
	Name              string
	Deleted           string
	Color             string
	RealName          string
	TZ                string
	TZLabel           string
	TZOffset          string
	IsBot             string
	IsAdmin           string
	IsOwner           string
	IsPrimaryOwner    string
	IsRestricted      string
	IsUltraRestricted string
	IsStranger        string
	IsAppUser         string
	IsInvitedUser     string
	Has2fa            string
	HasFiles          string
	Presence          string
	Locale            string
	CreatedAt         string
	UpdatedAt         string
}{
	ID:                "id",
	TeamID:            "team_id",
	Name:              "name",
	Deleted:           "deleted",
	Color:             "color",
	RealName:          "real_name",
	TZ:                "tz",
	TZLabel:           "tz_label",
	TZOffset:          "tz_offset",
	IsBot:             "is_bot",
	IsAdmin:           "is_admin",
	IsOwner:           "is_owner",
	IsPrimaryOwner:    "is_primary_owner",
	IsRestricted:      "is_restricted",
	IsUltraRestricted: "is_ultra_restricted",
	IsStranger:        "is_stranger",
	IsAppUser:         "is_app_user",
	IsInvitedUser:     "is_invited_user",
	Has2fa:            "has_2fa",
	HasFiles:          "has_files",
	Presence:          "presence",
	Locale:            "locale",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// Generated where

var UserWhere = struct {
	ID                whereHelperstring
	TeamID            whereHelperstring
	Name              whereHelperstring
	Deleted           whereHelperbool
	Color             whereHelperstring
	RealName          whereHelperstring
	TZ                whereHelperstring
	TZLabel           whereHelperstring
	TZOffset          whereHelperint64
	IsBot             whereHelperbool
	IsAdmin           whereHelperbool
	IsOwner           whereHelperbool
	IsPrimaryOwner    whereHelperbool
	IsRestricted      whereHelperbool
	IsUltraRestricted whereHelperbool
	IsStranger        whereHelperbool
	IsAppUser         whereHelperbool
	IsInvitedUser     whereHelperbool
	Has2fa            whereHelperbool
	HasFiles          whereHelperbool
	Presence          whereHelperstring
	Locale            whereHelperstring
	CreatedAt         whereHelpertime_Time
	UpdatedAt         whereHelpertime_Time
}{
	ID:                whereHelperstring{field: "\"user\".\"id\""},
	TeamID:            whereHelperstring{field: "\"user\".\"team_id\""},
	Name:              whereHelperstring{field: "\"user\".\"name\""},
	Deleted:           whereHelperbool{field: "\"user\".\"deleted\""},
	Color:             whereHelperstring{field: "\"user\".\"color\""},
	RealName:          whereHelperstring{field: "\"user\".\"real_name\""},
	TZ:                whereHelperstring{field: "\"user\".\"tz\""},
	TZLabel:           whereHelperstring{field: "\"user\".\"tz_label\""},
	TZOffset:          whereHelperint64{field: "\"user\".\"tz_offset\""},
	IsBot:             whereHelperbool{field: "\"user\".\"is_bot\""},
	IsAdmin:           whereHelperbool{field: "\"user\".\"is_admin\""},
	IsOwner:           whereHelperbool{field: "\"user\".\"is_owner\""},
	IsPrimaryOwner:    whereHelperbool{field: "\"user\".\"is_primary_owner\""},
	IsRestricted:      whereHelperbool{field: "\"user\".\"is_restricted\""},
	IsUltraRestricted: whereHelperbool{field: "\"user\".\"is_ultra_restricted\""},
	IsStranger:        whereHelperbool{field: "\"user\".\"is_stranger\""},
	IsAppUser:         whereHelperbool{field: "\"user\".\"is_app_user\""},
	IsInvitedUser:     whereHelperbool{field: "\"user\".\"is_invited_user\""},
	Has2fa:            whereHelperbool{field: "\"user\".\"has_2fa\""},
	HasFiles:          whereHelperbool{field: "\"user\".\"has_files\""},
	Presence:          whereHelperstring{field: "\"user\".\"presence\""},
	Locale:            whereHelperstring{field: "\"user\".\"locale\""},
	CreatedAt:         whereHelpertime_Time{field: "\"user\".\"created_at\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"user\".\"updated_at\""},
}

// UserRels is where relationship names are stored.
var UserRels = struct {
}{}

// userR is where relationships are stored.
type userR struct {
}

// NewStruct creates a new relationship struct
func (*userR) NewStruct() *userR {
	return &userR{}
}

// userL is where Load methods for each relationship are stored.
type userL struct{}

var (
	userAllColumns            = []string{"id", "team_id", "name", "deleted", "color", "real_name", "tz", "tz_label", "tz_offset", "is_bot", "is_admin", "is_owner", "is_primary_owner", "is_restricted", "is_ultra_restricted", "is_stranger", "is_app_user", "is_invited_user", "has_2fa", "has_files", "presence", "locale", "created_at", "updated_at"}
	userColumnsWithoutDefault = []string{"id", "team_id", "name", "deleted", "color", "real_name", "tz", "tz_label", "tz_offset", "is_bot", "is_admin", "is_owner", "is_primary_owner", "is_restricted", "is_ultra_restricted", "is_stranger", "is_app_user", "is_invited_user", "has_2fa", "has_files", "presence", "locale", "created_at", "updated_at"}
	userColumnsWithDefault    = []string{}
	userPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserSlice is an alias for a slice of pointers to User.
	// This should generally be used opposed to []User.
	UserSlice []*User
	// UserHook is the signature for custom User hook methods
	UserHook func(context.Context, boil.ContextExecutor, *User) error

	userQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userType                 = reflect.TypeOf(&User{})
	userMapping              = queries.MakeStructMapping(userType)
	userPrimaryKeyMapping, _ = queries.BindMapping(userType, userMapping, userPrimaryKeyColumns)
	userInsertCacheMut       sync.RWMutex
	userInsertCache          = make(map[string]insertCache)
	userUpdateCacheMut       sync.RWMutex
	userUpdateCache          = make(map[string]updateCache)
	userUpsertCacheMut       sync.RWMutex
	userUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userBeforeInsertHooks []UserHook
var userBeforeUpdateHooks []UserHook
var userBeforeDeleteHooks []UserHook
var userBeforeUpsertHooks []UserHook

var userAfterInsertHooks []UserHook
var userAfterSelectHooks []UserHook
var userAfterUpdateHooks []UserHook
var userAfterDeleteHooks []UserHook
var userAfterUpsertHooks []UserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *User) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *User) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *User) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *User) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *User) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *User) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *User) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *User) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *User) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserHook registers your hook function for all future operations.
func AddUserHook(hookPoint boil.HookPoint, userHook UserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userBeforeInsertHooks = append(userBeforeInsertHooks, userHook)
	case boil.BeforeUpdateHook:
		userBeforeUpdateHooks = append(userBeforeUpdateHooks, userHook)
	case boil.BeforeDeleteHook:
		userBeforeDeleteHooks = append(userBeforeDeleteHooks, userHook)
	case boil.BeforeUpsertHook:
		userBeforeUpsertHooks = append(userBeforeUpsertHooks, userHook)
	case boil.AfterInsertHook:
		userAfterInsertHooks = append(userAfterInsertHooks, userHook)
	case boil.AfterSelectHook:
		userAfterSelectHooks = append(userAfterSelectHooks, userHook)
	case boil.AfterUpdateHook:
		userAfterUpdateHooks = append(userAfterUpdateHooks, userHook)
	case boil.AfterDeleteHook:
		userAfterDeleteHooks = append(userAfterDeleteHooks, userHook)
	case boil.AfterUpsertHook:
		userAfterUpsertHooks = append(userAfterUpsertHooks, userHook)
	}
}

// One returns a single user record from the query.
func (q userQuery) One(ctx context.Context, exec boil.ContextExecutor) (*User, error) {
	o := &User{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all User records from the query.
func (q userQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserSlice, error) {
	var o []*User

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to User slice")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all User records in the query.
func (q userQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user exists")
	}

	return count > 0, nil
}

// Users retrieves all the records using an executor.
func Users(mods ...qm.QueryMod) userQuery {
	mods = append(mods, qm.From("\"user\""))
	return userQuery{NewQuery(mods...)}
}

// FindUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUser(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*User, error) {
	userObj := &User{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user")
	}

	return userObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *User) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userInsertCacheMut.RLock()
	cache, cached := userInsertCache[key]
	userInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAllColumns,
			userColumnsWithDefault,
			userColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userType, userMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userType, userMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"user\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, userPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into user")
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
		return errors.Wrap(err, "models: unable to populate default values for user")
	}

CacheNoHooks:
	if !cached {
		userInsertCacheMut.Lock()
		userInsertCache[key] = cache
		userInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the User.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *User) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userUpdateCacheMut.RLock()
	cache, cached := userUpdateCache[key]
	userUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAllColumns,
			userPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, userPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userType, userMapping, append(wl, userPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user")
	}

	if !cached {
		userUpdateCacheMut.Lock()
		userUpdateCache[key] = cache
		userUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all user")
	}
	return rowsAff, nil
}

// Delete deletes a single User record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *User) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no User provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userPrimaryKeyMapping)
	sql := "DELETE FROM \"user\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user")
	}

	if len(userAfterDeleteHooks) != 0 {
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
func (o *User) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user\".* FROM \"user\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserSlice")
	}

	*o = slice

	return nil
}

// UserExists checks if the User row exists.
func UserExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user exists")
	}

	return exists, nil
}
