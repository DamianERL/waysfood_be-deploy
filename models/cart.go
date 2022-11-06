package models

import "time"

type Cart struct {
	ID        int         `json:"id" gorm:"primary_key:auto_increment" `
	UserID    int         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	User      UserProfile `json:"user"`
	Order     []Order     `json:"order"`
	QTY       int         `json:"qty"`
	Total     int         `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt time.Time   `json:"-"`
}
