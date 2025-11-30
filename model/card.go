package model

type CartItem struct {
	ID        int64   `json:"id"`
	UserID    int64   `json:"user_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	UserID int64      `json:"user_id"`
	Items  []CartItem `json:"items"`
	Total  float64    `json:"total"`
}
