package asuransi

import (
	"errors"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.Asuransi{}, &model.User{})
	assert.NoError(t, err)

	asuransi := []model.Asuransi{
		{
			Judul: "Asuransi Kesehatan",
			Premi: 300000,
			UangPertanggungan: 100000000,
			Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
			Syarat: "Minimal 17 Tahun",
			Foto: "",
		},
		{
			Judul: "Asuransi Mobil",
			Premi: 200000,
			UangPertanggungan: 200000000,
			Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
			Syarat: "Minimal 18 Tahun",
			Foto: "",
		},
	}
	err = db.Create(&asuransi).Error
	assert.NoError(t, err)

	dataUser := []model.User{
		{
			Username: "CS01",
			Password: "$2a$10$BQHCjmHmEsFGJXCGWm7et.2lvVPecg0ibhFd/tgOCCCncTu5ieiA.",
			Name:     "Customer Service",
			Role:     2,
		},
	}
	db.Create(&dataUser)

	return db
}

func TestGetAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetAsuransi("Asuransi Kesehatan")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Ini Asuransi Kesehatan")

	res, err = repo.GetAsuransi("Asuransi Mobil")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Ini Asuransi Mobil")

	res, err = repo.GetAsuransi("No Link")
	assert.Equal(t, err.Error(), errors.New("Asuransi Tidak Ditemukan").Error())
	assert.Equal(t, res, model.Asuransi{})
}

func TestUpdateAsuransi(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	asuransi := model.Asuransi{
		Judul: "Asuransi Kesehatan",
		Premi: 300000,
		UangPertanggungan: 100000000,
		Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
		Syarat: "Minimal 17 Tahun",
		Foto: "",
	}
	res, err := repo.UpdateAsuransi(asuransi)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Ini Asuransi Update")

	asuransi = model.Asuransi{
		Judul: "Asuransi Mobil",
		Premi: 200000,
		UangPertanggungan: 200000000,
		Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
		Syarat: "Minimal 18 Tahun",
		Foto: "",
	}
	res, err = repo.UpdateAsuransi(asuransi)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Ini Asuransi Update")

	asuransi = model.Asuransi{
		Judul: "",
		Premi: 0,
		UangPertanggungan: 0,
		Deskripsi: "",
		Syarat: "",
		Foto: "",
	}
	res, err = repo.UpdateAsuransi(asuransi)
	assert.Equal(t, err.Error(), errors.New("wrong insurance").Error())
	assert.Equal(t, res, model.Asuransi{})

}
