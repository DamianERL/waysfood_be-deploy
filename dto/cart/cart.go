package cartdto

type CreateCart struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	QTY      int `json:"qty"`
	Shipping int `json:"shipping"`
	OrderID  int `json:"order_id"`
	Status   int `json:"Status"`
	Total    int `json:"total"`
}

type UpdateCart struct {
	Status   string `json:"Status"`
	SubTotal int    `json:"Subtotal"`
	Shipping int    `json:"shipping"`
	Total    int    `json:"total"`
	QTY      int    `json:"qty"`
}
