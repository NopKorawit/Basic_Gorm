package repository

import (
	"fmt"
	"orm/handler"
	"orm/model"
	"strconv"
	"time"
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
func CreateTest( name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	code := Generate()
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

func Generate()(NewId int) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	test := model.Test{}
	db.Last(&test)
	// fmt.Println(test.Code / 10000)
	
	last := (test.Code / 10000)
	now,err := strconv.Atoi(time.Now().Format("200601"))
	if err !=nil{
		panic(err)
	}
	// fmt.Println(now.Format("200601"))
	
	if last == now {
		NewId := test.Code+1
		println(NewId) 
		return NewId
	}
	NewId = now*10000+1
	println(NewId)
	return NewId

}
