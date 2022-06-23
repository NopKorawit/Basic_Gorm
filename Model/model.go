package model

import "gorm.io/gorm"

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size(10)"`
}

type Test struct {
	gorm.Model
	Name string `gorm:"column:Firstname;type:varchar(30);unique;default:Steven;not null"`
	Code string `gorm:"size:10"`
}

func (t Test) TableName() string {
	return "MyTest"
}
