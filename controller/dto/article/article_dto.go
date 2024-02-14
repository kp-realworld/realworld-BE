package articledto

import "time"

type CreateArticleRequestDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tag_list,omitempty"`
}

type ArticleAuthor struct {
	AuthorID     int64   `json:"author_id"`
	Username     string  `json:"username"`
	Bio          *string `json:"bio,omitempty"`
	ProfileImage string  `json:"profile_image"`
	Following    *bool   `json:"following,omitempty"`
}

type CreateArticleResponseDTO struct {
	ID            int64         `json:"article_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
}

type CreateArticleResponseWrapperDTO struct {
	Article CreateArticleResponseDTO `json:"article"`
}

type ReadArticleResponseDTO struct {
	ID            int64         `json:"article_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
}

type ReadArticleResponseWrapperDTO struct {
	Article ReadArticleResponseDTO `json:"article"`
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
	ID            int64         `json:"article_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
}

type UpdateArticleResponseWrapperDTO struct {
	Article UpdateArticleResponseDTO `json:"article"`
}

type ReadMyArticlesResponseWrapperDTO struct {
	Articles []ReadArticleResponseDTO `json:"articles"`
}
