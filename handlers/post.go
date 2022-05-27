package handlers

import (
	"encoding/json"
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/repository"
	"go-advanced-rest-websockets/server"
	"net/http"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := helpers.GetUserByJWT(s, w, r)

		if err != nil {
			return
		}

		var postRequest = InsertPostRequest{}

		if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := helpers.GenerateId()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post := models.Post{
			Id:          id,
			PostContent: postRequest.PostContent,
			UserId:      user.Id,
		}
		err = repository.InsertPost(r.Context(), &post)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		})
	}
}
