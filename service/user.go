package service

import (
	"e-commerce/model"
	"e-commerce/repository"
	"errors"
)

type UserService interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
	GetUser(id int64) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email and password are required")
	}

	existing, _ := s.repo.FindByEmail(user.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *userService) GetUser(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}
