package controllers

import (
	"net/http"
	"winning_fish_backend/config"
)

func StartMainServer() error {
	http.HandleFunc("/quizzes", indexQuiz)
	http.HandleFunc("/quiz", HandleQuizRequst)
	// http.HandleFunc("/quiz/delete/", deleteQuiz)
	// ↑↓の順番注意
	http.HandleFunc("/quiz/", HandleQuizUpdateRequest)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
