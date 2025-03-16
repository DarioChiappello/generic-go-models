package routers

import (
	"github.com/DarioChiappello/generic-go-models/controllers"
	"github.com/DarioChiappello/generic-go-models/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupOrderRoutes(router *gin.Engine, db *gorm.DB) {
	orderService := services.NewOrderService(db)
	orderController := controllers.NewOrderController(orderService)

	router.GET("/orders", orderController.GetAllOrders)
	router.GET("/orders/:id", orderController.GetOrderByID)
	router.POST("/orders", orderController.CreateOrder)
	router.PUT("/orders/:id", orderController.UpdateOrder)
	router.PATCH("/orders/:id", orderController.PatchOrder)
	router.DELETE("/orders/:id", orderController.DeleteOrder)

}
