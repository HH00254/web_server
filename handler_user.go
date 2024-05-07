package main

import (
	"encoding/json"
	"net/http"
)

func (apiCon *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}

	decoder := json.NewDecoder(r.body)

	param := parameters{}
	err := decoder.Decode(&param)
	if err != nil {
		respongWithError(w, 400, "Failed to decode JSON body")
		return
	}

	responWithJSON(w, 200, struct{}{})
}
