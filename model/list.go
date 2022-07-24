package model

import (
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID   string `json:"id"`
	Task string `json:"task"`
}
