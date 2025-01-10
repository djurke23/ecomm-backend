package main

import (
	"log"

	"github.com/djurke23/ecomm-backend/config"
	"github.com/djurke23/ecomm-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Povezivanje na bazu
	config.ConnectDB()

	// 2. Pokretanje migracija (samo za razvoj)
	config.MigrateDB(config.DB)

	// 3. Inicijalizacija Gin router-a
	router := gin.Default()

	// 4. Postavljanje ruta
	routes.SetupRoutes(router)

	// 5. Pokretanje servera
	log.Println("Pokrećem server na portu 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Greška pri pokretanju servera:", err)
	}
}
