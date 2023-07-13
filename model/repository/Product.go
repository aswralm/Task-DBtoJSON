package model

import (
	"database/sql"
	"encoding/json"
	
)

type Product struct {
	ID         string `dbq:"ID"`
	NameProduct string `dbq:"namaproduct"`
	Price      string `dbq:"price"`
}

func (Product) GetTableName() string {
	return `product`
}

func ListRowsInsertProduct() []string {
	return []string{
		"ID",
		"namaproduct",
		"price",
	}
}

func GetProduct(db *sql.DB, id string) (Product, error) {
	var product Product
	err := db.QueryRow("SELECT * FROM product WHERE ID = ?", id).Scan(&product.ID, &product.NameProduct, &product.Price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func GetListProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.NameProduct, &product.Price)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}
	return result, nil
}

func GetProductJSON(db *sql.DB, id string) (string, error) {
	product, err := GetProduct(db, id)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(product)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}

func GetListProductJSON(db *sql.DB) (string, error) {
	productList, err := GetListProduct(db)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(productList)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}
