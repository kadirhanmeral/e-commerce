package repository

import (
	"database/sql"
	"e-commerce/model"
)

type ProductRepository interface {
	Create(p *model.Product) error
	FindByID(id int64) (*model.Product, error)
	FindAll() ([]model.Product, error)
	Update(p *model.Product) error
	Delete(id int64) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *DB) ProductRepository {
	return &productRepository{db: db.Conn}
}

func (r *productRepository) Create(p *model.Product) error {
	_, err := r.db.Exec(`
        INSERT INTO products (name, description, price, stock)
        VALUES (?, ?, ?, ?)
    `, p.Name, p.Description, p.Price, p.Stock)
	return err
}

func (r *productRepository) FindByID(id int64) (*model.Product, error) {
	var p model.Product

	err := r.db.QueryRow(`
        SELECT id, name, description, price, stock 
        FROM products 
        WHERE id = ?
    `, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepository) FindAll() ([]model.Product, error) {
	rows, err := r.db.Query(`
        SELECT id, name, description, price, stock FROM products
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Product

	for rows.Next() {
		var p model.Product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)
		list = append(list, p)
	}

	return list, nil
}

func (r *productRepository) Update(p *model.Product) error {
	_, err := r.db.Exec(`
        UPDATE products 
        SET name=?, description=?, price=?, stock=? 
        WHERE id=?
    `, p.Name, p.Description, p.Price, p.Stock, p.ID)
	return err
}

func (r *productRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM products WHERE id=?`, id)
	return err
}
