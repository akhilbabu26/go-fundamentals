package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Fruit struct{
	ID int `gorm:"primaryKey;autoIncrement"`
	Name string
	Price float64
	Quantity int
}

func main(){
	dsn := "host=localhost user=postgres password=123 dbname=fruitsdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&Fruit{})
	fmt.Println("connected to database")
	// insert to db
	fruit := []Fruit{
		{Name: "Apple", Price: 1000, Quantity: 50},
		{Name: "Orange", Price:1200, Quantity: 200},
		{Name: "Mango", Price:750, Quantity: 40},
	}
	db.Create(&fruit)

	// update db
	db.Model(&Fruit{}).Where("id = ?", 2).Update("price", 80000)

	//delete from table
	db.Model(&Fruit{}).Where("id = ?", 1).Delete(&Fruit{})

	// read db
	var readfruits []Fruit
	db.Find(&readfruits)
	fmt.Println("All Fruits:")

	for _, f := range readfruits {
		fmt.Println(f)
	}
}