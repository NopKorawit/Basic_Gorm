package model

import "gorm.io/gorm"

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size(10)"`
}

type Test struct {
	gorm.Model
	Name string `gorm:"column:Firstname;type:varchar(30);unique;default:Steven;not null"`
	Code uint   `gorm:"size:50"`
}

func (t Test) TableName() string {
	return "MyTest"
}

type Customer struct {
	Id       uint
	Name     string
	Gender   Gender
	GenderId uint
}
