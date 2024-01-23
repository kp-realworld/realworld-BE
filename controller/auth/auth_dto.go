package auth

type SignUpRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name"`
}

type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponseDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
