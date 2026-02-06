package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type User struct{
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string
	Email string `gorm:"unique"`
}

func main(){
	dns := "host=localhost user=postgres password=123 dbname=test port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	fmt.Println("conected to database")

	db.AutoMigrate(&User{})

	// createUser(db)
	updateUser(db)

	readUsers(db)
	
	deleteUser(db)
	
	readUsers(db)
}

func createUser(db *gorm.DB){
	users := User{
		Name: "Akhil",
		Email: "a@gmail.com" ,
	}

	db.Create(&users)

	fmt.Println("Users table created", users.ID)
}

func readUsers(db *gorm.DB){
	var users []User

	db.Find(&users)

	fmt.Println("All users:")
	for _, u := range users{
		fmt.Println(u)
	}

}

func updateUser(db *gorm.DB){
	db.Model(&User{}).
	Where("id = ?", 1).
	Update("name", "junaid")
}

func deleteUser(db *gorm.DB){
	db.Delete(&User{}, 1)

	fmt.Println("user delete successfully")
}