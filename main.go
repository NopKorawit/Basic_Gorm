package main

import (
	"orm/repository"
)

func main() {

	// repository.AutoMigrate()

	// repository.CreateGender("TestaXX1aX")
	// repository.GetGenders()
	// repository.GetGender(10)
	// repository.GetGenderByName("Male")
	// repository.UpdateGender(6, "YYYY") //select +save  2qury
	// repository.UpdateGender2(6, "zzzz") //1qury แต่ห้ามเป็น 0 และ ""
	// repository.DeleteGender(5)

	// repository.CreateTest(45501, "Test6")
	// repository.CreateTest(660, "T2")
	// repository.CreateTest(046544, "Test4")
	// repository.DeleteTest(2)
	// repository.DeleterealTest(10)
	// repository.GetTests()

	// repository.CreateCustomer("Golf", 1)
	// repository.CreateCustomer("Nop", 1)
	// repository.CreateCustomer("Ing", 2)
	repository.GetCustomer()

}
