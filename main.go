package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mr-amirfazel/shopping-basket/db"
	"github.com/mr-amirfazel/shopping-basket/handlers"
	"github.com/mr-amirfazel/shopping-basket/models"
)

func main() {
	db.InitDB()
	defer db.CloseDB() 

	
	db.DB.AutoMigrate(&models.Basket{})


	db.DB.AutoMigrate(&models.User{})



	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping!")
	})

	// Routes
	basketRoute := e.Group("/basket")
	basketRoute.GET("/p", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping!")
	})
	basketRoute.GET("/", handlers.GetBaskets)
	basketRoute.POST("/", handlers.CreateBasket)
	basketRoute.GET("/:id", handlers.GetBasketByID)   
	basketRoute.PATCH("/:id", handlers.UpdateBasket) 
	basketRoute.DELETE("/:id", handlers.DeleteBasket)
	
	
	userRoute := e.Group("/user")
	userRoute.GET("/", handlers.GetAllUsers)
    userRoute.GET("/:id", handlers.GetUserByID)
    userRoute.POST("/", handlers.CreateUser)
    userRoute.POST("/login", handlers.LoginUser)
    userRoute.DELETE("/:id", handlers.DeleteUserByID)
    userRoute.PATCH("/:id", handlers.ChangePassword)
	log.Fatal(e.Start("0.0.0.0:8080"))
}
