package main

import (
	"orm/repository"
)

func main() {

	// err1 := db.AutoMigrate(Gender{}, Test{}, Customer{})
	// if err1 != nil {
	// 	panic(err1)
	// }

	repository.CreateGender("TestXX1aX")
	repository.GetGenders()
	// GetGender(2)
	// GetGenderByName("Male")
	// UpdateGender(5, "YYYY") //select +save  2qury
	// UpdateGender2(6, "zzzz") //1qury แต่ห้ามเป็น 0 และ ""
	// DeleteGender(5)

	// CreateTest(0, "Test1")
	// CreateTest(6460, "Test2")
	// CreateTest(046544, "Test4")
	// DeleteTest(3)
	// DeleterealTest(4)
	// GetTests()

	// CreateCustomer("Golf", 1)
	// CreateCustomer("Nop", 1)
	// CreateCustomer("Ing", 2)
	// GetCustomer()

}
