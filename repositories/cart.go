package repositories

import (
	"waysfood/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCart() ([]models.Cart, error)
	GEtCart(ID int) (models.Cart, error)
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart) (models.Cart, error)
	//takutnya cart yang udah suksess malah ke fetching juga
	//kita nyari cart yang statusnya masih pending /belum dibayar(biar data yang
	//udah di fetching bakal tinggal di chart dan yang sudah tidak tinggal di cart)
	FindbyIDCart(CartId int, Status string) (models.Cart, error)
	//ngk dipakai
	// GetOneCart(ID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCart() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Find(&cart).Error

	return cart, err
}

func (r *repository) GEtCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error
	return cart, err
}

func (r *repository) FindbyIDCart(CartId int, Status string) (models.Cart, error) {
	var Cart models.Cart
	err := r.db.Preload("User").Preload("Order").Preload("Order.Product").Preload("Order.Product.User").Where("user_id = ? AND status = ?", CartId, Status).First(&Cart).Error

	return Cart, err
}
