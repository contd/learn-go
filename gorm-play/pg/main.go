package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=jason password= dbname=links sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	user := User{
		Username: "kumpfjn",
		FirstName: "Jason",
		LastName: "Kumpf",
	}

	db.Create(&user)

	log.Println(user)
}

type User struct {
	ID uint
	Username string
	FirstName string
	LastName string
}
