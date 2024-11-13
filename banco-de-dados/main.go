package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	p := Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
	return &p
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_expert")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// INSERT E UPDATE
	// product := NewProduct("Mouse", 90.99)
	// InsertProduct(db, product)
	// UpdateProduct(db, product)

	// SELECT
	// product, err := SelectProduct(db, "59c74ee1-6513-4e67-a975-c4a4aa84c9fc")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(product.Name)

	// SELECT ALL
	// products, err := SelectAllProduct(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range products {
	// 	println(p.Name)
	// }

	// DELETE
	_ = DeleteProduct(db, "59c74ee1-6513-4e67-a975-c4a4aa84c9fc")
}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Id)
	if err != nil {
		return err
	}

	return nil
}

func SelectProduct(db *sql.DB, id string) (*Product, error) {
	p := Product{}
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_ = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Price)
	return &p, nil
}

func SelectAllProduct(db *sql.DB) ([]Product, error) {
	products := []Product{}

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.QueryRow(id)
	return nil
}
