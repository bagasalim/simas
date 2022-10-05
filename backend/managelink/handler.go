package managelink

import (
	"fmt"
	"net/http"

	"github.com/bagasalim/simas/custom"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetLink(c *gin.Context) {
	dataUser, exist := c.Get("user")
	if !exist {
		return
	}

	var req GetLinkRequest
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	messageErr := custom.ParseError(err)
	// 	if messageErr == nil {
	// 		messageErr = []string{"Input data not suitable"}
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
	// 	return
	// }
	linktype := c.Query("linktype")
	_ = linktype
	if linktype == "" {
		messageErr := []string{"Input data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	req = GetLinkRequest{LinkType: linktype}
	link, status, err := h.Service.GetLink(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    link,
	})

	fmt.Println(dataUser)
}

func (h *Handler) UpdateLink(c *gin.Context) {
	dataUser, exist := c.Get("user")
	if !exist {
		return
	}

	var req UpdateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	link, status, err := h.Service.UpdateLink(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    link,
	})

	fmt.Println(dataUser)
}
