package handlers

import (
	"birthday-promo-sim/pkg/entity"
	"birthday-promo-sim/pkg/notification"
	"birthday-promo-sim/pkg/usecase"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type APIHandler struct {
	PromoUsecase usecase.PromoUsecaseItf
}

func NewAPIHandler(promoUsecase usecase.PromoUsecaseItf) *APIHandler {
	return &APIHandler{PromoUsecase: promoUsecase}
}

func (h *APIHandler) NotificationHandler(w http.ResponseWriter, r *http.Request) {

	redisClient := initRedis()

	payload := entity.GeneratePromoPayload{
		PromoName: "birthday",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 1),
		Amount:    0.25,
	}
	userWithPromo, err := h.PromoUsecase.GeneratePromoCode(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
		return
	}

	err = notification.SendNotification(redisClient, userWithPromo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Notification Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Notification handled successfully"})
}

func SetupRouter(apiHandler *APIHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/notify", apiHandler.NotificationHandler).Methods("POST")
	return r
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return rdb
}
