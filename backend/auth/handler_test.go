package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	User := model.User{
		Username: "remasertu",
		Password: string(passwordHash),
		Name:     "rema",
		Role:     2,
	}
	// task := "task 1"
	repo.AddUser(User)
	service := NewService(repo)
	handler := NewHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/login", handler.Login)
	payload := `{"username": "remasertu", "password":"123456"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	type responseErrorValidation struct {
		Error []string `json:"error"`
	}
	type responseSuccess struct {
		Data  map[string]any `json:"data"`
		Token string         `json:"token"`
	}
	type responseError struct {
		Message string `json:"message"`
	}
	var success responseSuccess

	// assert.Equal(t, )
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))

	//validation
	payload = ``
	req, err = http.NewRequest("POST", "/login", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var errValid responseErrorValidation
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	//validation
	payload = `{"username": "remasertu", "password":"12345"}`
	req, err = http.NewRequest("POST", "/login", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var errorMessage responseError
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errorMessage))

}
func TestCreateUser(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	// passwordHash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	// User := model.User{
	// 	Username: "remasertu",
	// 	Password: string(passwordHash),
	// 	Name:     "rema",
	// 	Role:     2,
	// }
	// task := "task 1"
	// repo.AddUser(User)
	service := NewService(repo)
	handler := NewHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/create-user", handler.CreateUser)

	payload := `{"username": "remasertu", "password":"123456", "name":"rema"}`
	req, err := http.NewRequest("POST", "/create-user", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	type responseMess struct {
		Message string `json:"message"`
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	var success responseMess
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))

	// error validation
	payload = ``
	req, err = http.NewRequest("POST", "/create-user", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}
	var errValid responseErrorValidation
	// fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	//error duplicate
	payload = `{"username": "remasertu", "password":"123456", "name":"rema"}`
	req, err = http.NewRequest("POST", "/create-user", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	assert.NotEqual(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))
}
