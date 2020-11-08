package repository

import (
	"time"
)

//Users is pattern from table users
type Users struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// FindAuthorize is for find authorize data
func FindAuthorize(email string, password string) {

}
