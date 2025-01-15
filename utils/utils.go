package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"twitch/types"
)

func FormatStatsMessage(username string, stats types.Stats) string {
	var message string

	if os.Getenv("LANG") == "ru" {
		message = fmt.Sprintf("@"+username+" Победы: %d || Поражение: %d", stats.Wins, stats.Loses)
	} else {
		message = fmt.Sprintf("@"+username+" Wins: %d || Looses: %d", stats.Wins, stats.Loses)
	}

	return message
}

func FormatMatchStatsString(username string, stats types.MathStats) string {
	var message string
	if os.Getenv("LANG") == "ru" {
		message = fmt.Sprintf("@"+username+" Убийства: %s || Смерти: %s || Помощь: %s || КД: %s || %% В голову: %s%%", stats.Kills, stats.Deads, stats.Assists, stats.Kd, stats.Headshots)
	} else {
		message = fmt.Sprintf("@"+username+" Kills: %s || Deaths: %s || Assists: %s || K/D: %s || %% Headshots: %s%%", stats.Kills, stats.Deads, stats.Assists, stats.Kd, stats.Headshots)
	}

	return message
}

func ParseUsername(rawMessage string) string {
	// Example raw message: ":username!username@username.tmi.twitch.tv PRIVMSG #channel :Hello, world!"
	if len(rawMessage) > 0 && rawMessage[0] == ':' {
		// Split by space to isolate the prefix
		parts := strings.SplitN(rawMessage, " ", 2)
		if len(parts) > 0 {
			// Prefix is the first part, e.g., ":username!username@username.tmi.twitch.tv"
			prefix := parts[0]

			// Remove the leading ':' and split by '!' to get the username
			username := strings.SplitN(prefix[1:], "!", 2)[0]
			return username
		}
	}
	return ""
}

func BeginningOfMonthTimestamp() string {
	// Get the current time
	now := time.Now()

	// Create a new time value set to the beginning of the current month
	beginningOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	// Convert to Unix timestamp
	return strconv.Itoa(int(beginningOfMonth.Unix()))
}

func EndOfMonthTimestamp() string {
	// Get the current time
	now := time.Now()

	// Calculate the beginning of the next month
	beginningOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC)

	// Subtract one second to get the last moment of the current month
	endOfMonth := beginningOfNextMonth.Add(-time.Second)

	// Convert to Unix timestamp
	return strconv.Itoa(int(endOfMonth.Unix()))
}

func StartOfDay(t time.Time) string {
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return strconv.Itoa(int(startOfDay.Unix()))
}

func EndOfDay(t time.Time) string {
	year, month, day := t.Date()
	endOfDay := time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)
	return strconv.Itoa(int(endOfDay.Unix()))
}
