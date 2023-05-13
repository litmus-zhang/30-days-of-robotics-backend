package models

type Task struct {
	ID          string `json:"id"`
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
