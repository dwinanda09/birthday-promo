package notification

import (
	"birthday-promo-sim/pkg/entity"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func SendNotification(redisClient *redis.Client, users []entity.UserPromoRelation) error {
	ctx := context.TODO()
	for _, user := range users {
		notification := entity.NotificationPayload{
			NotificationType: "email",
			Subject:          "Happy Birthday! And Congratulation You Got Our Birthday Promo!",
			Body: fmt.Sprintf(`🎉 Happy Birthday %s! 🎂
			Celebrate your special day with us! 🎁 As a token of our appreciation, we're delighted to offer you an exclusive birthday promo. 🎈 Use code %s at checkout to enjoy a %f off on your next purchase.
			May your day be filled with joy, laughter, and fantastic surprises! 🥳 Don't miss out on this special birthday treat – it's our way of saying thank you for being a valued part of our Buyer family.
			Wishing you a year ahead filled with happiness and wonderful moments! 🎊
			Best regards,
			SayaKaya.id Team`, user.Email, user.PromoCode, user.Amount),
			Target: user.Email,
		}
		// Convert the NotificationPayload to JSON
		notificationJSON, err := json.Marshal(notification)
		if err != nil {
			return err
		}
		err = redisClient.Publish(ctx, "notifications", notificationJSON).Err()
		fmt.Printf("notification is sent to %s's %s\n", user.Email, notification.NotificationType)
		if err != nil {
			return err
		}
	}
	return nil
}
