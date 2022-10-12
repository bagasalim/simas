package zoomhistory

import (
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
	err = db.AutoMigrate(&model.User{}, &model.Riwayat{})
	assert.NoError(t, err)

	return db
}

func TestCreateUser(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	// repo.
	Riwayat := model.Riwayat{
		Nama:       "cayo",
		Email:      "cayo@gmail.com",
		Kategori:   "Kredit",
		Keterangan: "Gatau",
	}
	// task := "task 1"
	res, err := repo.AddUser(Riwayat)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	Riwayat = model.Riwayat{
		Nama: "cayo",
	}
	_, err = repo.AddUser(Riwayat)
	// fmt.Println(err())
	assert.NotNil(t, err)

}
