package utility

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/moutend/slack/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

// Ago converts time to human readable string.
func Ago(t time.Time) string {
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

func MessageReplacer() *strings.Replacer {
	return strings.NewReplacer(
		"&gt;", ">",
		"&lt;", "<",
		"&amp;", "&",
	)
}

func UserNameReplacer(users []*models.User) *strings.Replacer {
	patterns := make([]string, len(users)*4)

	for i, user := range users {
		patterns[i*4] = fmt.Sprintf("<@%s>", user.ID)
		patterns[i*4+1] = fmt.Sprintf("@%s", user.Name)
		patterns[i*4+2] = fmt.Sprintf("%s", user.ID)
		patterns[i*4+3] = fmt.Sprintf("%s", user.Name)
	}

	return strings.NewReplacer(patterns...)
}

func GetChannelIDByName(ctx context.Context, tx boil.ContextTransactor, name string) (string, error) {
	query := `
SELECT c.id AS id
FROM channel c
LEFT JOIN user u ON u.id = c.user
WHERE u.name = ? OR c.name = ?
`

	var channels []*struct {
		ID string `boil:"id"`
	}

	if err := queries.Raw(query, name, name).Bind(ctx, tx, &channels); err != nil {
		return "", fmt.Errorf("failed to find channel or user '%s': %w", name, err)
	}
	if len(channels) != 1 {
		return "", fmt.Errorf("failed to load channel or user '%s'", name)
	}

	return channels[0].ID, nil
}

func ExtractChannelIDAndMessageID(url string) (channelID, messageID string, err error) {
	if !strings.HasPrefix(url, "https://") {
		err = fmt.Errorf("utility: invalid URL")

		return
	}

	elem := strings.Split(strings.TrimPrefix(url, "https://"), "/")

	if len(elem) != 3 {
		err = fmt.Errorf("utility: URL must be 'https://example.slack.com/archives/xxx/yyy'")

		return
	}

	channelID = elem[1]
	messageID = elem[2]

	return
}
