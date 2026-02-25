package controller

import (
	"net/http"

	"github.com/cecepwahyu/rest-api-go/internal/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service service.LoginService
}

func NewLoginController(service service.LoginService) *LoginController {
	return &LoginController{service}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lc *LoginController) Login(c *gin.Context) {

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	result, err := lc.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
