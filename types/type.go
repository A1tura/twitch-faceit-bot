package types

import "time"

// Player represents a player in a team
type Player struct {
	PlayerID       string `json:"player_id"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	SkillLevel     int    `json:"skill_level"`
	GamePlayerID   string `json:"game_player_id"`
	GamePlayerName string `json:"game_player_name"`
	FaceitURL      string `json:"faceit_url"`
}

// Team represents a team in the match
type Team struct {
	TeamID   string   `json:"team_id"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Type     string   `json:"type"`
	Players  []Player `json:"players"`
}

// Match represents a single match
type Match struct {
	MatchID        string          `json:"match_id"`
	GameID         string          `json:"game_id"`
	Region         string          `json:"region"`
	GameMode       string          `json:"game_mode"`
	MaxPlayers     int             `json:"max_players"`
	TeamsSize      int             `json:"teams_size"`
	Teams          map[string]Team `json:"teams"`
	PlayingPlayers []string        `json:"playing_players"`
	CompetitionID  string          `json:"competition_id"`
	Status         string          `json:"status"`
	StartedAt      int             `json:"started_at"`
	Results        struct {
		Winner string         `json:"winner"`
		Score  map[string]int `json:"score"`
	} `json:"results"`
}

type Data struct {
	Items []Match `json:"items"`
}

type Stats struct {
	Loses int
	Wins  int
}

type ChatConfig struct {
	IrcServer       string
	Channel         string
	OauthToken      string
	BotUsername     string
	MessageCooldown time.Duration
}
