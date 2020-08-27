package database

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	maxDepth = 64
)

var (
	boilPath = filepath.Join("github.com", "volatiletech", "sqlboiler")

	whitespaceReplacer = strings.NewReplacer("\n", " ",
		"        ", " ",
		"       ", " ",
		"      ", " ",
		"     ", " ",
		"    ", " ",
		"   ", " ",
		"  ", " ",
	)
)

// SQLBuffer records queries.
type SQLBuffer struct {
	id        string
	callCount int
	Entries   []Entry
	Logger    *log.Logger
}

// Entry represents an SQL.
type Entry struct {
	Source    string
	Query     string
	Arguments string
}

func formatQuery(s string) string {
	return strings.TrimSpace(whitespaceReplacer.Replace(s))
}

func isQueryIssuedLocation(prevSource, currentSource string) bool {
	return strings.Contains(prevSource, boilPath) && !strings.Contains(currentSource, boilPath)
}

// Write implements io.Writer.
func (v *SQLBuffer) Write(p []byte) (int, error) {
	v.callCount++

	prevSource := ""
	currentSource := ""
	pcs := [maxDepth]uintptr{}
	n := runtime.Callers(1, pcs[:])

	for _, pc := range pcs[0:n] {
		caller := runtime.FuncForPC(pc)

		if caller == nil {
			continue
		}

		file, line := caller.FileLine(pc)
		currentSource = fmt.Sprintf("%s:%d", file, line)

		if isQueryIssuedLocation(prevSource, currentSource) {
			break
		}

		prevSource = currentSource
	}
	if v.callCount%2 == 1 {
		v.Entries = append(v.Entries, Entry{
			Source: currentSource,
			Query:  formatQuery(string(p)),
		})
	} else {
		v.Entries[len(v.Entries)-1].Arguments = strings.TrimSpace(string(p))

		v.Logger.Printf(
			"database: transaction: %s: query source: %s\n",
			v.id,
			v.Entries[len(v.Entries)-1].Source,
		)
		v.Logger.Printf(
			"database: transaction: %s: query string: %s; [%s]\n",
			v.id,
			v.Entries[len(v.Entries)-1].Query,
			v.Entries[len(v.Entries)-1].Arguments,
		)
	}

	return len(p), nil
}
