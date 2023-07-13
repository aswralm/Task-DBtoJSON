package model

import (
	"database/sql"
	"encoding/json"

)

type Reviews struct {
	ID         string `dbq:"ID"`
	ProductID  string `dbq:"product_id"`
	OrderID    string `dbq:"order_id"`
	Rating     string `dbq:"rating"`
	ReviewText string `dbq:"review_text"`
}

func (Reviews) GetTableName() string {
	return `reviews`
}

func ListRowsInsertReviews() []string {
	return []string{
		"ID",
		"product_id",
		"order_id",
		"rating",
		"review_text",
	}
}

func GetReviews(db *sql.DB, id string) (Reviews, error) {
	var reviews Reviews
	err := db.QueryRow("SELECT * FROM reviews WHERE ID = ?", id).Scan(&reviews.ID, &reviews.ProductID, &reviews.OrderID, &reviews.Rating, &reviews.ReviewText)
	if err != nil {
		return Reviews{}, err
	}
	return reviews, nil
}

func GetListReview(db *sql.DB) ([]Reviews, error) {
	rows, err := db.Query("SELECT * FROM reviews")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Reviews
	for rows.Next() {
		var review Reviews
		err := rows.Scan(&review.ID, &review.ProductID, &review.OrderID, &review.Rating, &review.ReviewText)
		if err != nil {
			return nil, err
		}
		result = append(result, review)
	}
	return result, nil
}

func GetReviewJSON(db *sql.DB, id string) (string, error) {
	review, err := GetReviews(db, id)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(review)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}

func GetListReviewJSON(db *sql.DB) (string, error) {
	reviewList, err := GetListReview(db)
	if err != nil {
		return "", err
	}
	dataJSON, err := json.Marshal(reviewList)
	if err != nil {
		return "", err
	}
	return string(dataJSON), nil
}
