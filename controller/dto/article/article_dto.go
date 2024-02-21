package articledto

import "time"

type CreateArticleRequestDTO struct {
	// title은 필수값
	Title       string   `validate:"required" json:"title"`
	Description string   `validate:"required" json:"description"`
	Body        string   `validate:"required" json:"body"`
	TagList     []string `json:"tag_list,omitempty"`
}

type ArticleAuthor struct {
	AuthorID     int64   `validate:"required" json:"author_id"`
	Username     string  `validate:"required" json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `validate:"required" json:"profile_image"`
	Following    *bool   `json:"following,omitempty"`
}

type CreateArticleResponseDTO struct {
	ID            int64         `validate:"required" json:"article_id"`
	Title         string        `validate:"required" json:"title"`
	Description   string        `validate:"required" json:"description"`
	Body          string        `validate:"required" json:"body"`
	FavoriteCount int64         `validate:"required" json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `validate:"required" json:"is_favorited"`
	Author        ArticleAuthor `validate:"required" json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
}

type CreateArticleResponseWrapperDTO struct {
	Article CreateArticleResponseDTO `validate:"required" json:"article"`
}

type ReadArticleResponseDTO struct {
	ID            int64         `validate:"required" json:"article_id"`
	Title         string        `validate:"required" json:"title"`
	Description   string        `validate:"required" json:"description"`
	Body          string        `validate:"required" json:"body"`
	FavoriteCount int64         `validate:"required" json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `validate:"required" json:"is_favorited"`
	Author        ArticleAuthor `validate:"required" json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
}

type ReadArticleResponseWrapperDTO struct {
	Article ReadArticleResponseDTO `validate:"required" json:"article"`
}

type ReadArticleByOffsetResponseWrapperDTO struct {
	Articles []ReadArticleResponseDTO `json:"articles"`
}

type ReadArticlesByUserResponseWrapperDTO struct {
	Articles []ReadArticleResponseDTO `json:"articles"`
}

type UpdateArticleRequestDTO struct {
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Body        *string  `json:"body,omitempty"`
	TagList     []string `json:"tag_list,omitempty"`
}

type UpdateArticleResponseDTO struct {
	ID            int64         `validate:"required" json:"article_id"`
	Title         string        `validate:"required" json:"title"`
	Description   string        `validate:"required" json:"description"`
	Body          string        `validate:"required" json:"body"`
	FavoriteCount int64         `validate:"required" json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `validate:"required" json:"is_favorited"`
	Author        ArticleAuthor `validate:"required" json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
}

type UpdateArticleResponseWrapperDTO struct {
	Article UpdateArticleResponseDTO `validate:"required" json:"article"`
}

type ReadMyArticlesResponseWrapperDTO struct {
	Articles []ReadArticleResponseDTO `json:"articles"`
}
