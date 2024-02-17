package articledto

import "time"

type CreateArticleLikeResponseDTO struct {
	ID            int64         `json:"article_id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Body          string        `json:"body"`
	FavoriteCount int64         `json:"favorite_count"`
	TagList       []string      `json:"tag_list,omitempty"`
	IsFavorited   bool          `json:"is_favorited"`
	Author        ArticleAuthor `json:"author"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     *time.Time    `json:"updated_at,omitempty"`
}

type CreateArticleLikeResponseWrapperDTO struct {
	Article CreateArticleLikeResponseDTO `json:"article"`
}
