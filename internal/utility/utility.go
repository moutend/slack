package utility

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

// SortMessagesByTimestamp sorts messages by descending order.
func SortMessagesByTimestamp(messages []slack.Message) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("client: failed to sort messages: %v", r)
		}
	}()

	sort.Slice(messages, func(i, j int) bool {
		a := strings.Split(messages[i].Timestamp, ".")
		b := strings.Split(messages[j].Timestamp, ".")

		if len(a) != 2 || len(b) != 2 {
			panic("timestamp is broken")
		}

		sec1, err := strconv.ParseInt(a[0], 10, 64)

		if err != nil {
			panic(err)
		}

		nano1, err := strconv.ParseInt(a[1], 10, 64)

		if err != nil {
			panic(err)
		}
		sec2, err := strconv.ParseInt(b[0], 10, 64)

		if err != nil {
			panic(err)
		}

		nano2, err := strconv.ParseInt(b[1], 10, 64)

		if err != nil {
			panic(err)
		}

		t1 := time.Unix(sec1, nano1)
		t2 := time.Unix(sec2, nano2)

		return t1.After(t2)
	})

	return nil
}

// FormatMessageText formats slack.Message.Text.
func FormatMessageText(s string, userIdToName map[string]string) string {
	pairs := []string{}

	for userId, userName := range userIdToName {
		pairs = append(
			pairs,
			fmt.Sprintf("<@%s>", userId),
			"@"+userName,
		)
	}
	r := strings.NewReplacer(pairs...)

	return r.Replace(s)
}

// Ago converts UNIX timestamp to human readable string.
func Ago(timestamp string) string {
	s := strings.Split(timestamp, ".")

	if len(s) < 2 {
		return ""
	}

	sec, _ := strconv.Atoi(s[0])
	msec, _ := strconv.Atoi(s[1])
	t := time.Unix(int64(sec), int64(msec))
	duration := time.Now().UTC().Sub(t)

	const (
		day  = 24 * time.Hour
		week = 7 * day
	)

	if duration < 2*time.Minute {
		return "a minute ago"
	} else if duration < time.Hour {
		return fmt.Sprintf("%v minutes ago", int(duration/time.Minute))
	} else if duration < 2*time.Hour {
		return "an hour ago"
	} else if duration < day {
		return fmt.Sprintf("%v hours ago", int(duration/time.Hour))
	} else if duration < 2*day {
		return "yesterday"
	} else if duration < week {
		return fmt.Sprintf("%v days ago", int(duration/day))
	} else if duration < 2*week {
		return "a week ago"
	} else if duration < 5*week {
		return fmt.Sprintf("%v weeks ago", int(duration/week))
	} else {
		return fmt.Sprintf("%v %v, %v", t.Day(), t.Month(), t.Year())
	}
}
