package responder

import (
	"encoding/json"
	commentdto "github.com/hotkimho/realworld-api/controller/dto/comment"
	"github.com/hotkimho/realworld-api/models"
	"net/http"
)

func CreateCommentResponse(w http.ResponseWriter, comment models.Comment) {

	wrapper := commentdto.CreateCommentResponseWrapperDTO{
		Comment: commentdto.CreateCommentResponseDTO{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			Body:      comment.Body,
			Author: commentdto.CommentAuthor{
				AuthorID:     comment.User.UserID,
				Username:     comment.User.Username,
				Bio:          comment.User.Bio,
				ProfileImage: comment.User.ProfileImage,
			},
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func ReadCommentsResponse(w http.ResponseWriter, comments []models.Comment) {

	commentList := make([]commentdto.ReadCommentsResponseDTO, 0)

	for _, comment := range comments {
		commentList = append(commentList, commentdto.ReadCommentsResponseDTO{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			Body:      comment.Body,
			Author: commentdto.CommentAuthor{
				AuthorID:     comment.User.UserID,
				Username:     comment.User.Username,
				Bio:          comment.User.Bio,
				ProfileImage: comment.User.ProfileImage,
			},
			UpdatedAt: comment.UpdatedAt,
		})
	}

	wrapper := commentdto.ReadCommentsResponseWrapperDTO{
		Comments: commentList,
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func UpdateCommentResponse(w http.ResponseWriter, comment models.Comment) {

	wrapper := commentdto.UpdateCommentResponseWrapperDTO{
		Comment: commentdto.UpdateCommentResponseDTO{
			ID:        comment.ID,
			CreatedAt: comment.CreatedAt,
			Body:      comment.Body,
			Author: commentdto.CommentAuthor{
				AuthorID:     comment.User.UserID,
				Username:     comment.User.Username,
				Bio:          comment.User.Bio,
				ProfileImage: comment.User.ProfileImage,
			},
		},
	}

	jsonData, err := json.Marshal(wrapper)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
