package routes

import (
	"ecommerce/controllers"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rute otentikasi
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Rute untuk produk (admin membutuhkan otorisasi)
	admin := r.Group("/admin")
	admin.Use(utils.AuthMiddleware())
	{
		admin.POST("/products", controllers.CreateProduct)
		admin.PUT("/products/:id", controllers.UpdateProduct)
		admin.DELETE("/products/:id", controllers.DeleteProduct)
	}

	// Rute untuk pengguna biasa
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)

	return r
}
