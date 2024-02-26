package entity

import "time"

// Promo model
type Promo struct {
	ID        int       `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Name      string    `gorm:"column:name"`
	PromoType string    `gorm:"column:promo_type"`
}

type GeneratePromoPayload struct {
	PromoName string
	StartDate time.Time
	EndDate   time.Time
	Amount    float64
}
