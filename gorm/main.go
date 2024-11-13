package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float32
	gorm.Model
}

func main() {
	dns := "root:123456@tcp(localhost:3306)/go_expert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	tx := db.Create(&Product{Name: "Vaio", Price: 3500})
	tx = db.Create(&Product{Name: "Acer", Price: 3000})
	tx = db.Create(&Product{Name: "Dell", Price: 3800})

	//product := Product{}
	products := []Product{}

	// tx = db.First(&product, "name = ?", "Vaio")
	if tx.Error != nil {
		log.Fatal(tx.Error)
	}

	// println(product.Name, product.Price)

	db.Where("price < ?", 3800).Limit(1).Offset(0).Find(&products)
	for _, p := range products {
		println(p.Name, p.Price)
	}

	db.Find(&products)
	products[1].Name = "Acer Raysen1"
	db.Save(&products[1])

	db.Delete(&products, "price < ?", 3500)
	println("DELETADO!")

	db.Find(&products)
	for _, p := range products {
		println(p.Name, p.Price)
	}

}
