package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahochbaum-rcg/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCon *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respongWithError(w, 400, fmt.Sprintf("Error parsing JOSN: %v", err))
		return
	}

	feed, err := apiCon.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
		Url:       param.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respongWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	responWithJSON(w, 201, databaseFeedToFeed(feed))
}
