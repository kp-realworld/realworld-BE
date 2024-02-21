package userdto

type ReadMyProfileResponseDTO struct {
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
}

type ReadMyProfileResponseWrapperDTO struct {
	User ReadMyProfileResponseDTO `validate:"required" json:"user"`
}

type ReadUserProfileResponseDTO struct {
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
	Following    *bool   `json:"following"`
}

type ReadUserProfileResponseWrapperDTO struct {
	User ReadUserProfileResponseDTO `validate:"required" json:"user"`
}

type UpdateUserProfileRequestDTO struct {
	Username     *string `json:"username"`
	Bio          *string `json:"bio"`
	ProfileImage *string `json:"profile_image"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
}

func (dtd *UpdateUserProfileRequestDTO) IsEmpty() bool {
	return dtd.Username == nil && dtd.Bio == nil && dtd.ProfileImage == nil
}

type UpdateUserProfileResponseDTO struct {
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
}

type UpdateUserProfileResponseWrapperDTO struct {
	User UpdateUserProfileResponseDTO `validate:"required" json:"user"`
}
