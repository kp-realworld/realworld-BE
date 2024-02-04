package authdto

// 회원가입
type SignUpRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpResponseDTO struct {
	UserID   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type SignUpResponseWrapperDTO struct {
	User SignUpResponseDTO `json:"user"`
}

// 로그인
type SignInRequestDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignInResponseDTO struct {
	UserID   int64  `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token"`
}

type SignInResponseWrapperDTO struct {
	User SignInResponseDTO `json:"user"`
}

// 토큰 갱신
type RefreshTokenResponseDTO struct {
	Token string `json:"token"`
}

// username 확인
type VerifyUsernameRequestDTO struct {
	Username string `json:"username" validate:"required"`
}

type VerifyUsernameResponseDTO struct {
	Username string `json:"username,omitempty"`
}

// email 확인
type VerifyEmailRequestDTO struct {
	Email string `json:"email" validate:"required"`
}

type VerifyEmailResponseDTO struct {
	Email string `json:"email,omitempty"`
}
