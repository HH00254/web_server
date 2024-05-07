package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahochbaum-rcg/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCon *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respongWithError(w, 400, fmt.Sprintf("Error parsing JOSN: %v", err))
		return
	}

	user, err := apiCon.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      param.Name,
	})
	if err != nil {
		respongWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	responWithJSON(w, 200, databaseUserToUser(user))
}
