package model

type GameInfo struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
	Deaths     map[string]int `json:"kills_by_means"`
}
