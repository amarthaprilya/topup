package database

import (
	"camera-rent/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	// dsn := "root:@tcp(127.0.0.1:3306)/camera-rent?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "freedb_camera-user:&*V$6MTbc%c7JE6@tcp(sql.freedb.tech:3306)/freedb_camera-rent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	db.AutoMigrate(&entity.User{})
	// db.AutoMigrate(&entity.Products{})
	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.Booking{})
	db.AutoMigrate(&entity.TopUp{})
	db.AutoMigrate(&entity.DoPayment{})
	db.AutoMigrate(&entity.PaymentSaldo{})
	errs := db.AutoMigrate(&entity.Products{})
	if errs != nil {
		log.Fatalf("Error during migration: %v", errs)
	}

	return db, nil

}
