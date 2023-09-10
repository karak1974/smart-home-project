package types

type EventLog struct {
	Id     int    `json:"id,omitempty"`
	Lamp   string `json:"lamp"`
	Date   string `json:"date,omitempty"`
	Status bool   `json:"status"`
}
