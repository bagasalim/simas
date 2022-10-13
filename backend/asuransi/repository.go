package asuransi

import (
	"errors"

	"github.com/bagasalim/simas/model"
	"gorm.io/gorm"
)
type AsuransiRepository interface {
	GetAsuransi(judul string) (model.Asuransi, error)
	UpdateAsuransi(asuransi model.Asuransi) (model.Asuransi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAsuransi(judul string) (model.Asuransi, error) {
	var Asuransi model.Asuransi
	if err := r.db.Where("judul = ?", judul).First(&Asuransi).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Asuransi{}, errors.New("insurance not found")
		}
		return model.Asuransi{}, err
	}
	return Asuransi, nil
}

func (r *repository) UpdateAsuransi(asuransi model.Asuransi) (model.Asuransi, error) {
	_, err := r.GetAsuransi(asuransi.Judul)
	if err != nil {
		return model.Asuransi{}, errors.New("wrong insurance input")
	}

	res := r.db.Where("judul=?", asuransi.Judul).Updates(model.Asuransi{
		Judul: asuransi.Judul,
	})
	if res.Error != nil {
		return model.Asuransi{}, res.Error
	}

	return asuransi, nil
}