package auth

import (
	"kuliner-kreasi-indo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(authService *AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		token, err := authService.Login(req.Email, req.Password)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
	}
}
