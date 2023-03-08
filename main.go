package main

import (
	"fmt"
	"intern_BCC/database"
	"intern_BCC/handler"
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

	customerHandler := handler.NewCustomerHandler(&customerRepo)
	spaceHandler := handler.NewSpaceHandler(&spaceRepo)
	ownerHandler := handler.NewOwnerHandler(&ownerRepo)

	// r.GET("/", ...)

	//---- Login dan Register ----
	r.POST("/login", customerHandler.Login)
	r.POST("/register/customer", customerHandler.CreateCustomer)
	r.POST("/login/owner", ownerHandler.Login)

	//---- Keperluan Cek Database ----
	r.GET("/customers", customerHandler.GetAllCustomer)
	r.GET("/customer/:id", customerHandler.GetCustomerByID)
	r.DELETE("/customer/:id", customerHandler.DeleteCustomerByID)
	r.GET("/owners", ownerHandler.GetAllOwner)
	r.GET("/owner/:id", ownerHandler.GetOwnerByID)
	r.DELETE("/owner/:id", ownerHandler.DeleteOwnerByID)

	//---- Dashboard Customer, Memilih Working Space ----
	r.GET("/spaces", spaceHandler.GetAllSpace)
	r.GET("/spaces/find", spaceHandler.GetSpaceByParam)
	// r.GET("/space/:id", spaceHandler.GetSpaceByID)

	//---- Pemesanan ----

	//---- Riwayat Pesanan ----
	// r.GET("/orders", ...)
	// r.GET("/order/:id", ...)
	// r.POST("/review", ...)

	//---- Update Data Space ----
	// r.POST("/login/owner", ...)
	// r.GET("/owner", ...)  berisi data semua space milik owner A
	// r.GET("/owner/:id", ...)
	// r.POST("/owner", ...)

	//---- Admin ----
	r.POST("/owner", ownerHandler.CreateOwner)
	r.POST("/space", spaceHandler.CreateSpace)
	// r.DELETE("/place/:id", placeHandler.DeletePlaceByID)

	r.Run(":" + port)
}
