package models

import "strconv"

var Tracks = []string{"Programming", "Design", "Electronics"}

//type Grade int64

//const (
//	Poor Grade = (iota + 1) * 10
//	Fair
//	Good
//	VeryGood
//	Excellent
//)

type Task struct {
	ID          uint `json:"id"`
	TrackID     int
	Track       Track  `json:"-"`
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *Task) SetDay(day string) {
	var _ error
	t.Day, _ = strconv.Atoi(day)
}

func (t *Task) SetTrack(track string) {
	var _ error
	t.TrackID, _ = strconv.Atoi(track)
}
