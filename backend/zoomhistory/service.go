package zoomhistory

import (
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	CreateZoomHistory(data ZoomHistoryRequest) (model.Zoom, int, error)
}

type service struct {
	repo ZoomRepository
}

func NewService(repo ZoomRepository) *service {
	return &service{repo}
}

func (s *service) CreateZoomHistory(data ZoomHistoryRequest) (model.Zoom, int, error) {
	Zoom := model.Zoom{
		Username: data.Username,
		Email:    data.Email,
		Kategori: data.Kategori,
		Keluhan:  data.Keluhan,
	}
	res, err := s.repo.AddUser(Zoom)
	if err != nil {
		return model.Zoom{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
