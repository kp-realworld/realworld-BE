package followdto

type CreateFollowResponseDTO struct {
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
	Following    bool    `validate:"required" json:"following"`
}

type CreateFollowResponseWrapperDTO struct {
	Profile CreateFollowResponseDTO `validate:"required" json:"user"`
}

type DeleteFollowResponseDTO struct {
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
	Following    bool    `validate:"required" json:"following"`
}

type DeleteFollowResponseWrapperDTO struct {
	Profile DeleteFollowResponseDTO `validate:"required" json:"user"`
}
