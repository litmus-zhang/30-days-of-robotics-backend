package models

var Tracks = []string{"Programming", "Design", "Electronics"}

type Grade int64

const (
	Poor Grade = (iota + 1) * 10
	Fair
	Good
	VeryGood
	Excellent
)

type Task struct {
	ID          uint   `json:"id"`
	Track       string `json:"track"`
	Day         string `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
