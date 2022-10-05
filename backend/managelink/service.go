package managelink

import (
	"net/http"

	"github.com/bagasalim/simas/model"
)

type Service interface {
	GetLink(data GetLinkRequest) (model.Link, int, error)
}

type service struct {
	repo LinkRepository
}

func NewService(repo LinkRepository) *service {
	return &service{repo}
}

func (s *service) GetLink(data GetLinkRequest) (model.Link, int, error) {

	link, err := s.repo.GetLink(data.Link_Type)
	if err != nil {
		return model.Link{}, http.StatusInternalServerError, err
	}

	return link, http.StatusOK, nil
}
