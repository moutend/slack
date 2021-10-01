// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMessages(t *testing.T) {
	t.Parallel()

	query := Messages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Messages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MessageExists(ctx, tx, o.Timestamp)
	if err != nil {
		t.Errorf("Unable to check if Message exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MessageExists to return true, but got false.")
	}
}

func testMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	messageFound, err := FindMessage(ctx, tx, o.Timestamp)
	if err != nil {
		t.Error(err)
	}

	if messageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Messages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Messages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	messageOne := &Message{}
	messageTwo := &Message{}
	if err = randomize.Struct(seed, messageOne, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}
	if err = randomize.Struct(seed, messageTwo, messageDBTypes, false, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = messageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = messageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func messageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func messageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Message) error {
	*o = Message{}
	return nil
}

func testMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Message{}
	o := &Message{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, messageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Message object: %s", err)
	}

	AddMessageHook(boil.BeforeInsertHook, messageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterInsertHook, messageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	messageAfterInsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterSelectHook, messageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	messageAfterSelectHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpdateHook, messageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpdateHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpdateHook, messageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	messageAfterUpdateHooks = []MessageHook{}

	AddMessageHook(boil.BeforeDeleteHook, messageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	messageBeforeDeleteHooks = []MessageHook{}

	AddMessageHook(boil.AfterDeleteHook, messageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	messageAfterDeleteHooks = []MessageHook{}

	AddMessageHook(boil.BeforeUpsertHook, messageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	messageBeforeUpsertHooks = []MessageHook{}

	AddMessageHook(boil.AfterUpsertHook, messageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	messageAfterUpsertHooks = []MessageHook{}
}

func testMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(messageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MessageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Messages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	messageDBTypes = map[string]string{`Type`: `STRING`, `Channel`: `STRING`, `User`: `STRING`, `Text`: `STRING`, `Timestamp`: `STRING`, `ThreadTimestamp`: `STRING`, `IsStarred`: `BOOLEAN`, `LastRead`: `STRING`, `Subscribed`: `BOOLEAN`, `UnreadCount`: `INTEGER`, `SubType`: `STRING`, `Hidden`: `BOOLEAN`, `DeletedTimestamp`: `STRING`, `EventTimestamp`: `STRING`, `BotID`: `STRING`, `Name`: `STRING`, `ReplyCount`: `INTEGER`, `CreatedAt`: `DATETIME`}
	_              = bytes.MinRead
)

func testMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(messageAllColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(messageAllColumns) == len(messagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Message{}
	if err = randomize.Struct(seed, o, messageDBTypes, true, messageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Messages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, messageDBTypes, true, messagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Message struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(messageAllColumns, messagePrimaryKeyColumns) {
		fields = messageAllColumns
	} else {
		fields = strmangle.SetComplement(
			messageAllColumns,
			messagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MessageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
