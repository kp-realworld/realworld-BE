package authdto

// 회원가입
type SignUpRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpResponseDTO struct {
	UserID   int64  `validate:"required" json:"user_id,omitempty"`
	Username string `validate:"required" json:"username,omitempty"`
	Email    string `validate:"required" json:"email,omitempty"`
}

type SignUpResponseWrapperDTO struct {
	User SignUpResponseDTO `validate:"required" json:"user"`
}

// 로그인
type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponseDTO struct {
	UserID       int64  `validate:"required" json:"user_id,omitempty"`
	Username     string `validate:"required" json:"username,omitempty"`
	AccessToken  string `validate:"required" json:"access_token"`
	RefreshToken string `validate:"required" json:"refresh_token"`
}

type SignInResponseWrapperDTO struct {
	User SignInResponseDTO `validate:"required" json:"user"`
}

// 토큰 갱신
type RefreshTokenResponseDTO struct {
	Token string `validate:"required" json:"token"`
}

// username 확인
type VerifyUsernameRequestDTO struct {
	Username string `json:"username" validate:"required"`
}

type VerifyUsernameResponseDTO struct {
	Username string `validate:"required" json:"username,omitempty"`
}

// email 확인
type VerifyEmailRequestDTO struct {
	Email string `json:"email" validate:"required"`
}

type VerifyEmailResponseDTO struct {
	Email string `validate:"required" json:"email,omitempty"`
}
