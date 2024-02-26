package entity

import "time"

// App_User model
type App_User struct {
	ID         int       `gorm:"column:id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	Name       string    `gorm:"column:name"`
	Email      string    `gorm:"column:email"`
	Phone      string    `gorm:"column:phone"`
	Birthdate  time.Time `gorm:"column:birthdate"`
	IsVerified bool      `gorm:"column:isverified"`
}
