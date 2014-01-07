package types

type SuperTag struct {
	Variety      float64    `json:"variety"`
	Name         string     `json:"name"`
	Categories   Categories `json:"categories"`
	Score        float64    `json:"score"`
	Depth        float64    `json:"depth"`
	Positions    []Position `json:"positions"`
	ID           int64      `json:"id"`
	ContentMatch float64    `json:"contentMatch"`
	Type         int        `json:"type"`
	SenseRank    int        `json:"senseRank"`
}
