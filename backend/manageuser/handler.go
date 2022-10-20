package manageuser

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetUser(c *gin.Context) {
	todos, status, err := h.Service.GetUser()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	fmt.Println("UpdateHandler")
	todos, status, err := h.Service.GetUser()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
	//tinggal diubah
}
func (h *Handler) DeleteUser(c *gin.Context) {
	fmt.Println("DeleteHandler")
	todos, status, err := h.Service.GetUser()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
	//tinggal diubah
}
