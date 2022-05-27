package handlers

import (
	"encoding/json"
	"go-advanced-rest-websockets/helpers"
	"go-advanced-rest-websockets/models"
	"go-advanced-rest-websockets/repository"
	"go-advanced-rest-websockets/server"
	"net/http"

	"github.com/gorilla/mux"
)

type UpsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"post_content"`
}

type PostUpdateResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := helpers.GetUserByJWT(s, w, r)

		if err != nil {
			return
		}

		var postRequest = UpsertPostRequest{}

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

		helpers.HttpJsonResponse(w, PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		})
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetPostById(r.Context(), params["id"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(post.Id) == 0 {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}

		helpers.HttpJsonResponse(w, post)
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := helpers.GetUserByJWT(s, w, r)

		if err != nil {
			return
		}

		var postRequest = UpsertPostRequest{}

		if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		params := mux.Vars(r)
		post := models.Post{
			Id:          params["id"],
			PostContent: postRequest.PostContent,
			UserId:      user.Id,
		}
		err = repository.UpdatePost(r.Context(), &post)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		helpers.HttpJsonResponse(w, PostUpdateResponse{
			Message: "Post updated",
		})
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := helpers.GetUserByJWT(s, w, r)

		if err != nil {
			return
		}

		params := mux.Vars(r)
		err = repository.DeletePostById(r.Context(), params["id"], user.Id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		helpers.HttpJsonResponse(w, PostUpdateResponse{
			Message: "Post deleted",
		})
	}
}
