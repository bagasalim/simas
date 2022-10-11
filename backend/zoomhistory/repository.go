package zoomhistory

import (
	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type ZoomRepository interface {
	AddUser(zoom model.Zoom) (model.Zoom, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) AddUser(zoom model.Zoom) (model.Zoom, error) {
	res := r.db.Create(&zoom)
	if res.Error != nil {
		return model.Zoom{}, res.Error
	}

	return zoom, nil
}
