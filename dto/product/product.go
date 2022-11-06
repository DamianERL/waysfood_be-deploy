package productdto

type ProductRequest struct {
	Name   string `json:"name" form:"name"`
	Price  int    `json:"price" form:"price"`
	Image  string `json:"image" form:"image"`
	UserID int    `json:"user_id" form:"user_id"`
}
