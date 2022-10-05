package managelink

import (
	"errors"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type LinkRepository interface {
	GetLink(linktype string) (model.Link, error)
	UpdateLink(link model.Link) (model.Link, error)
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

func (r *repository) UpdateLink(link model.Link) (model.Link, error) {
	res := r.db.Where("link_type=?", link.Link_Type).Updates(model.Link{
		Link_Value: link.Link_Value,
		Updated_By: link.Updated_By,
	})
	if res.Error != nil {
		return model.Link{}, res.Error
	}

	return link, nil
}
