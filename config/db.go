package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/korisnik/ecomm-backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Primer: kredencijali su "hardkodirani" za demo potrebe.
	// Najbolje je koristiti environment promenljive ili .env fajl.
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "ecommerce"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Greška pri povezivanju na bazu:", err)
	}

	DB = db
	log.Println("Uspješna konekcija na bazu!")
}

// MigrateDB - automatska migracija modela
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Product{},
		// Dodaj i ostale modele koje kreiraš (Order, User, itd.)
	)
	if err != nil {
		log.Fatal("Greška pri migraciji:", err)
	}
	log.Println("Migracija uspešno obavljena!")
}
