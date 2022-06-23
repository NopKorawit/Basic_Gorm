package repository

import (
	"database/sql"
	"fmt"
	"orm/handler"
	"orm/model"
)

// แยกไม่ได้ใส่มาก่อน

//---------------------------------------------------------------

func GetGenders() {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	genders := []model.Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	gender := model.Gender{}
	tx := db.Find(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName(name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	gender := model.Gender{}
	tx := db.Where("name=?", name).Find(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func CreateGender(name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	gender := model.Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func UpdateGender(id uint, name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	gender := model.Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func UpdateGender2(id uint, name string) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	gender := model.Gender{Name: name} //ไม่อัพเดตตัว 0 หรือ ""
	// tx := db.Model(&model.Gender{}).Where("id=?", id).Updates(gender)
	tx := db.Model(&model.Gender{}).Where("id=@id", sql.Named("id", id)).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func DeleteGender(id uint) {
	db, err := handler.DB()
	if err != nil {
		panic(err)
	}
	tx := db.Delete(&model.Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Printf("delete id:%v", id)
	GetGender(id)
}
