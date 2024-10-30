package models

import (
	"time"

	"gorm.io/gorm"
)

type ItemTagEnum uint

const (
	ItemTagEmpty ItemTagEnum = iota
)

type Item struct {
	gorm.Model
	Name          string    `json:"name" form:"name" gorm:"not null"`
	Location      string    `json:"location" form:"location" gorm:"not null"`
	FoundTime     time.Time `json:"found_time" form:"found_time" gorm:"not null"`
	IsComplete    bool      `json:"is_complete"`
	PhotoFileName string    `json:"photo_file_name"`
}
