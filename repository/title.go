package repository

import (
	"errors"
	"time"
	"up-meetup-clone/database"

	"gorm.io/gorm"
)

//Title is pattern from table users
type Title struct {
	ID        int            `gorm:"column:id"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

var tableName = "title"

// AllTitle is get all title
func AllTitle() interface{} {
	db := database.Connect()
	var title []Title
	result := db.Table(tableName).Find(&title)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return []string{}
	}

	return title
}

// FindByID is find title by id
func FindByID(id string) interface{} {
	db := database.Connect()
	var title Title
	result := db.Table(tableName).First(&title, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return title
}

// Insert is for insert data
func Insert() interface{} {
	db := database.Connect()
	result := db.Create(&Title{
		Name: "jinzhu",
	})

	return result
}
