package asuransi

import (
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
	var req GetAsuransiRequest
	judul := c.Query("judul")
	if judul == "" {
		messageErr := []string{"Input data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	req = GetAsuransiRequest{Judul: judul}
	asuransi, status, err := h.Service.GetAsuransi(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data": asuransi,
	})
}

func (h *Handler) UpdateAsuransi(c *gin.Context) {
	

	var req UpdateAsuransiRequest
	judul := c.Query("judul")

	if judul == "" {
		messageErr := []string{"Param data not suitable"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	reqFix := UpdateAsuransiRequest{
		Judul:  judul,
		Premi: req.Premi,
		UangPertanggungan: req.UangPertanggungan,
		Deskripsi: req.Deskripsi,
		Syarat: req.Syarat,
		Foto: req.Foto,
	}
	asuransi, status, err := h.Service.UpdateAsuransi(reqFix)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    asuransi,
	})

}
func (h *Handler) GetAsuransiRequest(c *gin.Context) {
	judul, _ := c.Params.Get("judul")
	req := GetAsuransiRequest{Judul: judul}
	asuransi, status, err := h.Service.GetAsuransi(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    asuransi,
	})
}
