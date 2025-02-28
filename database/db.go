package database

import (
	"camera-rent/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	// if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
	// 	if err := godotenv.Load(); err != nil {
	// 		log.Fatal("error loading .env file:", err)
	// 	}
	// }

	// databaseURL := os.Getenv("DATABASE_URL")

	// // Check if DATABASE_URL is set
	// if databaseURL == "" {
	// 	log.Fatal("DATABASE_URL environment variable is not set")
	// }

	// // Open the database connection using the URL
	// db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB Connection Error:", err)
	// }

	dsn := "freedb_customeruser:pHQQcE@ttPyB#5Z@tcp(sql.freedb.tech:3306)/freedb_customerdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}
	// Auto Migration

	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(&entity.User{})
	// db.AutoMigrate(&entity.Products{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Booking{})
	db.AutoMigrate(&entity.TopUp{})
	db.AutoMigrate(&entity.DoPayment{})
	db.AutoMigrate(&entity.PaymentSaldo{})
	errs := db.AutoMigrate(&entity.Products{})
	if errs != nil {
		// Tangani kesalahan di sini, misalnya dengan mencetak pesan kesalahan atau mengembalikan kesalahan
		log.Fatalf("Error during migration: %v", errs)
	}

	return db, nil

}
