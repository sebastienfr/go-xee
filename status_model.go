package xee

// Status struct
type Status struct {
	Location *Location `json:"location"`
	Signals  []Signal  `json:"signals"`
}
