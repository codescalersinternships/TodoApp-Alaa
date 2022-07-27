package model

type List struct {
	ID   string `json:"id" gorm:"primary_key"`
	Task string `json:"task"`
}
