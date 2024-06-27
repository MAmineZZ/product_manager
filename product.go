package main

import (
	"fmt"
	"log"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func addProduct(name, description string, price float64) {
	query := "INSERT INTO products (name, description, price) VALUES (?, ?, ?)"
	_, err := db.Exec(query, name, description, price)
	if err != nil {
		log.Fatalf("Error adding product: %v", err)
	}
	fmt.Println("Product added successfully")
}

func listProducts() {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			log.Fatalf("Error scanning product: %v", err)
		}
		fmt.Printf("%d: %s - %s - $%.2f\n", p.ID, p.Name, p.Description, p.Price)
	}
}

func updateProduct(id int, name, description string, price float64) {
	query := "UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?"
	_, err := db.Exec(query, name, description, price, id)
	if err != nil {
		log.Fatalf("Error updating product: %v", err)
	}
	fmt.Println("Product updated successfully")
}

func deleteProduct(id int) {
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("Error deleting product: %v", err)
	}
	fmt.Println("Product deleted successfully")
}
