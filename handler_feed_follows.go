package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepak04112002/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	}

    err = apiCfg.DB.CreateFeedFollow(r.Context(), createFeedFollowParams)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error create Feed follow: %s", err))
		return
	}
	feedFollow, err := apiCfg.DB.GetFeedFollowByID(r.Context(), createFeedFollowParams.ID)
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error retrieving created feed: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollows(r.Context())
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("couldn't get feeds follow: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
     feedFollowIDStr:=chi.URLParam(r,"feedFollowID")
	 feedFollowID,err:=uuid.Parse(feedFollowIDStr)
	 if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt parse feed follow ID: %s", err))
		return
	}
	err=apiCfg.DB.DeleteFeedFollows(r.Context(),database.DeleteFeedFollowsParams{
		ID: feedFollowID.String(),
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldnt delete feed follow: %s", err))
		return
	}
	responseWithJSON(w,200,struct{}{})
}
