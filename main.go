package main

import (
	"fmt"
	"intern_BCC/database"
	"intern_BCC/handler"
	"intern_BCC/middleware"
	"intern_BCC/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// remove this comment
func main() {
	fmt.Println("Hello")
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load env file")
	}
	port := os.Getenv("PORT")

	db := database.InitDB()

	if err := database.AutoMigrate(db); err != nil {
		log.Fatalln("Auto migrate error", err)
	}

	customerRepo := repository.NewCustomerRepository(db)
	spaceRepo := repository.NewSpaceRepository(db)
	ownerRepo := repository.NewOwnerRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	customerHandler := handler.NewCustomerHandler(&customerRepo)
	spaceHandler := handler.NewSpaceHandler(&spaceRepo)
	ownerHandler := handler.NewOwnerHandler(&ownerRepo)
	orderHandler := handler.NewOrderHandler(&orderRepo)

	// r.GET("/", ...)

	//---- Login dan Register ----
	r.POST("/login", customerHandler.Login)
	r.POST("/register/customer", customerHandler.CreateCustomer)

	//---- Keperluan Cek Database ----
	r.GET("/customers", customerHandler.GetAllCustomer)
	// r.GET("/customer/:id", customerHandler.GetCustomerByID)
	r.DELETE("/customer/:id", customerHandler.DeleteCustomerByID)
	r.GET("/owners", ownerHandler.GetAllOwner)
	r.GET("/owner/:id", ownerHandler.GetOwnerByID)
	r.DELETE("/owner/:id", ownerHandler.DeleteOwnerByID)

	//---- Memilih Working Space ----
	r.GET("/spaces", spaceHandler.GetAllSpace)
	r.GET("/spaces/find", spaceHandler.GetSpaceByParam)
	r.GET("/space/:id", spaceHandler.GetSpaceByID)
	r.POST("/space/:id", middleware.JwtMiddleware(), orderHandler.CreateOrder)

	//---- Pemesanan ----

	//---- Riwayat Pesanan ----
	// r.GET("/orders", ...)
	// r.GET("/order/:id", ...)
	// r.POST("/review", ...)

	//---- Update Data Space ----
	r.POST("/login/owner", ownerHandler.Login)
	// r.GET("/owner", ...)
	// r.GET("/owner/:id", ...)		ID Space
	// r.POST("/owner/:id", ...)

	//---- Admin ----
	r.POST("/owner", ownerHandler.CreateOwner)
	r.POST("/space", spaceHandler.CreateSpace)
	// r.DELETE("/place/:id", placeHandler.DeletePlaceByID)

	r.Run(":" + port)
}
