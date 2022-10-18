package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

const (
	login           = "/login"
	createuser      = "/create-user"
	updateLastLogin = "/updatelastlogin"
)

type responseLoginData struct {
	Name     string `json:"name"`
	Role     int8   `json:"role"`
	Username string `json:"username"`
}
type resposeLogin struct {
	Data  responseLoginData `json:"data"`
	Token string            `json:"token"`
}
type responseSuccess struct {
	Message string     `json:"message"`
	Data    model.Link `json:"data"`
}

func initialRepoAuth(t *testing.T) *Handler {
	db := newTestDB(t)
	repoUser := NewRepository(db)

	repoService := NewService(repoUser)
	repoHandler := NewHandler(repoService)
	return repoHandler
}

func initialRepo(t *testing.T) *Handler {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return handler
}

type ResponseMessage struct {
	Message string
}

func getToken(t *testing.T) string {
	handler := initialRepoAuth(t)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/login", handler.Login)
	payload := `{"username": "cindu", "password":"123456"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)
	return data.Token

}

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
	r.POST(login, handler.Login)
	payload := `{"username": "remasertu", "password":"123456"}`
	req, err := http.NewRequest("POST", login, strings.NewReader(payload))
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
	req, err = http.NewRequest("POST", login, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var errValid responseErrorValidation
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &errValid))

	//validation
	payload = `{"username": "remasertu", "password":"12345"}`
	req, err = http.NewRequest("POST", login, strings.NewReader(payload))
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
	r.POST(createuser, handler.CreateUser)

	payload := `{"username": "remasertu", "password":"123456", "name":"rema"}`
	req, err := http.NewRequest("POST", createuser, strings.NewReader(payload))
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
	req, err = http.NewRequest("POST", createuser, strings.NewReader(payload))
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
	req, err = http.NewRequest("POST", createuser, strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	fmt.Println("res", w.Code, string(w.Body.Bytes()[:]))
	assert.NotEqual(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &success))
}

func TestUpdateLastLoginHandler(t *testing.T) {
	token := getToken(t)
	handler := initialRepo(t)
	fmt.Println(token)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.POST(updateLastLogin, handler.UpdateLastLogin)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}

	//success
	payload := `{"username": "cindu", "lastlogin":"2022-10-18T10:12:07.000Z"}`
	req, _ := http.NewRequest("POST", updateLastLogin, strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Code, string(w.Body.Bytes()[:]))
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

	//fail wrong input
	payload = `{"username": "cindu", "lastlogin":"2022-10-1810:12:07.000Z"}`
	req, _ = http.NewRequest("POST", updateLastLogin, strings.NewReader(payload))
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseMessage := responseErrorValidation{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))
}
