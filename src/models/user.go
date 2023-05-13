package models

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email,unique"`
	Password  string `json:"password"`
}

type UserTrack struct {
	UserID    string `json:"user_id"`
	TrackID   string `json:"track_id"`
	TrackName string `json:"track_name"`
}
