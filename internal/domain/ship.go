package domain

type Ship struct {
	LocString string
	Location [][]string
	Damage [][]string
	Knock bool
	End bool
}

type Board struct {
	Size int
	Loc [][]string
}

type Location struct {
	*Board
	Loc [][]string
}

type Shot struct {
	Destroy bool `json:"destroy"`
	Knock   bool `json:"knock"`
	End     bool `json:"end"`
}