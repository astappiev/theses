package topic

import (
	"time"

	"theses/user"
)

type Topic struct {
	Id        uint        `json:"id" gorm:"primaryKey"`
	Title     string      `json:"title"`
	Abstract  string      `json:"abstract"`
	Users     []user.User `gorm:"many2many:topic_user;"`
	UpdatedAt time.Time   `json:"updated_at"`
	CreatedAt time.Time   `json:"created_at"`
	ExpiresAt time.Time   `json:"expires_at"`
}
