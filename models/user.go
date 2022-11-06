package models

import "time"

type User struct {
	ID           int               `json:"id" gorm:"primary_key:auto_increment"`
	Name         string            `json:"name" gorm:"type: varchar(255)"`
	Email        string            `json:"email" gorm:"type: varchar(255)"`
	Password     string            `json:"password" gorm:"type: varchar(255)"`
	Gender       string            `json:"gender" gorm:"type: varchar(255)"`
	Phone        string            `json:"phone" gorm:"type: varchar(255)"`
	Role         string            `json:"role" gorm:"type: varchar(255)"`
	Location     string            `json:"location" gorm:"type:varchar(255)"`
	Image        string            `json:"image" gorm:"type:varchar(255)"`
	Products     []ProductResponse `json:"product" `
	Carts        []Cart            `json:"carts"`
	Transactions []Transaction     `json:"transactions" gorm:"foreignKey:BuyerID" `
	Income       []Transaction     `json:"income" gorm:"foreignKey:BuyerID" `
	CreatedAt    time.Time         `json:"-"`
	UpdatedAt    time.Time         `json:"-"`
}

type UserProfile struct {
	ID       int    `json:"id" `
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender" `
	Phone    string `json:"phone" `
	Role     string `json:"role" `
	Location string `json:"location" `
	Image    string `json:"image" `
}

func (UserProfile) TableName() string {
	return "users"
}
