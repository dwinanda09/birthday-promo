package scheduler

import (
	"birthday-promo-sim/pkg/entity"
	"birthday-promo-sim/pkg/notification"
	"birthday-promo-sim/pkg/usecase"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CronScheduler struct {
	promoUsecase usecase.PromoUsecaseItf
}

func NewCronScheduler(pr usecase.PromoUsecaseItf) *CronScheduler {
	return &CronScheduler{promoUsecase: pr}
}

func (cs *CronScheduler) Start() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", cs.triggerNotification) // Run daily at midnight
	if err != nil {
		log.Fatal(err)
	}

	c.Start()

	select {}
}

func (cs *CronScheduler) triggerNotification() {
	log.Println("Notification cron job triggered.")
	// Perform the notification logic here
	// Initialize necessary dependencies (e.g., database, Redis, etc.)
	redisClient := initRedis()

	payload := entity.GeneratePromoPayload{
		PromoName: "birthday",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Amount:    0.25,
	}
	userWithPromo, err := cs.promoUsecase.GeneratePromoCode(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Use the notification package to send notifications
	err = notification.SendNotification(redisClient, userWithPromo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notification cron job completed.")
}

// initRedis initializes the Redis client
func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Replace with your Redis server address
		DB:   0,
	})
	return rdb
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

	return db
}
