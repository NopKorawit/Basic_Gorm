package repository

import (
	"fmt"
	"orm/handler"
	"orm/model"

	"gorm.io/gorm/clause"
)

func CreateCustomer(name string, genderID uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}

	customer := model.Customer{
		Name:     name,
		GenderId: genderID,
	}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetCustomer()
}

func GetCustomer() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	customers := []model.Customer{}
	// tx := db.Find(&customers)
	// tx := db.Preload("Gender").Find(&customers)
	tx := db.Preload(clause.Associations).Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	// fmt.Println(customers)
	for _, customers := range customers {
		fmt.Printf("%v|%v|%v\n", customers.Id, customers.Name, customers.Gender.Name)
	}
}
