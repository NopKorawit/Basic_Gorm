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

	repository.CreateTest("Golf11sadaa")
	// repository.CreateTest("Golf11sadas")
	// repository.CreateTest("Golf11sadsa")
	// repository.CreateTest(202206002, "Test2")
	// repository.CreateTest(202206003, "Test3")
	// repository.CreateTest(660, "T2")
	// repository.CreateTest(046544, "Test4")
	// repository.DeleteTest(27)
	// repository.DeleterealTest(23)
	repository.GetTests()
	// repository.Generate()

	// repository.CreateCustomer("Golf", 1)
	// repository.CreateCustomer("Nop", 1)
	// repository.CreateCustomer("Ing", 2)
	// repository.GetCustomer()

}
