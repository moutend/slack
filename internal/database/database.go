package database

import (
	"context"
	"database/sql"
	"io"
	"io/ioutil"
	"log"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattn/go-sqlite3"
)

var (
	transactionInstanceContextKey      = &contextKey{"transaction-instance-context-key"}
	transactionForceRollbackContextKey = &contextKey{"transaction-force-rollback-context-key"}
)

type contextKey struct {
	Name string
}

type transaction struct {
	ID         string
	Transactor boil.ContextTransactor
}

// Manager manages database.
type Manager struct {
	dsn               string
	db                *sql.DB
	debug             *log.Logger
	enableDebugOutput bool
}

// TransactionalExecutable represents a callback function.
type TransactionalExecutable func(context.Context, boil.ContextTransactor) error

func withForceRollback(ctx context.Context) context.Context {
	return context.WithValue(ctx, transactionForceRollbackContextKey, struct{}{})
}

func (m *Manager) withTransaction(ctx context.Context) (context.Context, error) {
	tx, err := m.db.Begin()

	if err != nil {
		return ctx, errors.Wrap(err, "database: failed to begin transaction")
	}

	id := xid.New().String()

	return context.WithValue(ctx, transactionInstanceContextKey, transaction{Transactor: tx, ID: id}), nil
}

// Transaction begins transaction with given context.
func (m *Manager) Transaction(ctx context.Context, fn TransactionalExecutable) (err error) {
	if m == nil || m.db == nil {
		return errors.New("database: manager is not initialized yet")
	}

	v, hasTransaction := ctx.Value(transactionInstanceContextKey).(transaction)

	if !hasTransaction || v.Transactor == nil || v.ID == "" {
		ctx, err = m.withTransaction(ctx)

		if err != nil {
			return errors.Wrap(err, "database: failed to begin transaction")
		}

		v, _ = ctx.Value(transactionInstanceContextKey).(transaction)
	}

	id := v.ID
	tx := v.Transactor

	if !hasTransaction {
		m.debug.Printf("database: start transaction %s\n", id)

		defer m.debug.Printf("database: end transaction %s\n", id)
	}

	buffer := &SQLBuffer{id: id, Logger: m.debug}

	ctx = boil.WithDebug(ctx, m.enableDebugOutput)
	ctx = boil.WithDebugWriter(ctx, buffer)

	err = fn(ctx, tx)

	if hasTransaction {
		if err != nil {
			return errors.Wrapf(err, "database: error occurs while using transaction: %s", id)
		}

		return nil
	}
	if err != nil {
		err = errors.Wrapf(err, "database: error occurs while using transaction: %s", id)

		if rollbackError := tx.Rollback(); rollbackError != nil {
			return errors.Wrapf(rollbackError, "database: failed to rollback transaction: %s", id)
		}

		m.debug.Println("database:  complete rollback transaction: %s", id)

		return err
	}
	if _, ok := ctx.Value(transactionForceRollbackContextKey).(struct{}); ok {
		if err := tx.Rollback(); err != nil {
			return errors.Wrapf(err, "database: failed to force rollback transaction: %s", id)
		}

		m.debug.Printf("database: complete force rollback transaction: %s\n", id)

		return nil
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "database: failed to commit transaction: %s", id)
	}

	m.debug.Printf("database: complete commit transaction: %s\n", id)

	return nil
}

// Close closes connection.
func (m *Manager) Close() error {
	if err := m.db.Close(); err != nil {
		return errors.Wrap(err, "database: failed to close database")
	}

	return nil
}

// SetLogger sets logger for debugging.
func (m *Manager) SetLogger(w io.Writer) {
	if w != nil {
		m.debug = log.New(w, "debug: ", 0)
		m.enableDebugOutput = true

		return
	}

	m.debug = log.New(ioutil.Discard, "", 0)
	m.enableDebugOutput = false
}

// New returns prepared Manager.
func New(databaseFilePath, migrationDirectoryPath string) (*Manager, error) {
	dsn := "file:" + databaseFilePath + "?cache=shared&mode=rwc"
	db, err := sql.Open(`sqlite3`, dsn)

	if err != nil {
		return nil, errors.Wrap(err, "database: failed to create manager")
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationDirectoryPath,
		"sqlite3", driver)

	if err != nil {
		return nil, errors.Wrap(err, "database: failed to create manager")
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, errors.Wrap(err, "database: failed to create manager")
	}

	return &Manager{
		dsn:   dsn,
		db:    db,
		debug: log.New(ioutil.Discard, "", 0),
	}, nil
}
