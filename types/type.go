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

type MathStats struct {
	Kills     string
	Deads     string
	Assists   string
	Headshots string
	Kd        string
}

type ChatConfig struct {
	IrcServer       string
	Channel         string
	OauthToken      string
	BotUsername     string
	MessageCooldown time.Duration
}

type MatchData struct {
	Rounds []Round `json:"rounds"`
}

type Round struct {
	BestOf        string      `json:"best_of"`
	CompetitionID *string     `json:"competition_id"`
	GameID        string      `json:"game_id"`
	GameMode      string      `json:"game_mode"`
	MatchID       string      `json:"match_id"`
	MatchRound    string      `json:"match_round"`
	Played        string      `json:"played"`
	RoundStats    RoundStats  `json:"round_stats"`
	Teams         []TeamMatch `json:"teams"`
}

type RoundStats struct {
	Region string `json:"Region"`
	Rounds string `json:"Rounds"`
	Score  string `json:"Score"`
	Winner string `json:"Winner"`
	Map    string `json:"Map"`
}

type TeamMatch struct {
	TeamID    string        `json:"team_id"`
	Premade   bool          `json:"premade"`
	TeamStats TeamStats     `json:"team_stats"`
	Players   []PlayerMatch `json:"players"`
}

type TeamStats struct {
	FinalScore      string `json:"Final Score"`
	TeamHeadshots   string `json:"Team Headshots"`
	SecondHalfScore string `json:"Second Half Score"`
	Team            string `json:"Team"`
	TeamWin         string `json:"Team Win"`
	OvertimeScore   string `json:"Overtime score"`
	FirstHalfScore  string `json:"First Half Score"`
}

type PlayerMatch struct {
	PlayerID    string      `json:"player_id"`
	Nickname    string      `json:"nickname"`
	PlayerStats PlayerStats `json:"player_stats"`
}

type PlayerStats struct {
	// Only including a few fields as an example. Add others as needed.
	Kills                    string `json:"Kills"`
	Deaths                   string `json:"Deaths"`
	Assists                  string `json:"Assists"`
	KD                       string `json:"K/D Ratio"`
	ADR                      string `json:"ADR"`
	Headshots                string `json:"Headshots"`
	HeadshotsPercentage      string `json:"Headshots %"`
	FirstKills               string `json:"First Kills"`
	UtilityDamage            string `json:"Utility Damage"`
	FlashSuccessRatePerMatch string `json:"Flash Success Rate per Match"`
}
