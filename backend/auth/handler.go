package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bagasalim/simas/custom"
	_ "github.com/bagasalim/simas/custom"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}
func (h *Handler) CreateUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, status, err := h.Service.CreateAccount(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    res,
	})
	return
}
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	res, status, err := h.Service.Login(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	// custom.DataJWT{}
	token, err := custom.GenerateJWT(res.Username, res.Name, res.Role)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println(time.Now().Add(60 * time.Minute))
	c.JSON(status, gin.H{
		"token": token,
		"user":  res,
	})
	return
}
func (h *Handler) Test(c *gin.Context) {
	// auth := c.Request.Header["Authorization"]
	// fmt.Println(auth)
	// if len(auth) == 0 {
	// 	fmt.Println("no auth")
	// 	return
	// }
	// token := auth[0]
	// fmt.Println("token", token)
	// dataUser, err := custom.ClaimToken(token)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// c.Set("user", dataUser)
	// dataUser := custom.DataJWT{}
	dataUser, exist := c.Get("user")
	if exist == false {

	}
	fmt.Println(dataUser)
}
