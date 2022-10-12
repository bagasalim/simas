package zoomhistory

import (
	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)

type ZoomRepository interface {
	AddUser(Riwayat model.Riwayat) (model.Riwayat, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) AddUser(Riwayat model.Riwayat) (model.Riwayat, error) {
	res := r.db.Create(&Riwayat)
	if res.Error != nil {
		return model.Riwayat{}, res.Error
	}

	return Riwayat, nil
}
