package manageuser

import (
	"fmt"
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetUser() ([]model.User, int, error)
	UpdateUser() ([]model.User, int, error) //tinggal di ubah
	DeleteUser() ([]model.User, int, error) //tinggal di ubah
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *service {
	return &service{repo}
}

func (s *service) GetUser() ([]model.User, int, error) {

	user, err := s.repo.GetUser()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
func (s *service) UpdateUser() ([]model.User, int, error) {
	fmt.Println("UpdateService")
	user, err := s.repo.UpdateUser()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
	//tinggal diubah
}
func (s *service) DeleteUser() ([]model.User, int, error) {
	fmt.Println("DeleteService")
	user, err := s.repo.UpdateUser()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
	//tinggal diubah
}
