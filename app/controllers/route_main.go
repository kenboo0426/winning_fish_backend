package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"winning_fish_backend/app/models"
)

func indexQuiz(w http.ResponseWriter, r *http.Request) {
	quizzes, err := models.GetQuizzes()
	if err != nil {
		log.Fatalln(err)
	}
	res, _ := json.Marshal(quizzes)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	w.Write(res)
}

func HandleQuizRequst(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	switch r.Method {
	case http.MethodPost:
		createQuiz(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func HandleQuizUpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	switch r.Method {
	case http.MethodGet:
		showQuiz(w, r)
	case http.MethodPut:
		updateQuiz(w, r)
	case http.MethodDelete:
		deleteQuiz(w, r)
	case http.MethodOptions:
	// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func showQuiz(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/quiz")
	_, id := filepath.Split(sub)

	quiz_id, _ := strconv.Atoi(id)
	quiz, err := models.GetQuiz(quiz_id)
	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(quiz)
	w.Write(res)
}

func createQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &quiz); err != nil {
		fmt.Println(err)
	}
	quiz.CreatedAt = time.Now()

	id, err := quiz.CreateQuiz()
	if err != nil {
		log.Fatalln(err)
	}
	quiz.ID = int(id)
	correct_option_id, _ := quiz.CreateOptions()
	quiz.CorrectID = int(correct_option_id)
	quiz.UpdateQuiz()
	res, _ := json.Marshal(quiz)
	w.Write(res)
}

func updateQuiz(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/quiz")
	_, id := filepath.Split(sub)

	quiz_id, _ := strconv.Atoi(id)
	quiz, err := models.GetQuiz(quiz_id)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &quiz); err != nil {
		fmt.Println(err)
	}
	err = quiz.UpdateQuiz()

	res, _ := json.Marshal(quiz)
	w.Write(res)
}

func deleteQuiz(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/quiz")
	_, id := filepath.Split(sub)

	quiz_id, _ := strconv.Atoi(id)
	quiz, err := models.DeleteQuiz(quiz_id)

	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(quiz)
	w.Write(res)
}
