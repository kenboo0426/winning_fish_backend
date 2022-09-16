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

func HandleAnswerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodPost:
		createAnswer(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func createAnswer(w http.ResponseWriter, r *http.Request) {
	var answer models.Answer
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &answer); err != nil {
		fmt.Println(err)
	}
	answer.CreatedAt = time.Now()
	answer.CreateAnswer()
	res, _ := json.Marshal(answer)
	w.Write(res)
}
