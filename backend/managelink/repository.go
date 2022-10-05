package managelink

import (
	"errors"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type LinkRepository interface {
	GetLink(linktype string) (model.Link, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLink(linktype string) (model.Link, error) {
	var Link model.Link
	if err := r.db.Where("link_type = ?", linktype).First(&Link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Link{}, errors.New("link not found")
		}
		return model.Link{}, err
	}
	return Link, nil
}
