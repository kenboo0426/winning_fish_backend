package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	w.Write(res)
}

func HandleQuizRequst(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showQuiz(w, r)
	case "POST":
		createQuiz(w, r)
	case "PUT":
		updateQuiz(w, r)
	case "DELETE":
		deleteQuiz(w, r)
	default:
		w.WriteHeader(405)
	}
}

func showQuiz(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query()["id"][0])
	quiz, err := models.GetQuiz(id)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(quiz.ID)
}

func createQuiz(w http.ResponseWriter, r *http.Request) {
	quiz := models.Quiz{
		Image:       r.PostFormValue("image"),
		CorrectID:   1,
		CorrectRate: 1,
		Level:       1,
		CreatedAt:   time.Now(),
	}
	err := quiz.CreateQuiz()
	if err != nil {
		log.Fatalln(err)
	}
}

func updateQuiz(w http.ResponseWriter, r *http.Request) {

}

func deleteQuiz(w http.ResponseWriter, r *http.Request) {

}
