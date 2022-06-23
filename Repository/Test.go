package repository

import (
	"fmt"
	"orm/handler"
	"orm/model"
)

func AutoMigrate() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	err1 := db.AutoMigrate(model.Gender{}, model.Test{}, model.Customer{})
	if err1 != nil {
		panic(err1)
	}
}

//----------------------------------------------------------------------------
func CreateTest(code uint, name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	test := model.Test{Name: name, Code: code}
	db.Create(&test)

}

func GetTests() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	tests := []model.Test{}
	db.Find(&tests)
	for _, t := range tests {
		fmt.Printf("%v|%v|%v\n", t.ID, t.Name, t.Code)
	}
}

func DeleteTest(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	db.Delete(&model.Test{}, id)
}

func DeleterealTest(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	db.Unscoped().Delete(&model.Test{}, id)
}
