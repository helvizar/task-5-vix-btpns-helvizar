package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model

	ID       uint `gorm:"primaryKey;autoIncrement:true"`
	Assignee string
	Content  string
	Date     string
	IsDone   bool
}
