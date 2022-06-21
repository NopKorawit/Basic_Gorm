package main

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {
	dsn := "server=localhost\\SQLEXPRESS;Database=GormTest;praseTime=true"
	dial := sqlserver.Open(dsn)
	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// DryRun: true, //ใช้รันโดยไม่สน database
	})
	if err != nil {
		panic(err)
	}

	// err1 := db.AutoMigrate(Gender{}, Test{})

	// if err1 != nil {
	// 	panic(err)
	// }

	// CreateGender("Female")
	GetGenders()

}

func GetGenders() {
	genders := []Gender{}
	tx := db.Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

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
