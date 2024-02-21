package commentdto

import "time"

// article author 구조체가 있어도 분리
type CommentAuthor struct {
	AuthorID     int64   `validate:"required" json:"author_id"`
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
}

type CreateCommentRequestDTO struct {
	Body string `validate:"required" json:"body"`
}

type CreateCommentResponseDTO struct {
	ID        int64         `validate:"required" json:"comment_id"`
	Body      string        `validate:"required" json:"body"`
	Author    CommentAuthor `validate:"required" json:"author"`
	CreatedAt time.Time     `json:"created_at"`
}

type CreateCommentResponseWrapperDTO struct {
	Comment CreateCommentResponseDTO `validate:"required" json:"comment"`
}

type ReadCommentsResponseDTO struct {
	ID        int64         `validate:"required" json:"comment_id"`
	Body      string        `validate:"required" json:"body"`
	Author    CommentAuthor `validate:"required" json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

type ReadCommentsResponseWrapperDTO struct {
	Comments []ReadCommentsResponseDTO `json:"comments"`
}

type UpdateCommentRequestDTO struct {
	Body string `validate:"required" json:"body"`
}

type UpdateCommentResponseDTO struct {
	ID        int64         `validate:"required" json:"comment_id"`
	Body      string        `validate:"required" json:"body"`
	Author    CommentAuthor `validate:"required" json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

type UpdateCommentResponseWrapperDTO struct {
	Comment UpdateCommentResponseDTO `validate:"required" json:"comment"`
}
