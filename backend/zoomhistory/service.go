package zoomhistory

import (
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	CreateZoomHistory(data ZoomHistoryRequest) (model.Riwayat, int, error)
}

type service struct {
	repo ZoomRepository
}

func NewService(repo ZoomRepository) *service {
	return &service{repo}
}

func (s *service) CreateZoomHistory(data ZoomHistoryRequest) (model.Riwayat, int, error) {
	Riwayat := model.Riwayat{
		Nama:       data.Nama,
		Email:      data.Email,
		Kategori:   data.Kategori,
		Keterangan: data.Keterangan,
	}
	res, err := s.repo.AddUser(Riwayat)
	if err != nil {
		return model.Riwayat{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
