package entity

import "time"

// UserPromoRelation model to store the relation between users and promos
type UserPromoRelation struct {
	ID        int       `gorm:"column:id"`
	Email     string    `gorm:"column:email"`
	UserID    int       `gorm:"column:user_id;uniqueIndex:idx_user_promo"`
	PromoID   int       `gorm:"column:promo_id;uniqueIndex:idx_user_promo"`
	Amount    float64   `gorm:"column:amount"`
	PromoCode string    `gorm:"column:promo_code"`
	StartDate time.Time `gorm:"column:start_date"`
	EndDate   time.Time `gorm:"clumn:end_date"`
}
