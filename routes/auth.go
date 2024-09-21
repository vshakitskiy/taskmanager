package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gotaskmaster/models"
	"net/http"
)

func HandleAuth() chi.Router {
	r := chi.NewRouter()

	r.Post("/cred/signup", signUpWithCredentials)
	r.Post("/cred/signin", signInWithCredentials)

	r.Post("/github", authWithGithub)

	return r
}

func signUpWithCredentials(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Printf("%v", newUser)
	if err != nil {
		SendJSON(w, http.StatusBadRequest, err.Error())
	}
}

func signInWithCredentials(w http.ResponseWriter, r *http.Request) {}

func authWithGithub(w http.ResponseWriter, r *http.Request) {}
