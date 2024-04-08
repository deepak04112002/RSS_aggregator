package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/deepak04112002/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	}

    err = apiCfg.DB.CreateFeed(r.Context(), createFeedParams)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating Feed: %s", err))
		return
	}
	createdFeed, err := apiCfg.DB.GetFeedByID(r.Context(), createFeedParams.ID)
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error retrieving created feed: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(createdFeed))
}
func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
    feeds,err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error get feeds: %s", err))
		return
	}
	responseWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}