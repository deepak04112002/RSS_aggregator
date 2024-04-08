package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepak04112002/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreatedUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	createParams := database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), createParams)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:10,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't get posts : %v", err))
		return
	}
	responseWithJSON(w,200, databasePostsToPosts(posts))
}
