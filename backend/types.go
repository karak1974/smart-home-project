package main

type EventLog struct {
	Id     int    `json:"id"`
	Lamp   string `json:"state"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}
