package models

import "time"

type Transaction struct {
	ID        int         `json:"id" gorm:"primary_key:auto_increment" `
	BuyerID   int         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" `
	Buyer     UserProfile `json:"buyer"`
	SellerID  int         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Seller    UserProfile `json:"seller"`
	CartID    int         `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  `
	Cart      Cart        `json:"cart"  `
	Total     int         `json:"total"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"create_at"`
	UpdatedAt time.Time   `json:"-"`
}
