package auth

import (
	"errors"
	"time"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUser(username string) (model.User, error)
	AddUser(user model.User) (model.User, error)
	AddLastLogin(username string, lastlogin time.Time) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindUser(username string) (model.User, error) {
	var User model.User
	if err := r.db.Where("username = ?", username).First(&User).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("not found")
		}
		return model.User{}, err
	}
	return User, nil
}
func (r *repository) AddUser(user model.User) (model.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}
func (r *repository) AddLastLogin(username string, lastlogin time.Time) (model.User, error) {
	var User model.User
	res := r.db.Model(&User).Where("username=?", username).Updates(model.User{
		LastLogin: lastlogin,
	})

	if res.Error != nil {
		return model.User{}, res.Error
	}

	return User, nil
}
