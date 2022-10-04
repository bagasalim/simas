package auth

import (
	"errors"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUser(username string) (model.User, error)
	addUser(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindUser(username string) (model.User, error) {
	var User model.User
	res := r.db.Where("username = ?", username).First(&User)
	// fmt.Println("findUser", res)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return model.User{}, errors.New("Username or Password is wrong")
		}
		return model.User{}, res.Error
	}
	return User, nil
}
func (r *repository) addUser(user model.User) (model.User, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return model.User{}, res.Error
	}

	return user, nil
}