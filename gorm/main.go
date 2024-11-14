package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:primaryKey`
	Name     string
	Products []Product
}

type Line struct {
	ID       int `gorm:primaryKey`
	Name     string
	Products []Product `gorm:"many2many:products_lines;"`
}

type Product struct {
	ID           int `gorm:"primaryKey`
	Name         string
	Price        float32
	CategoryId   int
	Category     Category
	SerialNumber SerialNumber
	Lines        []Line `gorm:"many2many:products_lines;"`
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dns := "root:123456@tcp(localhost:3306)/go_expert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// tx := db.Create(&Category{Name: "Informática"})
	// tx = db.Create(&Category{Name: "Papelaria"})

	// line_1 := Line{
	// 	Name: "Linha 1",
	// }
	// line_2 := Line{
	// 	Name: "Linha 2",
	// }
	// tx = db.Create(&line_1)
	// tx = db.Create(&line_2)

	// tx = db.Create(&Product{Name: "Vaio", Price: 3500, CategoryId: 1, Lines: []Line{line_1, line_2}})
	// tx = db.Create(&Product{Name: "Acer", Price: 3000, CategoryId: 1, Lines: []Line{line_1, line_2}})
	// tx = db.Create(&Product{Name: "Dell", Price: 3800, CategoryId: 1, Lines: []Line{line_1, line_2}})

	// tx = db.Create(&Product{Name: "Caneta", Price: 1, CategoryId: 2, Lines: []Line{line_2}})
	// tx = db.Create(&Product{Name: "Caderno", Price: 12, CategoryId: 2, Lines: []Line{line_1}})

	// tx = db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: 1})
	// tx = db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: 2})
	// tx = db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: 3})
	// tx = db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: 4})
	// tx = db.Create(&SerialNumber{Number: uuid.New().String(), ProductID: 5})

	//product := Product{}
	//products := []Product{}

	// tx = db.First(&product, "name = ?", "Vaio")
	// if tx.Error != nil {
	// 	log.Fatal(tx.Error)
	// }

	// println(product.Name, product.Price)

	// db.Where("price < ?", 3800).Limit(1).Offset(0).Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Price, p.CategoriId, p.Category.Name)
	// }

	// db.Find(&products)
	// products[1].Name = "Acer Raysen1"
	// db.Save(&products[1])

	// db.Delete(&products, "price < ?", 3500)
	// println("DELETADO!")

	// db.Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Price)
	// }

	// db.Preload("Category").Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Price, p.CategoryId, p.Category.Name)
	// }

	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, p := range products {
	// 	println(p.Name, p.Price, p.CategoryId, p.Category.Name, p.SerialNumber.Number)
	// }

	// categories := []Category{}
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, c := range categories {
	// 	println(c.Name)
	// 	for _, p := range c.Products {
	// 		println("- ", p.Name, p.Price, p.SerialNumber.Number)
	// 	}
	// }

	lines := []Line{}
	err = db.Model(&Line{}).Preload("Products").Preload("Products.Category").Preload("Products.SerialNumber").Find(&lines).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range lines {
		println(l.Name)
		for _, p := range l.Products {
			println("- ", p.Category.Name, p.Name, p.Price, p.SerialNumber.Number)
		}
	}

}
