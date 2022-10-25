package topic

import (
	"time"

	"theses/user"
)

type Topic struct {
	Id           uint        `json:"id" gorm:"primaryKey"`
	Title        string      `json:"title"`
	Language     []string    `json:"language"`
	Keywords     []string    `json:"keywords"`
	Description  string      `json:"description"`
	Requirements []string    `json:"requirements"`
	Users        []user.User `gorm:"many2many:topic_user;"`
	ExpiresAt    time.Time   `json:"expires_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	CreatedAt    time.Time   `json:"created_at"`
}
