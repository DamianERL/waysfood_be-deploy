package transactiondto

type RequestTransaction struct {
	BuyerID  int    `json:"buyer"`
	SellerID int    `json:"seller_id"`
	CartID   int    `json:"cart_id"`
	Total    int    `json:"total"`
	Status   string `jspn:"status"`
}

type ResponseTransaction struct {
	ID     int    `json:"id"`
	Cart   string `json:"cart"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller_id"`
	Total  int    `json:"total"`
	Status string `json:"status"`
}
