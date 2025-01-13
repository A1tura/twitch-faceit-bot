package faceit

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"twitch/types"
)

func isWinner(teamA, teamB types.Team, winner string) bool {
	var streamerTeam string
	for _, player := range teamA.Players {
		if player.Nickname == "gleban22" {
			streamerTeam = "faction1"
		}
	}
	for _, player := range teamB.Players {
		if player.Nickname == "gleban22" {
			streamerTeam = "faction2"
		}
	}

	if winner == streamerTeam {
		return true
	} else {
		return false
	}
}

func Get_mounth_stats() types.Stats {
	url := "https://open.faceit.com/data/v4/players/" + os.Getenv("FACEIT_ID") + "/history?game=cs2&limit=100"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer " + os.Getenv("FACEIT_API"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data types.Data
	err = json.Unmarshal(reader, &data)
	if err != nil {
		panic(err)
	}

	var stats types.Stats

	for _, match := range data.Items {
		winner := isWinner(match.Teams["faction1"], match.Teams["faction2"], match.Results.Winner)

		if winner {
			stats.Wins = stats.Wins + 1
		} else {
			stats.Loses = stats.Loses + 1
		}
	}

	return stats
}

func Get_day_stats(streamStartedAt time.Time) types.Stats {
	url := "https://open.faceit.com/data/v4/players/" + os.Getenv("FACEIT_ID") + "/history?game=cs2&from=" + strconv.Itoa(int(streamStartedAt.Unix()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer " + os.Getenv("FACEIT_API"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data types.Data
	err = json.Unmarshal(reader, &data)
	if err != nil {
		panic(err)
	}

	var stats types.Stats

	for _, match := range data.Items {
		if match.StartedAt >= int(streamStartedAt.Unix()) {
			winner := isWinner(match.Teams["faction1"], match.Teams["faction2"], match.Results.Winner)

			if winner {
				stats.Wins = stats.Wins + 1
			} else {
				stats.Loses = stats.Loses + 1
			}
		}
	}

	return stats
}
