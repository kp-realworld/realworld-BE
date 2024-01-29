package articledto

import "time"

type CreateArticleRequestDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tag_list",omitempty`
}

type ArticleAuthor struct {
	Username     string  `json:"username"`
	Bio          *string `json:"bio"`
	ProfileImage string  `json:"profile_image"`
}

type CreateArticleResponseDTO struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
}

type CreateArticleResponseWrapperDTO struct {
	Article CreateArticleResponseDTO `json:"article"`
}

type ReadArticleResponseDTO struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
}

type ReadArticleResponseWrapperDTO struct {
	Article ReadArticleResponseDTO `json:"article"`
}

type ReadArticleListResponseDTO struct {
	Articles []ReadArticleResponseDTO `json:"articles"`
}

type UpdateArticleRequestDTO struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Body        *string  `json:"body"`
	TagList     []string `json:"tag_list"`
}

type UpdateArticleResponseDTO struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int           `json:"favorite_count"`
	TagList       []string      `json:"tag_list"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
}

type UpdateArticleResponseWrapperDTO struct {
	Article UpdateArticleResponseDTO `json:"article"`
}
