package repository

import (
	"database/sql"
	"e-commerce/model"
)

type CartRepository interface {
	AddItem(item *model.CartItem) error
	GetCart(userID int64) ([]model.CartItem, error)
	ClearCart(userID int64) error
}

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *DB) CartRepository {
	return &cartRepository{db: db.Conn}
}

func (r *cartRepository) AddItem(item *model.CartItem) error {
	_, err := r.db.Exec(`
        INSERT INTO cart_items (user_id, product_id, quantity, price)
        VALUES (?, ?, ?, ?)
    `, item.UserID, item.ProductID, item.Quantity, item.Price)
	return err
}

func (r *cartRepository) GetCart(userID int64) ([]model.CartItem, error) {
	rows, err := r.db.Query(`
        SELECT id, user_id, product_id, quantity, price 
        FROM cart_items
        WHERE user_id = ?
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.CartItem

	for rows.Next() {
		var c model.CartItem
		rows.Scan(&c.ID, &c.UserID, &c.ProductID, &c.Quantity, &c.Price)
		items = append(items, c)
	}

	return items, nil
}

func (r *cartRepository) ClearCart(userID int64) error {
	_, err := r.db.Exec(`DELETE FROM cart_items WHERE user_id=?`, userID)
	return err
}
