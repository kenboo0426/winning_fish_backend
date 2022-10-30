package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"winning_fish_backend/app/models"
)

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	switch r.Method {
	case http.MethodPost:
		findOrCreateUser(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}

}

func HandleUserUpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	switch r.Method {
	case http.MethodGet:
		showUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodOptions:
	// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func showUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func findOrCreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println(err)
	}
	existerdUser, notRecordErr := models.GetUserByIDOrUUID(user.ID, user.UUID)
	if notRecordErr == nil {
		user = existerdUser
	} else {
		user.CreatedAt = time.Now()
		user.CreateUser()
		w.WriteHeader(201)
	}
	res, _ := json.Marshal(user)
	w.Write(res)
}
