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

// File is an object representing the database table.
type File struct {
	ID                 string `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name               string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Title              string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Mimetype           string `boil:"mimetype" json:"mimetype" toml:"mimetype" yaml:"mimetype"`
	ImageExifRotation  int64  `boil:"image_exif_rotation" json:"image_exif_rotation" toml:"image_exif_rotation" yaml:"image_exif_rotation"`
	Filetype           string `boil:"filetype" json:"filetype" toml:"filetype" yaml:"filetype"`
	PrettyType         string `boil:"pretty_type" json:"pretty_type" toml:"pretty_type" yaml:"pretty_type"`
	User               string `boil:"user" json:"user" toml:"user" yaml:"user"`
	Mode               string `boil:"mode" json:"mode" toml:"mode" yaml:"mode"`
	Editable           bool   `boil:"editable" json:"editable" toml:"editable" yaml:"editable"`
	IsExternal         bool   `boil:"is_external" json:"is_external" toml:"is_external" yaml:"is_external"`
	ExternalType       string `boil:"external_type" json:"external_type" toml:"external_type" yaml:"external_type"`
	Size               int64  `boil:"size" json:"size" toml:"size" yaml:"size"`
	URLPrivate         string `boil:"url_private" json:"url_private" toml:"url_private" yaml:"url_private"`
	URLPrivateDownload string `boil:"url_private_download" json:"url_private_download" toml:"url_private_download" yaml:"url_private_download"`
	OriginalH          int64  `boil:"original_h" json:"original_h" toml:"original_h" yaml:"original_h"`
	OriginalW          int64  `boil:"original_w" json:"original_w" toml:"original_w" yaml:"original_w"`
	Thumb64            string `boil:"thumb64" json:"thumb64" toml:"thumb64" yaml:"thumb64"`
	Thumb80            string `boil:"thumb80" json:"thumb80" toml:"thumb80" yaml:"thumb80"`
	Thumb160           string `boil:"thumb160" json:"thumb160" toml:"thumb160" yaml:"thumb160"`
	Thumb360           string `boil:"thumb360" json:"thumb360" toml:"thumb360" yaml:"thumb360"`
	Thumb360Gif        string `boil:"thumb360_gif" json:"thumb360_gif" toml:"thumb360_gif" yaml:"thumb360_gif"`
	Thumb360W          int64  `boil:"thumb360_w" json:"thumb360_w" toml:"thumb360_w" yaml:"thumb360_w"`
	Thumb360H          int64  `boil:"thumb360_h" json:"thumb360_h" toml:"thumb360_h" yaml:"thumb360_h"`
	Thumb480           string `boil:"thumb480" json:"thumb480" toml:"thumb480" yaml:"thumb480"`
	Thumb480W          int64  `boil:"thumb480_w" json:"thumb480_w" toml:"thumb480_w" yaml:"thumb480_w"`
	Thumb480H          int64  `boil:"thumb480_h" json:"thumb480_h" toml:"thumb480_h" yaml:"thumb480_h"`
	Thumb720           string `boil:"thumb720" json:"thumb720" toml:"thumb720" yaml:"thumb720"`
	Thumb720W          int64  `boil:"thumb720_w" json:"thumb720_w" toml:"thumb720_w" yaml:"thumb720_w"`
	Thumb720H          int64  `boil:"thumb720_h" json:"thumb720_h" toml:"thumb720_h" yaml:"thumb720_h"`
	Thumb960           string `boil:"thumb960" json:"thumb960" toml:"thumb960" yaml:"thumb960"`
	Thumb960W          int64  `boil:"thumb960_w" json:"thumb960_w" toml:"thumb960_w" yaml:"thumb960_w"`
	Thumb960H          int64  `boil:"thumb960_h" json:"thumb960_h" toml:"thumb960_h" yaml:"thumb960_h"`
	Thumb1024          string `boil:"thumb1024" json:"thumb1024" toml:"thumb1024" yaml:"thumb1024"`
	Thumb1024W         int64  `boil:"thumb1024_w" json:"thumb1024_w" toml:"thumb1024_w" yaml:"thumb1024_w"`
	Thumb1024H         int64  `boil:"thumb1024_h" json:"thumb1024_h" toml:"thumb1024_h" yaml:"thumb1024_h"`
	Permalink          string `boil:"permalink" json:"permalink" toml:"permalink" yaml:"permalink"`
	PermalinkPublic    string `boil:"permalink_public" json:"permalink_public" toml:"permalink_public" yaml:"permalink_public"`
	EditLink           string `boil:"edit_link" json:"edit_link" toml:"edit_link" yaml:"edit_link"`
	Preview            string `boil:"preview" json:"preview" toml:"preview" yaml:"preview"`
	PreviewHighlight   string `boil:"preview_highlight" json:"preview_highlight" toml:"preview_highlight" yaml:"preview_highlight"`
	Lines              int64  `boil:"lines" json:"lines" toml:"lines" yaml:"lines"`
	LinesMore          int64  `boil:"lines_more" json:"lines_more" toml:"lines_more" yaml:"lines_more"`
	IsPublic           bool   `boil:"is_public" json:"is_public" toml:"is_public" yaml:"is_public"`
	PublicURLShared    bool   `boil:"public_url_shared" json:"public_url_shared" toml:"public_url_shared" yaml:"public_url_shared"`
	CommentsCount      int64  `boil:"comments_count" json:"comments_count" toml:"comments_count" yaml:"comments_count"`
	NumStars           int64  `boil:"num_stars" json:"num_stars" toml:"num_stars" yaml:"num_stars"`
	IsStarred          bool   `boil:"is_starred" json:"is_starred" toml:"is_starred" yaml:"is_starred"`

	R *fileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FileColumns = struct {
	ID                 string
	Name               string
	Title              string
	Mimetype           string
	ImageExifRotation  string
	Filetype           string
	PrettyType         string
	User               string
	Mode               string
	Editable           string
	IsExternal         string
	ExternalType       string
	Size               string
	URLPrivate         string
	URLPrivateDownload string
	OriginalH          string
	OriginalW          string
	Thumb64            string
	Thumb80            string
	Thumb160           string
	Thumb360           string
	Thumb360Gif        string
	Thumb360W          string
	Thumb360H          string
	Thumb480           string
	Thumb480W          string
	Thumb480H          string
	Thumb720           string
	Thumb720W          string
	Thumb720H          string
	Thumb960           string
	Thumb960W          string
	Thumb960H          string
	Thumb1024          string
	Thumb1024W         string
	Thumb1024H         string
	Permalink          string
	PermalinkPublic    string
	EditLink           string
	Preview            string
	PreviewHighlight   string
	Lines              string
	LinesMore          string
	IsPublic           string
	PublicURLShared    string
	CommentsCount      string
	NumStars           string
	IsStarred          string
}{
	ID:                 "id",
	Name:               "name",
	Title:              "title",
	Mimetype:           "mimetype",
	ImageExifRotation:  "image_exif_rotation",
	Filetype:           "filetype",
	PrettyType:         "pretty_type",
	User:               "user",
	Mode:               "mode",
	Editable:           "editable",
	IsExternal:         "is_external",
	ExternalType:       "external_type",
	Size:               "size",
	URLPrivate:         "url_private",
	URLPrivateDownload: "url_private_download",
	OriginalH:          "original_h",
	OriginalW:          "original_w",
	Thumb64:            "thumb64",
	Thumb80:            "thumb80",
	Thumb160:           "thumb160",
	Thumb360:           "thumb360",
	Thumb360Gif:        "thumb360_gif",
	Thumb360W:          "thumb360_w",
	Thumb360H:          "thumb360_h",
	Thumb480:           "thumb480",
	Thumb480W:          "thumb480_w",
	Thumb480H:          "thumb480_h",
	Thumb720:           "thumb720",
	Thumb720W:          "thumb720_w",
	Thumb720H:          "thumb720_h",
	Thumb960:           "thumb960",
	Thumb960W:          "thumb960_w",
	Thumb960H:          "thumb960_h",
	Thumb1024:          "thumb1024",
	Thumb1024W:         "thumb1024_w",
	Thumb1024H:         "thumb1024_h",
	Permalink:          "permalink",
	PermalinkPublic:    "permalink_public",
	EditLink:           "edit_link",
	Preview:            "preview",
	PreviewHighlight:   "preview_highlight",
	Lines:              "lines",
	LinesMore:          "lines_more",
	IsPublic:           "is_public",
	PublicURLShared:    "public_url_shared",
	CommentsCount:      "comments_count",
	NumStars:           "num_stars",
	IsStarred:          "is_starred",
}

// Generated where

var FileWhere = struct {
	ID                 whereHelperstring
	Name               whereHelperstring
	Title              whereHelperstring
	Mimetype           whereHelperstring
	ImageExifRotation  whereHelperint64
	Filetype           whereHelperstring
	PrettyType         whereHelperstring
	User               whereHelperstring
	Mode               whereHelperstring
	Editable           whereHelperbool
	IsExternal         whereHelperbool
	ExternalType       whereHelperstring
	Size               whereHelperint64
	URLPrivate         whereHelperstring
	URLPrivateDownload whereHelperstring
	OriginalH          whereHelperint64
	OriginalW          whereHelperint64
	Thumb64            whereHelperstring
	Thumb80            whereHelperstring
	Thumb160           whereHelperstring
	Thumb360           whereHelperstring
	Thumb360Gif        whereHelperstring
	Thumb360W          whereHelperint64
	Thumb360H          whereHelperint64
	Thumb480           whereHelperstring
	Thumb480W          whereHelperint64
	Thumb480H          whereHelperint64
	Thumb720           whereHelperstring
	Thumb720W          whereHelperint64
	Thumb720H          whereHelperint64
	Thumb960           whereHelperstring
	Thumb960W          whereHelperint64
	Thumb960H          whereHelperint64
	Thumb1024          whereHelperstring
	Thumb1024W         whereHelperint64
	Thumb1024H         whereHelperint64
	Permalink          whereHelperstring
	PermalinkPublic    whereHelperstring
	EditLink           whereHelperstring
	Preview            whereHelperstring
	PreviewHighlight   whereHelperstring
	Lines              whereHelperint64
	LinesMore          whereHelperint64
	IsPublic           whereHelperbool
	PublicURLShared    whereHelperbool
	CommentsCount      whereHelperint64
	NumStars           whereHelperint64
	IsStarred          whereHelperbool
}{
	ID:                 whereHelperstring{field: "\"file\".\"id\""},
	Name:               whereHelperstring{field: "\"file\".\"name\""},
	Title:              whereHelperstring{field: "\"file\".\"title\""},
	Mimetype:           whereHelperstring{field: "\"file\".\"mimetype\""},
	ImageExifRotation:  whereHelperint64{field: "\"file\".\"image_exif_rotation\""},
	Filetype:           whereHelperstring{field: "\"file\".\"filetype\""},
	PrettyType:         whereHelperstring{field: "\"file\".\"pretty_type\""},
	User:               whereHelperstring{field: "\"file\".\"user\""},
	Mode:               whereHelperstring{field: "\"file\".\"mode\""},
	Editable:           whereHelperbool{field: "\"file\".\"editable\""},
	IsExternal:         whereHelperbool{field: "\"file\".\"is_external\""},
	ExternalType:       whereHelperstring{field: "\"file\".\"external_type\""},
	Size:               whereHelperint64{field: "\"file\".\"size\""},
	URLPrivate:         whereHelperstring{field: "\"file\".\"url_private\""},
	URLPrivateDownload: whereHelperstring{field: "\"file\".\"url_private_download\""},
	OriginalH:          whereHelperint64{field: "\"file\".\"original_h\""},
	OriginalW:          whereHelperint64{field: "\"file\".\"original_w\""},
	Thumb64:            whereHelperstring{field: "\"file\".\"thumb64\""},
	Thumb80:            whereHelperstring{field: "\"file\".\"thumb80\""},
	Thumb160:           whereHelperstring{field: "\"file\".\"thumb160\""},
	Thumb360:           whereHelperstring{field: "\"file\".\"thumb360\""},
	Thumb360Gif:        whereHelperstring{field: "\"file\".\"thumb360_gif\""},
	Thumb360W:          whereHelperint64{field: "\"file\".\"thumb360_w\""},
	Thumb360H:          whereHelperint64{field: "\"file\".\"thumb360_h\""},
	Thumb480:           whereHelperstring{field: "\"file\".\"thumb480\""},
	Thumb480W:          whereHelperint64{field: "\"file\".\"thumb480_w\""},
	Thumb480H:          whereHelperint64{field: "\"file\".\"thumb480_h\""},
	Thumb720:           whereHelperstring{field: "\"file\".\"thumb720\""},
	Thumb720W:          whereHelperint64{field: "\"file\".\"thumb720_w\""},
	Thumb720H:          whereHelperint64{field: "\"file\".\"thumb720_h\""},
	Thumb960:           whereHelperstring{field: "\"file\".\"thumb960\""},
	Thumb960W:          whereHelperint64{field: "\"file\".\"thumb960_w\""},
	Thumb960H:          whereHelperint64{field: "\"file\".\"thumb960_h\""},
	Thumb1024:          whereHelperstring{field: "\"file\".\"thumb1024\""},
	Thumb1024W:         whereHelperint64{field: "\"file\".\"thumb1024_w\""},
	Thumb1024H:         whereHelperint64{field: "\"file\".\"thumb1024_h\""},
	Permalink:          whereHelperstring{field: "\"file\".\"permalink\""},
	PermalinkPublic:    whereHelperstring{field: "\"file\".\"permalink_public\""},
	EditLink:           whereHelperstring{field: "\"file\".\"edit_link\""},
	Preview:            whereHelperstring{field: "\"file\".\"preview\""},
	PreviewHighlight:   whereHelperstring{field: "\"file\".\"preview_highlight\""},
	Lines:              whereHelperint64{field: "\"file\".\"lines\""},
	LinesMore:          whereHelperint64{field: "\"file\".\"lines_more\""},
	IsPublic:           whereHelperbool{field: "\"file\".\"is_public\""},
	PublicURLShared:    whereHelperbool{field: "\"file\".\"public_url_shared\""},
	CommentsCount:      whereHelperint64{field: "\"file\".\"comments_count\""},
	NumStars:           whereHelperint64{field: "\"file\".\"num_stars\""},
	IsStarred:          whereHelperbool{field: "\"file\".\"is_starred\""},
}

// FileRels is where relationship names are stored.
var FileRels = struct {
}{}

// fileR is where relationships are stored.
type fileR struct {
}

// NewStruct creates a new relationship struct
func (*fileR) NewStruct() *fileR {
	return &fileR{}
}

// fileL is where Load methods for each relationship are stored.
type fileL struct{}

var (
	fileAllColumns            = []string{"id", "name", "title", "mimetype", "image_exif_rotation", "filetype", "pretty_type", "user", "mode", "editable", "is_external", "external_type", "size", "url_private", "url_private_download", "original_h", "original_w", "thumb64", "thumb80", "thumb160", "thumb360", "thumb360_gif", "thumb360_w", "thumb360_h", "thumb480", "thumb480_w", "thumb480_h", "thumb720", "thumb720_w", "thumb720_h", "thumb960", "thumb960_w", "thumb960_h", "thumb1024", "thumb1024_w", "thumb1024_h", "permalink", "permalink_public", "edit_link", "preview", "preview_highlight", "lines", "lines_more", "is_public", "public_url_shared", "comments_count", "num_stars", "is_starred"}
	fileColumnsWithoutDefault = []string{"id", "name", "title", "mimetype", "image_exif_rotation", "filetype", "pretty_type", "user", "mode", "editable", "is_external", "external_type", "size", "url_private", "url_private_download", "original_h", "original_w", "thumb64", "thumb80", "thumb160", "thumb360", "thumb360_gif", "thumb360_w", "thumb360_h", "thumb480", "thumb480_w", "thumb480_h", "thumb720", "thumb720_w", "thumb720_h", "thumb960", "thumb960_w", "thumb960_h", "thumb1024", "thumb1024_w", "thumb1024_h", "permalink", "permalink_public", "edit_link", "preview", "preview_highlight", "lines", "lines_more", "is_public", "public_url_shared", "comments_count", "num_stars", "is_starred"}
	fileColumnsWithDefault    = []string{}
	filePrimaryKeyColumns     = []string{"id"}
)

type (
	// FileSlice is an alias for a slice of pointers to File.
	// This should generally be used opposed to []File.
	FileSlice []*File
	// FileHook is the signature for custom File hook methods
	FileHook func(context.Context, boil.ContextExecutor, *File) error

	fileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fileType                 = reflect.TypeOf(&File{})
	fileMapping              = queries.MakeStructMapping(fileType)
	filePrimaryKeyMapping, _ = queries.BindMapping(fileType, fileMapping, filePrimaryKeyColumns)
	fileInsertCacheMut       sync.RWMutex
	fileInsertCache          = make(map[string]insertCache)
	fileUpdateCacheMut       sync.RWMutex
	fileUpdateCache          = make(map[string]updateCache)
	fileUpsertCacheMut       sync.RWMutex
	fileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fileBeforeInsertHooks []FileHook
var fileBeforeUpdateHooks []FileHook
var fileBeforeDeleteHooks []FileHook
var fileBeforeUpsertHooks []FileHook

var fileAfterInsertHooks []FileHook
var fileAfterSelectHooks []FileHook
var fileAfterUpdateHooks []FileHook
var fileAfterDeleteHooks []FileHook
var fileAfterUpsertHooks []FileHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *File) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *File) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *File) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *File) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *File) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *File) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *File) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *File) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *File) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fileAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFileHook registers your hook function for all future operations.
func AddFileHook(hookPoint boil.HookPoint, fileHook FileHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fileBeforeInsertHooks = append(fileBeforeInsertHooks, fileHook)
	case boil.BeforeUpdateHook:
		fileBeforeUpdateHooks = append(fileBeforeUpdateHooks, fileHook)
	case boil.BeforeDeleteHook:
		fileBeforeDeleteHooks = append(fileBeforeDeleteHooks, fileHook)
	case boil.BeforeUpsertHook:
		fileBeforeUpsertHooks = append(fileBeforeUpsertHooks, fileHook)
	case boil.AfterInsertHook:
		fileAfterInsertHooks = append(fileAfterInsertHooks, fileHook)
	case boil.AfterSelectHook:
		fileAfterSelectHooks = append(fileAfterSelectHooks, fileHook)
	case boil.AfterUpdateHook:
		fileAfterUpdateHooks = append(fileAfterUpdateHooks, fileHook)
	case boil.AfterDeleteHook:
		fileAfterDeleteHooks = append(fileAfterDeleteHooks, fileHook)
	case boil.AfterUpsertHook:
		fileAfterUpsertHooks = append(fileAfterUpsertHooks, fileHook)
	}
}

// One returns a single file record from the query.
func (q fileQuery) One(ctx context.Context, exec boil.ContextExecutor) (*File, error) {
	o := &File{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for file")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all File records from the query.
func (q fileQuery) All(ctx context.Context, exec boil.ContextExecutor) (FileSlice, error) {
	var o []*File

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to File slice")
	}

	if len(fileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all File records in the query.
func (q fileQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count file rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fileQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if file exists")
	}

	return count > 0, nil
}

// Files retrieves all the records using an executor.
func Files(mods ...qm.QueryMod) fileQuery {
	mods = append(mods, qm.From("\"file\""))
	return fileQuery{NewQuery(mods...)}
}

// FindFile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFile(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*File, error) {
	fileObj := &File{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"file\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, fileObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from file")
	}

	return fileObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *File) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fileColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fileInsertCacheMut.RLock()
	cache, cached := fileInsertCache[key]
	fileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fileAllColumns,
			fileColumnsWithDefault,
			fileColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fileType, fileMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"file\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"file\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"file\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, filePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into file")
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
		return errors.Wrap(err, "models: unable to populate default values for file")
	}

CacheNoHooks:
	if !cached {
		fileInsertCacheMut.Lock()
		fileInsertCache[key] = cache
		fileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the File.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *File) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fileUpdateCacheMut.RLock()
	cache, cached := fileUpdateCache[key]
	fileUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fileAllColumns,
			filePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update file, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"file\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, filePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fileType, fileMapping, append(wl, filePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update file row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for file")
	}

	if !cached {
		fileUpdateCacheMut.Lock()
		fileUpdateCache[key] = cache
		fileUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fileQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for file")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FileSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"file\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all file")
	}
	return rowsAff, nil
}

// Delete deletes a single File record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *File) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no File provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), filePrimaryKeyMapping)
	sql := "DELETE FROM \"file\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for file")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fileQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FileSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(fileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file")
	}

	if len(fileAfterDeleteHooks) != 0 {
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
func (o *File) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFile(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FileSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"file\".* FROM \"file\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, filePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FileSlice")
	}

	*o = slice

	return nil
}

// FileExists checks if the File row exists.
func FileExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"file\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if file exists")
	}

	return exists, nil
}
