package models

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	TrackID   int
	Track     Track `json:"-"`
}

type UserTask struct {
	UserID    string `json:"user_id"`
	TrackName string `json:"track_name"`
	Tasks     []Task `json:"tasks"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = string(hashedPassword)
}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}
func (user *User) SetTrack(track string) {
	user.TrackID, _ = strconv.Atoi(track)
}
