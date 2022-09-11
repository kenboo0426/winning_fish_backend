package controllers

import (
	"net/http"
	"winning_fish_backend/config"
)

func StartMainServer() error {
	http.HandleFunc("/quizzes", indexQuiz)
	http.HandleFunc("/quiz", HandleQuizRequst)
	// ↑↓の順番注意
	http.HandleFunc("/quiz/", HandleQuizUpdateRequest)
	http.HandleFunc("/online_matches", indexOnlineMatch)
	http.HandleFunc("/socket", handleWebSocket)
	go ListenToWsChannel()
	http.HandleFunc("/online_match/start", startOnlineMatch)
	http.HandleFunc("/online_match", HandleOnlineMatchRequest)
	// ↑↓の順番注意
	http.HandleFunc("/online_match/", HandleOnlineMatchUpdateRequest)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
