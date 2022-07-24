package models

import (
	"fmt"

	"github.com/project/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Value       float64
	Amount      int
}

func GetProducts() ([]Product, error) {
	db := db.ConectDB()

	query, err := db.Query("SELECT id, name, description, value, amount FROM products ORDER BY id ASC")
	if err != nil {
		return nil, fmt.Errorf("Error querying: %s", err)
	}

	p := Product{}
	products := []Product{}

	for query.Next() {
		var id, amount int
		var name, description string
		var value float64

		err = query.Scan(&id, &name, &description, &value, &amount)
		if err != nil {
			return nil, fmt.Errorf("Error scanning query: ", err)
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Value = value
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products, nil
}

func GetProductByID(id string) (Product, error) {
	db := db.ConectDB()

	query, err := db.Query("SELECT id, name, description, value, amount FROM products WHERE id=$1", id)
	if err != nil {
		return Product{}, fmt.Errorf("Error preparing query: %s", err)
	}

	product := Product{}

	for query.Next() {
		var id, amount int
		var name, description string
		var value float64

		err = query.Scan(&id, &name, &description, &value, &amount)
		if err != nil {
			return Product{}, fmt.Errorf("Error scanning query: %s", err)
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Value = value
		product.Amount = amount
	}
	defer db.Close()
	return product, nil
}

func InsertProduct(name, description string, value float64, amount int) error {
	db := db.ConectDB()

	query, err := db.Prepare("INSERT INTO products(name, description, value, amount) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("Error preparing query: %s", err)
	}
	query.Exec(name, description, value, amount)

	defer db.Close()
	return nil
}

func DeleteProduct(id string) error {
	db := db.ConectDB()

	query, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		return fmt.Errorf("Error preparing query: %s", err)
	}
	query.Exec(id)

	defer db.Close()
	return nil
}

func UpdateProduct(id int, name, description string, value float64, amount int) error {
	db := db.ConectDB()

	query, err := db.Prepare("UPDATE products SET name=$1, description=$2, value=$3, amount=$4 WHERE id=$5")
	if err != nil {
		return fmt.Errorf("Error querying: %s", err)
	}

	query.Exec(name, description, value, amount, id)

	defer db.Close()
	return nil
}
