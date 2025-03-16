package routers

import (
	"github.com/DarioChiappello/generic-go-models/controllers"
	"github.com/DarioChiappello/generic-go-models/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	productService := services.NewProductService(db)
	productController := controllers.NewProductController(productService)

	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductByID)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:id", productController.UpdateProduct)
	router.PATCH("/products/:id", productController.PatchProduct)
	router.DELETE("/products/:id", productController.DeleteProduct)

}
