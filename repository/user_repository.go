package repository

import (
	"database/sql"
	"e-commerce/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id int64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll() ([]model.User, error)
	Update(user *model.User) error
	Delete(id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *DB) UserRepository {
	return &userRepository{db: db.Conn}
}

func (r *userRepository) Create(u *model.User) error {
	_, err := r.db.Exec(`
        INSERT INTO users (username, email, password)
        VALUES (?, ?, ?)
    `, u.Username, u.Email, u.Password)

	return err
}

func (r *userRepository) FindByID(id int64) (*model.User, error) {
	u := &model.User{}
	err := r.db.QueryRow(`
        SELECT id, username, email, password FROM users WHERE id = ?
    `, id).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.db.QueryRow(`
        SELECT id, username, email, password FROM users WHERE email = ?
    `, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	rows, err := r.db.Query(`
        SELECT id, username, email, password FROM users
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password)
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) Update(u *model.User) error {
	_, err := r.db.Exec(`
        UPDATE users SET username=?, email=?, password=? WHERE id=?
    `, u.Username, u.Email, u.Password, u.ID)

	return err
}

func (r *userRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=?`, id)
	return err
}
