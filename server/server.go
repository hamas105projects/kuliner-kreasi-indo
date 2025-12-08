package server

import (
	"kuliner-kreasi-indo/config"
	"kuliner-kreasi-indo/middleware"
	"kuliner-kreasi-indo/modules/auth"
	"kuliner-kreasi-indo/modules/saleorder"
	"kuliner-kreasi-indo/modules/user"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.LoadConfig()
	db := cfg.DB // harus sudah connect ke DB yang tabelnya manual

	r := gin.Default()

	authService := &auth.AuthService{
		DB:        db,
		JWTSecret: cfg.JWTSecret,
		ExpireHrs: cfg.JWTExpire,
	}
	userService := &user.UserService{DB: db}
	saleOrderRepo := &saleorder.SaleOrderRepository{DB: db}
	saleOrderService := &saleorder.SaleOrderService{Repo: saleOrderRepo}

	api := r.Group("/api")
	api.POST("/auth/login", auth.LoginHandler(authService))

	// Users
	usersGroup := api.Group("/users")
	usersGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	usersGroup.Use(middleware.RBACMiddleware([]string{"owner"}))
	usersGroup.POST("/cashier", user.CreateCashierHandler(userService))
	usersGroup.GET("/cashier", user.GetAllCashiersHandler(userService))
	usersGroup.GET("/cashier/:id", user.GetCashierByIDHandler(userService))
	usersGroup.PUT("/cashier/:id", user.UpdateCashierHandler(userService))
	usersGroup.DELETE("/cashier/:id", user.DeleteCashierHandler(userService))

	// Sale Orders
	salesGroup := api.Group("/sale-orders")
	salesGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	salesGroup.Use(middleware.RBACMiddleware([]string{"owner", "cashier"}))
	salesGroup.POST("", saleorder.CreateSaleOrderHandler(saleOrderService))
	salesGroup.GET("", saleorder.GetAllSaleOrdersHandler(saleOrderService))
	salesGroup.GET("/:id", saleorder.GetSaleOrderByIDHandler(saleOrderService))
	salesGroup.PUT("/:id", saleorder.UpdateSaleOrderHandler(saleOrderService))
	salesGroup.DELETE("/:id", saleorder.DeleteSaleOrderHandler(saleOrderService))

	log.Printf("Server running on port %s", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
