package model

import "time"

type Review struct {
	ReviewId     string    `json:"reviewId,omitempty"`
	OrderId      string    `json:"orderId"`
	BookId       string    `json:"bookId"`
	CustomerName string    `json:"customerName"`
	Rating       float64   `json:"rating"`
	ReviewText   string    `json:"reviewText"`
	ReviewDate   time.Time `json:"reviewDate"`
}
