package model

import (
	"database/sql"
	"encoding/json"
	
)

type Orders struct {
	ID        string `dbq:"ID"`
	ProductID string `dbq:"product_id"`
	OrdersID  string `dbq:"orders_id"`
	Quantity  string `dbq:"quantity"`
}

func (Orders) GetTableName() string {
	return `orders`
}

func ListRowsInsertOrders() []string {
	return []string{
		"ID",
		"product_id",
		"orders_id",
		"quantity",
	}
}

func GetOrders(db *sql.DB, id string) (Orders, error) {
	var order Orders
	err := db.QueryRow("SELECT * FROM orders WHERE ID = ?", id).Scan(&order.ID, &order.ProductID, &order.OrdersID, &order.Quantity)
	if err != nil {
		return Orders{}, err
	}
	return order, nil
}

func GetListOrders(db *sql.DB) ([]Orders, error) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Orders
	for rows.Next() {
		var order Orders
		err := rows.Scan(&order.ID, &order.ProductID, &order.OrdersID, &order.Quantity)
		if err != nil {
			return nil, err
		}
		result = append(result, order)
	}
	return result, nil
}

func GetOrdersJSON(db *sql.DB, id string) (string, error) {
	order, err := GetOrders(db, id)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(order)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}

func GetListOrdersJSON(db *sql.DB) (string, error) {
	orderList, err := GetListOrders(db)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(orderList)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}
