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
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	_, status, err := h.Service.CreateAccount(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
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
		"data": map[string]any{
			"name":     res.Name,
			"username": res.Username,
			"role":     res.Role,
		},
	})
	return
}
func (h *Handler) Test(c *gin.Context) {
	dataUser, exist := c.Get("user")
	if exist == false {

	}
	fmt.Println(dataUser)
}
