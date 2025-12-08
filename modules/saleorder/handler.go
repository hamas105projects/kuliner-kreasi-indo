package saleorder

import (
	"kuliner-kreasi-indo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSaleOrderHandler(service *SaleOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Items []SaleOrderItem `json:"items"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		cashierID := c.GetString("user_id")
		order, err := service.CreateSaleOrder(cashierID, req.Items)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.SuccessResponse(c, http.StatusOK, "Sale order created", order)
	}
}

func GetAllSaleOrdersHandler(service *SaleOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, limit := utils.GetPagination(c)
		offset := (page - 1) * limit

		orders, err := service.GetAllSaleOrders(offset, limit)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "List sale orders", orders)
	}
}

func GetSaleOrderByIDHandler(service *SaleOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		order, err := service.GetSaleOrderByID(id)
		if err != nil {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Sale order detail", order)
	}
}

func UpdateSaleOrderHandler(service *SaleOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req struct {
			Items []SaleOrderItem `json:"items"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		order, err := service.UpdateSaleOrder(id, req.Items)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Sale order updated", order)
	}
}

func DeleteSaleOrderHandler(service *SaleOrderService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := service.DeleteSaleOrder(id); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		utils.SuccessResponse(c, http.StatusOK, "Sale order deleted", nil)
	}
}
