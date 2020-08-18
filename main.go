package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=172.17.0.2 port=5432 user=postgres dbname=go-gorm password=mysecretpassword sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	initDb(db)

	updateUser(db)

	db.Where(&User{
		Username: "Test3",
	}).Delete(&User{})

	println("done")
}

func updateUser(db *gorm.DB) {
	searchedUser := searchUser(db)
	searchedUser.LastName = "TestUpdated"
	db.Save(&searchedUser)
	_ = searchUser(db)
}

func searchUser(db *gorm.DB) User {
	searchedUser := User{Username: "Test1"}
	db.Where(&searchedUser).First(&searchedUser)
	fmt.Println(searchedUser)
	return searchedUser
}

func initDb(db *gorm.DB) {
	db.DropTable(&User{})
	db.CreateTable(&User{})
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})

	for _, user := range users {
		db.Create(&user)
	}

	db.Debug().Model(&Calendar{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Debug().Save(&User{
		Username: "test",
		Calendar: Calendar{
			Name: "Calendar Shit",
		},
	})
}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Calendar  Calendar
}

type Calendar struct {
	gorm.Model
	Name   string
	UserId uint
}

var users []User = []User{
	User{
		Username:  "Test1",
		FirstName: "Test1",
		LastName:  "Test1",
	},
	User{
		Username:  "Test2",
		FirstName: "Test2",
		LastName:  "Test2",
	},
	User{
		Username:  "Test3",
		FirstName: "Test3",
		LastName:  "Test3",
	},
}
