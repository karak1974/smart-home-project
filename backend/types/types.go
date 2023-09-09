package types

type EventLog struct {
	Id     int    `json:",omitempty"`
	Lamp   string `json:"lamp"`
	Date   string `json:",omitempty"`
	Status bool   `json:"status"`
}
