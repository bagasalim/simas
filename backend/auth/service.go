package auth

import (
	"errors"
	"net/http"

	"github.com/bagasalim/simas/model"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(data LoginRequest) (model.User, int, error)
	CreateAccount(data RegisterRequest) (model.User, int, error)
	UpdateLastLogin(data LastLoginRequest) (model.User, int, error)
}

type service struct {
	repo AuthRepository
}

func NewService(repo AuthRepository) *service {
	return &service{repo}
}
func (s *service) Login(data LoginRequest) (model.User, int, error) {
	username := data.Username
	User, err := s.repo.FindUser(username)

	if err != nil {
		if err.Error() == "not found" {
			return model.User{}, http.StatusUnauthorized, errors.New("username or password is wrong")
		}
		return model.User{}, http.StatusInternalServerError, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errors.New("password is wrong")
	}
	User.Password = ""
	return User, http.StatusOK, nil
}
func (s *service) CreateAccount(data RegisterRequest) (model.User, int, error) {
	found, err := s.repo.FindUser(data.Username)
	if err == nil && found.Name != "" {
		return model.User{}, http.StatusBadRequest, errors.New("duplicate data")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User := model.User{
		Username: data.Username,
		Password: string(passwordHash),
		Name:     data.Name,
		Role:     2,
	}
	res, err := s.repo.AddUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}

func (s *service) UpdateLastLogin(data LastLoginRequest) (model.User, int, error) {
	res, err := s.repo.AddLastLogin(data.Username, data.LastLogin)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
