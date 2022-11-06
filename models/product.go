package models

import (
	"time"
)

type Product struct {
	ID        int         `json:"id" gorm:"primary_key:auto_increment"`
	Name      string      `json:"name" gorm:"type:varchar (255)"`
	Price     int         `json:"price" gorm:"type:int"`
	Image     string      `json:"image" gorm:"type:varchar(255)"`
	UserID    int         `json:"-" `
	User      UserProfile `json:"user"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}

type ProductResponse struct {
	ID     int         `json:"id"`
	Name   string      `json:"Name"`
	Price  int         `json:"price"`
	Image  string      `json:"image"`
	UserID int         `json:"-"`
	User   UserProfile `json:"user"`
}

func (ProductResponse) TableName() string {
	return "products"
}
