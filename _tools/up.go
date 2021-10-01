package main

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	dsn := "file:" + filepath.Join(wd, "cache.db3") + "?cache=shared&mode=rwc"

	db, err := sql.Open(`sqlite3`, dsn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+filepath.Join(wd, "_migrations"),
		"sqlite3", driver)

	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	m.Close()
}
