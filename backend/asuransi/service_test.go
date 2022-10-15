package asuransi

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/stretchr/testify/assert"
)

func TestGetAsuransiService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := GetAsuransiRequest{
		Judul : "Asuransi Kesehatan",
	}

	asuransi, status, err := service.GetAsuransi(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, asuransi)

	
	req = GetAsuransiRequest{
		Judul: "Asuransi Mobil",
	}

	asuransi, status, err = service.GetAsuransi(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, asuransi)

}

func TestUpdateAsuransiService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	asuransi := UpdateAsuransiRequest{
		Judul: "Asuransi Kesehatan",
		Premi: 300000,
		UangPertanggungan: 100000000,
		Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
		Syarat: "Minimal 17 Tahun",
		Foto: "",
	}
	res, status, err := service.UpdateAsuransi(asuransi)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Asuransi Kesehatan")

	asuransi = UpdateAsuransiRequest{
		Judul: "Asuransi Mobil",
		Premi: 200000,
		UangPertanggungan: 200000000,
		Deskripsi: "Asuransi kesehatan yang dikeluarkan oleh AsuransiKu",
		Syarat: "Minimal 18 Tahun",
		Foto: "",
	}
	res, status, err = service.UpdateAsuransi(asuransi)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, res)
	assert.Equal(t, res.Judul, "Asuransi Mobil")

	
	asuransi = UpdateAsuransiRequest{
		Judul: "",
		Premi: 0,
		UangPertanggungan: 0,
		Deskripsi: "",
		Syarat: "",
		Foto: "",
	}
	res, status, err = service.UpdateAsuransi(asuransi)
	assert.Equal(t, err.Error(), errors.New("Asuransi not Found").Error())
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.Equal(t, res, model.Asuransi{})

}
