package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx/v3"
)

func exportProductsToExcel(filename string) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Products")
	if err != nil {
		log.Fatalf("Error creating sheet: %v", err)
	}

	row := sheet.AddRow()
	row.AddCell().Value = "ID"
	row.AddCell().Value = "Name"
	row.AddCell().Value = "Description"
	row.AddCell().Value = "Price"

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

		row := sheet.AddRow()
		row.AddCell().Value = fmt.Sprintf("%d", p.ID)
		row.AddCell().Value = p.Name
		row.AddCell().Value = p.Description
		row.AddCell().Value = fmt.Sprintf("%.2f", p.Price)
	}

	err = file.Save(filename)
	if err != nil {
		log.Fatalf("Error saving file: %v", err)
	}
	fmt.Println("Products exported to", filename)
}
