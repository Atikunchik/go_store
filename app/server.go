package main

import (
	"GolangwithFrame/app/controller"
	"GolangwithFrame/src/app/service"
	"GolangwithFrame/src/infrastructure/cache"
	"GolangwithFrame/src/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	_ "github.com/lib/pq"
)

var (
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main() {

	Cache := cache.NewRedisCache("redis:6379", 0, 60)
	Repository := repository.NewRepository()
	Service := service.New(Repository)
	controller := Controller.New(*Service, Cache)

	defer Repository.CloseDB()
	server := gin.Default()
	server.Use()

	products := server.Group("/products", controller.RequireAuth)
	{
		products.GET("", controller.FindAllProducts)
		products.POST("", controller.CreateProduct)
		products.GET("/:id", controller.GetProduct)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}

	category := server.Group("/category", controller.RequireAuth)
	{

		category.GET("", controller.FindAllCategory)
		category.GET("/:id/products", controller.FindProductsByCategory)
		category.POST("", controller.CreateCategory)
		category.GET("/:id", controller.GetCategory)
		category.PUT("/:id", controller.UpdateCategory)
		category.DELETE("/:id", controller.DeleteCategory)
	}

	cart := server.Group("/cart", controller.RequireAuth)
	{

		cart.GET("", controller.FindAllCarts)
		cart.GET("/:user_login", controller.GetUserCart)
		cart.POST("", controller.CreateCart)
		cart.DELETE("/:user_login", controller.DeleteCart)
	}

	users := server.Group("/account")
	{
		users.POST("/signup", controller.SignUp)
		users.POST("/login", controller.Login)
		users.GET("/validate", controller.RequireAuth, controller.Validate)

	}

	server.Run(":8080")
}
