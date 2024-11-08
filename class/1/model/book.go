package model

type Book struct {
	BookId   string `json:"bookId"`
	BookName string `json:"bookName"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	Subtotal int    `json:"subtotal"`
	Discount int    `json:"discount"`
}
