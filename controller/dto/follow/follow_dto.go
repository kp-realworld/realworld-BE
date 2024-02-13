package followdto

type CreateFollowResponseDTO struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
	Following    bool    `json:"following"`
}

type CreateFollowResponseWrapperDTO struct {
	Profile CreateFollowResponseDTO `json:"profile"`
}

type DeleteFollowResponseDTO struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
	Following    bool    `json:"following"`
}

type DeleteFollowResponseWrapperDTO struct {
	Profile DeleteFollowResponseDTO `json:"profile"`
}
