package main

import (
	"fmt"
	"net/http"
)

func handleProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintln(w, "<h1>List Produits</h1>")
	fmt.Fprintln(w, "<ul>")
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price)
		if err != nil {
			http.Error(w, "Error scanning product", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "<li>%d: %s - %s - $%.2f</li>", p.ID, p.Name, p.Description, p.Price)
	}
	fmt.Fprintln(w, "</ul>")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Server Produits</h1>")
	fmt.Fprintln(w, "<p><a href='/products'>Afficher les Produits</a></p>")
}

func startServer() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/products", handleProducts)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
