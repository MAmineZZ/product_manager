package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	initDB()
	defer closeDB()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Ajouter un produit")
		fmt.Println("2. Afficher la liste des produits")
		fmt.Println("3. Modifier un produit")
		fmt.Println("4. Supprimer un produit")
		fmt.Println("5. Exporter les informations produits dans un fichier Excel")
		fmt.Println("6. Lancer un serveur Http avec une page web")
		fmt.Println("7. Se connecter à une VM en ssh")
		fmt.Println("8. Se connecter à un serveur FTP")
		fmt.Println("9. Lancer l'interface graphique")
		fmt.Println("10. Quitter")
		fmt.Print("Choisissez une option: ")

		input, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(input))

		switch option {
		case 1:
			fmt.Print("Nom du produit: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Description du produit: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Prix du produit: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			addProduct(name, description, price)

		case 2:
			listProducts()

		case 3:
			fmt.Print("ID du produit à modifier: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			fmt.Print("Nouveau nom du produit: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Nouvelle description du produit: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Nouveau prix du produit: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			updateProduct(id, name, description, price)

		case 4:
			fmt.Print("ID du produit à supprimer: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))
			deleteProduct(id)

		case 5:
			fmt.Print("Nom du fichier Excel (avec extension .xlsx): ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSpace(filename)
			exportProductsToExcel(filename)

		case 6:
			startServer()

		case 7:
			fmt.Print("SSH User: ")
			user, _ := reader.ReadString('\n')
			user = strings.TrimSpace(user)

			fmt.Print("SSH Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			fmt.Print("SSH Host: ")
			host, _ := reader.ReadString('\n')
			host = strings.TrimSpace(host)

			fmt.Print("SSH Port: ")
			portInput, _ := reader.ReadString('\n')
			port, _ := strconv.Atoi(strings.TrimSpace(portInput))

			connectSSH(user, password, host, port)

		case 8:
			fmt.Print("FTP Host: ")
			host, _ := reader.ReadString('\n')
			host = strings.TrimSpace(host)

			fmt.Print("FTP Port: ")
			portInput, _ := reader.ReadString('\n')
			port, _ := strconv.Atoi(strings.TrimSpace(portInput))

			fmt.Print("FTP User: ")
			user, _ := reader.ReadString('\n')
			user = strings.TrimSpace(user)

			fmt.Print("FTP Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			connectFTP(host, port, user, password)

		case 9:
			//startGUI()

		case 10:
			fmt.Println("Au revoir!")
			return

		default:
			fmt.Println("Option invalide. Veuillez réessayer.")
		}
	}
}
