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
	usersGroup.Use(middleware.RBACMiddleware([]string{"admin"}))
	// Untuk membuat admin pertama kami mematikan atau mengkomen jwt
	// untuk yang pertama kalinya, middleware.AuthMiddleware(cfg.JWTSecret) dan middleware.RBACMiddleware([]string{"admin"})
	// 	{
	//     "code": 200,
	//     "data": {
	//         "ID": "ea7ac4da-63fe-4c67-ab4f-81e529ed7167",
	//         "Name": "Randi Imam",
	//         "Email": "randi@mail.com",
	//         "Password": "$2a$14....., telah di hash
	//         "Role": "admin",
	//         "CreatedAt": "2025-12-09T04:26:57.5702749+07:00",
	//         "UpdatedAt": "2025-12-09T04:26:57.5702749+07:00",
	//         "DeletedAt": null
	//     },
	//     "message": "Cashier created",
	//     "status": "success"
	// }
	usersGroup.POST("/", user.CreateUserHandler(userService))
	usersGroup.GET("/", user.GetAllCashiersHandler(userService))
	usersGroup.GET("/:id", user.GetCashierByIDHandler(userService))
	usersGroup.PUT("/:id", user.UpdateCashierHandler(userService))
	usersGroup.DELETE("/:id", user.DeleteCashierHandler(userService))

	// Sale Orders
	salesGroup := api.Group("/sale-orders")
	salesGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	salesGroup.Use(middleware.RBACMiddleware([]string{"admin", "cashier"}))
	salesGroup.POST("", saleorder.CreateSaleOrderHandler(saleOrderService))
	salesGroup.GET("", saleorder.GetAllSaleOrdersHandler(saleOrderService))
	salesGroup.GET("/:id", saleorder.GetSaleOrderByIDHandler(saleOrderService))
	salesGroup.PUT("/:id", saleorder.UpdateSaleOrderHandler(saleOrderService))
	salesGroup.DELETE("/:id", saleorder.DeleteSaleOrderHandler(saleOrderService))

	log.Printf("Server running on port %s", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
