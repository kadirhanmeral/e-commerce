package service

import (
	"e-commerce/model"
	"testing"
)

type MockUserRepo struct {
	FindByEmailFunc func(email string) (*model.User, error)
	CreateFunc      func(user *model.User) error
}

func (m *MockUserRepo) Create(user *model.User) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(user)
	}
	return nil
}

func (m *MockUserRepo) FindByEmail(email string) (*model.User, error) {
	if m.FindByEmailFunc != nil {
		return m.FindByEmailFunc(email)
	}
	return nil, nil
}

func (m *MockUserRepo) FindByID(id int64) (*model.User, error) { return nil, nil }
func (m *MockUserRepo) FindAll() ([]model.User, error)         { return nil, nil }
func (m *MockUserRepo) Update(user *model.User) error          { return nil }
func (m *MockUserRepo) Delete(id int64) error                  { return nil }

func TestRegister(t *testing.T) {
	mock := &MockUserRepo{
		FindByEmailFunc: func(email string) (*model.User, error) { return nil, nil },
		CreateFunc:      func(user *model.User) error { return nil },
	}
	service := NewUserService(mock)

	err := service.Register(&model.User{
		Username: "kadirhanmeral",
		Email:    "kadirhanmeral@example.com",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLogin_Success(t *testing.T) {
	mock := &MockUserRepo{
		FindByEmailFunc: func(email string) (*model.User, error) {
			return &model.User{Username: "kadirhanmeral", Email: email, Password: "password"}, nil
		},
	}
	service := NewUserService(mock)

	user, err := service.Login("kadirhanmeral@example.com", "password")
	if err != nil || user == nil || user.Username != "kadirhanmeral" {
		t.Errorf("Expected successful login, got user=%v, err=%v", user, err)
	}
}

func TestLogin_Failure(t *testing.T) {
	mock := &MockUserRepo{
		FindByEmailFunc: func(email string) (*model.User, error) {
			return &model.User{Username: "kadirhanmeral", Email: email, Password: "password"}, nil
		},
	}
	service := NewUserService(mock)

	_, err := service.Login("kadirhanmeral@example.com", "wrongpassword")
	if err == nil || err.Error() != "invalid credentials" {
		t.Errorf("Expected 'invalid credentials' error, got %v", err)
	}
}
