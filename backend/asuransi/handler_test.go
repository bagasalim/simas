package asuransi

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/bagasalim/simas/model"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetTodoHandler(t *testing.T) {
// 	db := newTestDB(t)
// 	repo := NewRepository(db)
// 	service := NewService(repo)
// 	handler := NewHandler(service)

// 	db.Create(&model.Asuransi{
// 		Judul: "Asuransi Kesehatan",
// 	})

// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()

// 	r.GET("/", handler.GetAsuransi)
// 	req, err := http.NewRequest("GET", "/", nil)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, req)

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	type response struct {
// 		Judul string        `json:"judul"`
// 		Data    []model.Asuransi `json:"data"`
// 	}

// 	var res response

// 	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
// 	assert.Equal(t, "success", res.Judul)
// 	assert.Equal(t, "Asuransi Kesehatan", res.Data[0].Judul)

// }
