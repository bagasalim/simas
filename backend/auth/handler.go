package auth

import (
	"net/http"

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

func (h *Handler) SendOTP(c *gin.Context){
	
	var req SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := custom.ParseError(err)
		if messageErr == nil {
			messageErr = []string{"Input data not suitable"}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	token, email,  status, err := h.Service.SetOtp(req.Username)
	if err != nil{
		mes := err.Error()
		if mes == "Username not found"{
			mes = "Fail to send otp"
		}
		c.JSON(status, gin.H{
			"message":mes,
		})
		return 
	}
    to := []string{email}
    cc := []string{}
    subject := "Simas Contact OTP Kode"
    message := "OTP Kode:"+token
	err = custom.SendMail(to, cc, subject, message)
	resp :="sukses"
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": resp,
	})
	
}