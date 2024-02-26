package db

import (
	"birthday-promo-sim/pkg/entity"
	"time"

	"gorm.io/gorm"
)

type AppUserRepositoryItf interface {
	FetchBirthdayUsers() ([]entity.App_User, error)
}

type AppUserRepository struct {
	db *gorm.DB
}

// FetchBirthdayUsers implements AppUserRepositoryItf.
func (r *AppUserRepository) FetchBirthdayUsers() ([]entity.App_User, error) {
	var users []entity.App_User

	today := time.Now().Format("2006-01-02")
	query := r.db.Where("TO_CHAR(birthdate, 'YYYY-MM-DD') = ? and isverified = ?", today, true).Find(&users)
	if query.Error != nil {
		return nil, query.Error
	}

	return users, nil
}

func NewAppUserRepository(db *gorm.DB) AppUserRepositoryItf {
	return &AppUserRepository{db: db}
}
