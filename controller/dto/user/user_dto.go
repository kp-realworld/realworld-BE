package userdto

type ReadMyProfileResponseDTO struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
}

type ReadMyProfileResponseWrapperDTO struct {
	User ReadMyProfileResponseDTO `json:"user"`
}

type ReadUserProfileResponseDTO struct {
	UserID       int64   `json:"user_id"`
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
	Following    *bool   `json:"following"`
	Email        string  `json:"email"`
}

type ReadUserProfileResponseWrapperDTO struct {
	User ReadUserProfileResponseDTO `json:"user"`
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
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
}

type UpdateUserProfileResponseWrapperDTO struct {
	User UpdateUserProfileResponseDTO `json:"user"`
}
