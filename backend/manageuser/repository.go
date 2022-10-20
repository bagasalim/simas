package manageuser

import (
	"fmt"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]model.User, error)
	UpdateUser() ([]model.User, error) //tinggal diubah
	DeleteUser() ([]model.User, error) //tinggal diubah
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUser() ([]model.User, error) {
	var User []model.User
	res := r.db.Find(&User)
	if res.Error != nil {
		return nil, res.Error
	}

	return User, nil
}

func (r *repository) UpdateUser() ([]model.User, error) {
	fmt.Println("UpdateRepo")
	var User []model.User
	res := r.db.Find(&User)
	if res.Error != nil {
		return nil, res.Error
	}

	return User, nil
	//tinggal diubah
}
func (r *repository) DeleteUser() ([]model.User, error) {
	fmt.Println("DeleteRepo")
	var User []model.User
	res := r.db.Find(&User)
	if res.Error != nil {
		return nil, res.Error
	}

	return User, nil
	//tinggal diubah
}
