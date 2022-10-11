package managelink

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bagasalim/simas/auth"
	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Link_Handler(t *testing.T) {
	type responseLoginData struct {
		Name     string `json:"name"`
		Role     int8   `json:"role"`
		Username string `json:"username"`
	}
	type resposeLogin struct {
		Data  responseLoginData `json:"data"`
		Token string            `json:"token"`
	}

	db := newTestDB(t)
	repoUser := auth.NewRepository(db)

	repoService := auth.NewService(repoUser)
	repoHandler := auth.NewHandler(repoService)

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/login", repoHandler.Login)
	payload := `{"username": "CS01", "password":"123456"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.Use(middleware.IsCS)
	manageLinkRoute.GET("/getlink", handler.GetLink)
	req, err = http.NewRequest("GET", "/getlink?linktype=wa", nil)
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type ResponseMessage struct {
		Message string
	}
	responseMessage := ResponseMessage{}
	assert.Equal(t, 500, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	req, err = http.NewRequest("GET", "/getlink?linktype=", nil)
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseMessage = ResponseMessage{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	req, err = http.NewRequest("GET", "/getlink?linktype=WA", nil)
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	type DataResponse struct {
		LinkValue string `json:"linkvalue"`
	}
	type responseSuccess struct {
		Message string     `json:"message"`
		Data    model.Link `json:"data"`
	}
	res := responseSuccess{}

	assert.Equal(t, 200, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

}
func Test_Handler_Update_Link(t *testing.T) {
	type responseLoginData struct {
		Name     string `json:"name"`
		Role     int8   `json:"role"`
		Username string `json:"username"`
	}
	type resposeLogin struct {
		Data  responseLoginData `json:"data"`
		Token string            `json:"token"`
	}

	db := newTestDB(t)
	repoUser := auth.NewRepository(db)

	repoService := auth.NewService(repoUser)
	repoHandler := auth.NewHandler(repoService)

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/login", repoHandler.Login)
	payload := `{"username": "CS01", "password":"123456"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	data := resposeLogin{}
	r.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &data)

	manageLinkRoute := r.Group("")
	middleware := custom.MiddleWare{}
	manageLinkRoute.Use(middleware.Auth)
	manageLinkRoute.Use(middleware.IsCS)
	manageLinkRoute.POST("/update-link", handler.UpdateLink)

	req, err = http.NewRequest("POST", "/update-link", nil)
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type responseErrorValidation struct {
		Error []string `json:"error"`
	}
	// fmt.Println(string(w.Body.Bytes()[:]))
	responseMessage := responseErrorValidation{}
	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	responseMessage = responseErrorValidation{}
	req, err = http.NewRequest("POST", "/update-link?linktype=wa", nil)
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	// fmt.Println(string(w.Body.Bytes()[:]))
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	payload = `{"linktype": "WA", "linkvalue":"test", "updatedby":"rema"}`
	req, err = http.NewRequest("POST", "/update-link?linktype=wa", strings.NewReader(payload))
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	// fmt.Println( w.Code,string(w.Body.Bytes()[:]))
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &responseMessage))

	payload = `{"linktype": "WA", "linkvalue":"test", "updatedby":"rema"}`
	req, err = http.NewRequest("POST", "/update-link?linktype=WA", strings.NewReader(payload))
	req.Header.Set("Authorization", data.Token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	type responseSuccess struct {
		Message string     `json:"message"`
		Data    model.Link `json:"data"`
	}
	res := responseSuccess{}
	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Code, string(w.Body.Bytes()[:]))
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))

}
