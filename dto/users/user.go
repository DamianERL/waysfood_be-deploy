package usersdto

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Location string `json:"location" form:"location"`
	Phone    string `json:"phone" form:"phone"`
	Image    string `json:"image" form:"image"`
}
