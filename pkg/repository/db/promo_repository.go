package db

import (
	"birthday-promo-sim/pkg/entity"

	"gorm.io/gorm"
)

type PromoRepositoryItf interface {
	GeneratePromoCode(payload []entity.UserPromoRelation) error
	FindPromo(promoType string) (entity.Promo, error)
}

type PromoRepository struct {
	db *gorm.DB
}

// FindPromo implements PromoRepositoryItf.
func (r *PromoRepository) FindPromo(promoType string) (entity.Promo, error) {
	var promo entity.Promo

	err := r.db.Table("promo").Where("promo_type = ?", promoType).First(&promo)
	if err != nil {
		return entity.Promo{}, err.Error
	}

	return promo, nil
}

// GeneratePromoCode implements PromoRepositoryItf.
func (r *PromoRepository) GeneratePromoCode(payload []entity.UserPromoRelation) error {
	return r.db.Table("user_promo_relation").Save(payload).Error
}

func NewPromoRepository(db *gorm.DB) PromoRepositoryItf {
	return &PromoRepository{db: db}
}
