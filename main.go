package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	configs "github.com/DarioChiappello/generic-go-models/config"
	"github.com/DarioChiappello/generic-go-models/models"
	"github.com/DarioChiappello/generic-go-models/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	// Inicializar base de datos
	db := configs.ConnectDB()

	// Realizar migraciones
	if err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	); err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
		return
	}

	// Configurar router
	router := gin.Default()

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	// Validar que los orígenes sean válidos
	for _, origin := range allowedOrigins {
		if origin == "*" {
			allowedOrigins = []string{"*"}
			break
		}
		if !strings.HasPrefix(origin, "http://") && !strings.HasPrefix(origin, "https://") {
			log.Fatal("Invalid origin format in ALLOWED_ORIGINS: ", origin)
		}
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configurar rutas
	routers.SetupUserRoutes(router, db)
	routers.SetupProductRoutes(router, db)
	routers.SetupOrderRoutes(router, db)

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
