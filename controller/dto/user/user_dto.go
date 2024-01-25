package user

type ReadUserProfileResponseDTO struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
}

type ReadUserProfileResponseWrapperDTO struct {
	User ReadUserProfileResponseDTO `json:"user"`
}

type UpdateUserProfileRequestDTO struct {
	Username     *string `json:"username"`
	Bio          *string `json:"bio"`
	ProfileImage *string `json:"profile_image"`
}

type UpdateUserProfileResponseDTO struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
}

type UpdateUserProfileResponseWrapperDTO struct {
	User UpdateUserProfileResponseDTO `json:"user"`
}
