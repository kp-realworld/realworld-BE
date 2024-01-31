package commentdto

import "time"

// article author 구조체가 있어도 분리
type CommentAuthor struct {
	AuthorID     int64   `json:"author_id"`
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
}

type CreateCommentRequestDTO struct {
	Body string `json:"body"`
}

type CreateCommentResponseDTO struct {
	ID        int64         `json:"comment_id"`
	Body      string        `json:"body"`
	Author    CommentAuthor `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
}

type CreateCommentResponseWrapperDTO struct {
	Comment CreateCommentResponseDTO `json:"comment"`
}

type ReadCommentsResponseDTO struct {
	ID        int64         `json:"comment_id"`
	Body      string        `json:"body"`
	Author    CommentAuthor `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

type ReadCommentsResponseWrapperDTO struct {
	Comments []ReadCommentsResponseDTO `json:"comments"`
}

type UpdateCommentRequestDTO struct {
	Body string `json:"body"`
}

type UpdateCommentResponseDTO struct {
	ID        int64         `json:"comment_id"`
	Body      string        `json:"body"`
	Author    CommentAuthor `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
}

type UpdateCommentResponseWrapperDTO struct {
	Comment UpdateCommentResponseDTO `json:"comment"`
}
