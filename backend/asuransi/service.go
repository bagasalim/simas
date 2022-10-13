package asuransi

import (
	"errors"
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetAsuransi(data GetAsuransiRequest) (model.Asuransi, int, error)
	UpdateAsuransi(data UpdateAsuransiRequest) (model.Asuransi, int,  error)
}

type service struct {
	repo AsuransiRepository
}

func NewService(repo AsuransiRepository) *service {
	return &service{repo}
}

func (s *service) GetAsuransi(data GetAsuransiRequest) (model.Asuransi, int, error) {
	asuransi, err := s.repo.GetAsuransi(data.Judul)
	if err != nil {
		return model.Asuransi{}, http.StatusInternalServerError, err
	}

	return asuransi, http.StatusOK, nil
}

func (s *service) UpdateAsuransi(data UpdateAsuransiRequest) (model.Asuransi, int, error) {
	found, err := s.repo.GetAsuransi(data.Judul)
	if err != nil {
		return model.Asuransi{}, http.StatusInternalServerError, err
	}
	if err == nil && found.Judul == "" {
		return model.Asuransi{}, http.StatusBadRequest, errors.New("insurance not found")
	}

	Asuransi := model.Asuransi{
		Judul: data.Judul,
	}

	res, err := s.repo.UpdateAsuransi(Asuransi)
	if err != nil {
		return model.Asuransi{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}