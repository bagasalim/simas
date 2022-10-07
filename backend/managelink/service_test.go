package managelink

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/stretchr/testify/assert"
)

func TestGetLinkService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	//Get WA
	req := GetLinkRequest{
		LinkType: "WA",
	}

	link, status, err := service.GetLink(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, link)

	//Get Zoom
	req = GetLinkRequest{
		LinkType: "Zoom",
	}

	link, status, err = service.GetLink(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
	assert.NotNil(t, link)

}

func TestUpdateLinkService(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	//Link WA
	link := model.Link{
		LinkType:  "WA",
		LinkValue: "Ini Link WA Update",
		UpdatedBy: "System",
	}
	res, err := repo.UpdateLink(link)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link WA Update")

	//Link Zoom
	link = model.Link{
		LinkType:  "Zoom",
		LinkValue: "Ini Link Zoom Update",
		UpdatedBy: "System",
	}
	res, err = repo.UpdateLink(link)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.LinkValue, "Ini Link Zoom Update")

	//No Link
	link = model.Link{
		LinkType:  "No Link",
		LinkValue: "No Link",
		UpdatedBy: "System",
	}
	res, err = repo.UpdateLink(link)
	assert.Equal(t, err.Error(), errors.New("wrong link type").Error())
	assert.Equal(t, res, model.Link{})

}
