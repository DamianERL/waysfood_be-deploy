package ordersdto

type OrderRequest struct {
	ProductID int `json:"product_id"`
	SubAmount int `json:"sub_amount"`
}

type OrderUpdate struct {
	QTY       int `json:"qty"  `
	SubAmount int `json:"sub_amount"`
}
