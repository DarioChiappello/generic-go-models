package routers

import (
	"github.com/DarioChiappello/generic-go-models/controllers"
	"github.com/DarioChiappello/generic-go-models/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.PUT("/users/:id", userController.UpdateUser)
	router.PATCH("/users/:id", userController.PatchUser)
	router.DELETE("/users/:id", userController.DeleteUser)
	router.GET("/users", userController.GetAllUsers)

}
