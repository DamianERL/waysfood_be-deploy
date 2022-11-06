package authdto

type RegisterRequest struct {
	Email    string `validate:"required"  form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	Name     string `form:"name" json:"name"`
	Gender   string `form:"gender" json:"gender"`
	Phone    string `form:"phone" json:"phone"`
	Role     string `form:"role" json:"role"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
