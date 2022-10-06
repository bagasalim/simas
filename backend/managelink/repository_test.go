package managelink

import (
	"errors"
	"testing"

	_ "errors"

	"github.com/bagasalim/simas/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
	err = db.AutoMigrate(&model.Link{})
	assert.NoError(t, err)

	link := model.Link{
		LinkType: "WA",
	}
	err = db.Create(&link).Error
	assert.NoError(t, err)

	return db
}

func TestGetLink(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	res, err := repo.GetLink("WA")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	res1, err1 := repo.GetLink("No Link")
	assert.Equal(t, err1.Error(), errors.New("link not found").Error())
	assert.Equal(t, res1, model.Link{})
}

func TestUpdateLink(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	link := model.Link{
		LinkType:  "WA",
		LinkValue: "Ini Link WA",
		UpdatedBy: "System",
	}
	res, err := repo.UpdateLink(link)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	link2 := model.Link{
		LinkType:  "No Link",
		LinkValue: "No Link",
		UpdatedBy: "System",
	}
	res2, err2 := repo.UpdateLink(link2)
	assert.Equal(t, err2.Error(), errors.New("wrong link type").Error())
	assert.Equal(t, res2, model.Link{})

}
