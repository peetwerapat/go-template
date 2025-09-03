package domain

import "time"

type User struct {
	ID        uint       `gorm:"column:id;primaryKey" json:"id"`
	Email     string     `gorm:"column:email" json:"email"`
	Password  string     `gorm:"password"`
	FirstName string     `gorm:"first_name" json:"firstName"`
	LastName  string     `gorm:"last_name" json:"lastName"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}
