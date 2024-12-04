package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	CommonModel
	UserDto
}

type UserDto struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        uint16 `json:"age"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Address    string `json:"address"`
}

func GetAllUsers() []User {
	user1 := User{
		CommonModel: CommonModel{
			Status:   true,
			Priority: 1,
		},
		UserDto: UserDto{
			Email:     "jEh0c@example.com",
			FirstName: "Prashant",
			LastName:  "Chapagain",
			Age:       21,
			Address:   "Kathmandu",
			Password:  "password",
		},
	}
	return []User{
		user1,
	}
}
