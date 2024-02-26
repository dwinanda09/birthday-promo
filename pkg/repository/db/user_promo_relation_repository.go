package db

import "birthday-promo-sim/pkg/entity"

type UserPromoRelationRepository interface {
	StoreUserPromoRelation(users []entity.App_User, promo entity.Promo) error
}
