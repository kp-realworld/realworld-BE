package articledto

import "time"

type CreateArticleLikeResponseDTO struct {
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

type CreateArticleLikeResponseWrapperDTO struct {
	Article CreateArticleLikeResponseDTO `validate:"required" json:"article"`
}
