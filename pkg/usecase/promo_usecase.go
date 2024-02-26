package usecase

import (
	"birthday-promo-sim/pkg/entity"
	"birthday-promo-sim/pkg/repository/db"
	"fmt"
)

type PromoUsecaseItf interface {
	GeneratePromoCode(payload entity.GeneratePromoPayload) ([]entity.UserPromoRelation, error)
}

type PromoUsecase struct {
	pr  db.PromoRepositoryItf
	aur db.AppUserRepositoryItf
}

// GeneratePromoCode implements PromoUsecaseItf.
func (r PromoUsecase) GeneratePromoCode(payload entity.GeneratePromoPayload) ([]entity.UserPromoRelation, error) {
	validUsers, err := r.aur.FetchBirthdayUsers()
	if err != nil {
		return nil, err
	}

	promoID, err := r.pr.FindPromo(payload.PromoName)
	if err != nil {
		return nil, err
	}

	userPromoEntity := []entity.UserPromoRelation{}

	for _, validUser := range validUsers {
		userPromoEntity = append(userPromoEntity, entity.UserPromoRelation{
			UserID:    validUser.ID,
			PromoID:   promoID.ID,
			Amount:    0.25,
			PromoCode: fmt.Sprintf("bday%s%d", validUser.Name, validUser.ID),
			StartDate: payload.StartDate,
			EndDate:   payload.EndDate,
		})
	}

	err = r.pr.GeneratePromoCode(userPromoEntity)
	if err != nil {
		return nil, err
	}
	return userPromoEntity, nil
}

func NewPromoUsecase(promoRepoItf db.PromoRepositoryItf, aur db.AppUserRepositoryItf) PromoUsecaseItf {
	return PromoUsecase{pr: promoRepoItf, aur: aur}
}
