package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Image    string `json:"image"`
	Token    string `json:"token"`
	Location string `json:"location"`
	// Password string `json:"password" `
}

type AuthResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Image string `json:"image"`
	Role  string `json:"role"`
}
