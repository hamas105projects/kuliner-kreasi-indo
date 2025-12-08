package user

import (
	"kuliner-kreasi-indo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCashierHandler(userService *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		user, err := userService.CreateCashier(req.Name, req.Email, req.Password)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Cashier created", user)
	}
}

func GetAllCashiersHandler(userService *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, limit := utils.GetPagination(c)
		offset := (page - 1) * limit

		users, err := userService.GetAllCashiers(offset, limit)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "List cashiers", users)
	}
}

func GetCashierByIDHandler(userService *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := userService.GetCashierByID(id)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if user.ID.String() == "" {
			utils.ErrorResponse(c, http.StatusNotFound, "Cashier not found")
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Cashier detail", user)
	}
}

func UpdateCashierHandler(userService *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		user, err := userService.UpdateCashier(id, req.Name, req.Email, req.Password)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Cashier updated", user)
	}
}

func DeleteCashierHandler(userService *UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := userService.DeleteCashier(id); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Cashier deleted", nil)
	}
}
