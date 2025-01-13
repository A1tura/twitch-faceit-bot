package stream

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// Define the structures that match the JSON response
type StreamData struct {
	ID           string   `json:"id"`
	UserID       string   `json:"user_id"`
	UserLogin    string   `json:"user_login"`
	UserName     string   `json:"user_name"`
	GameID       string   `json:"game_id"`
	GameName     string   `json:"game_name"`
	Type         string   `json:"type"`
	Title        string   `json:"title"`
	ViewerCount  int      `json:"viewer_count"`
	StartedAt    string   `json:"started_at"`
	Language     string   `json:"language"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Tags         []string `json:"tags"`
	IsMature     bool     `json:"is_mature"`
}

type Pagination struct {
	Cursor string `json:"cursor"`
}

type TwitchAPIResponse struct {
	Data       []StreamData `json:"data"`
	Pagination Pagination   `json:"pagination"`
}

type Stream struct {
	Streamer_username string
	ClientId          string
	Auth              string
}

func (stream *Stream) GetStreamStrartingTime() (bool, time.Time) {
	url := "https://api.twitch.tv/helix/streams?user_login=" + stream.Streamer_username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Client-ID", stream.ClientId)
	req.Header.Add("Authorization", "Bearer "+stream.Auth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	reader, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var response TwitchAPIResponse
    var startedAt time.Time
	err = json.Unmarshal([]byte(reader), &response)
	if err != nil {
		return false, time.Now()
	}

	if len(response.Data) == 1 {
		startedAt, err = time.Parse(time.RFC3339, response.Data[0].StartedAt)
		if err != nil {
			return false, time.Now()
		}
	} else {
        return false, time.Now()
    }

	return true, startedAt
}
