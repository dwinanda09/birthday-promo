package main

import (
	"birthday-promo-sim/pkg/entity"
	repo "birthday-promo-sim/pkg/repository/db"
	"birthday-promo-sim/pkg/scheduler"
	"birthday-promo-sim/pkg/usecase"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NotificationPayload struct {
	NotificationType string
	Subject          string
	Body             string
	Target           string
}

var NOW_TIME = time.Now()

func main() {
	// Initialize database connection
	db := initDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// for seeding the DB
	// err = SeedDB(db)
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Create instances of repositories
	appUserRepo := repo.NewAppUserRepository(db)
	promoRepo := repo.NewPromoRepository(db)

	// Create an instance of the PromoUsecase
	promoUsecase := usecase.NewPromoUsecase(promoRepo, appUserRepo)

	cronScheduler := scheduler.NewCronScheduler(promoUsecase)
	go cronScheduler.Start()

	// Keep the main Goroutine running
	select {}

}

// SeedDBUser populates the database with initial data
func SeedDBUser(db *gorm.DB) error {
	users := []entity.App_User{
		{Name: "Alice", Email: "alice@ona.com", Birthdate: NOW_TIME, IsVerified: true, CreatedAt: NOW_TIME, UpdatedAt: NOW_TIME},
		{Name: "Bob", Email: "bob@pol.com", Birthdate: NOW_TIME.AddDate(0, 0, -25), IsVerified: false, CreatedAt: NOW_TIME, UpdatedAt: NOW_TIME},
		{Name: "Clark", Email: "clark@kad.com", Birthdate: NOW_TIME.AddDate(0, 0, -14), IsVerified: true, CreatedAt: NOW_TIME, UpdatedAt: NOW_TIME},
	}

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func initDB() *gorm.DB {
	// Connection parameters from environment variables
	host := "localhost"
	user := "postgres"
	// password := "postgres"
	dbname := "postgres"
	port := "5432"

	fmt.Println("connecting..,.")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, user, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is connected")
	return db
}
