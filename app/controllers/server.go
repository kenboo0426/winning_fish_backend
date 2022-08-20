package controllers

import (
	"net/http"
	"winning_fish_backend/config"
)

func StartMainServer() error {
	http.HandleFunc("/quizzes", indexQuiz)
	http.HandleFunc("/quiz", HandleQuizRequst)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
