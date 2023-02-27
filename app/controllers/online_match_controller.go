package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"winning_fish_backend/app/models"
)

func HandleOnlineMatchUpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodGet:
		showOnlineMatch(w, r)
	case http.MethodPut:
		updateOnlineMatch(w, r)
	case http.MethodPost:
		s := r.URL.Path[len("/online_match/"):]
		if strings.Contains(s, "/") && strings.Contains(s, "join") {
			joinOnlineMatch(w, r)
		} else {
			w.WriteHeader(405)
		}
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func joinOnlineMatch(w http.ResponseWriter, r *http.Request) {
	user_or_guest_id := r.FormValue("user_or_guest_id")
	s := strings.Trim(r.URL.Path, "/online_match/")
	onlineMatchID, err := strconv.Atoi(strings.Trim(s, "/join"))
	if err != nil {
		w.WriteHeader(401)
		return
	}
	var online_match_joined_user models.OnlineMatchJoinedUser
	err = online_match_joined_user.CreateOnlineMatchJoinedUser(user_or_guest_id, onlineMatchID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	onlineMatch, err := models.GetOnlineMatchByID(onlineMatchID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(onlineMatch)
	w.Write(res)
}

func showOnlineMatch(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/online_match")
	_, id := filepath.Split(sub)

	online_match_id, _ := strconv.Atoi(id)
	onlineMatch, err := models.GetOnlineMatchByID(online_match_id)
	if err != nil {
		log.Fatalln(err)
	}

	onlineMatch.OnlineMatchJoinedUsers, err = onlineMatch.GetJoinedUsersByOnlineMatch()
	onlineMatch.OnlineMatchAskedQuizzes, err = onlineMatch.GetAskedQuizByOnlineMatch()
	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(onlineMatch)
	w.Write(res)
}

func updateOnlineMatch(w http.ResponseWriter, r *http.Request) {
}

func startOnlineMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	sub := strings.TrimPrefix(r.URL.Path, "/online_match/start")
	_, id := filepath.Split(sub)

	online_match_id, _ := strconv.Atoi(id)
	online_match, err := models.GetOnlineMatchByID(online_match_id)
	users, _ := online_match.GetJoinedUsersByOnlineMatch()

	if err != nil {
		w.WriteHeader(404)
		return
	}
	if len(users) > 4 || len(users) <= 1 {
		w.WriteHeader(404)
		return
	}
	online_match.Status = "processing"
	err = online_match.UpdateOnlineMatch()
	online_match.RegisterQuiz()

	if err != nil {
		w.WriteHeader(404)
		return
	}

	res, _ := json.Marshal(online_match)
	w.Write(res)
}

func searchOnlineMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	sub := strings.TrimPrefix(r.URL.Path, "/online_match/search/")
	_, roomID := filepath.Split(sub)
	onlineMatch, err := models.GetOnlineMatchByRoomID(roomID)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	res, _ := json.Marshal(onlineMatch)
	w.Write(res)
}

func finishOnlineMatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	sub := strings.TrimPrefix(r.URL.Path, "/online_match/finish")
	_, id := filepath.Split(sub)

	online_match_id, _ := strconv.Atoi(id)
	online_match, err := models.GetOnlineMatchByID(online_match_id)
	if err != nil {
		log.Fatalln(err)
	}
	online_match.Status = "finished"
	finished_at := time.Now()
	online_match.FinishedAt = &finished_at
	err = online_match.UpdateOnlineMatch()
	online_match.OnlineMatchJoinedUsers, err = online_match.GetJoinedUsersByOnlineMatch()
	if err != nil {
		log.Fatalln(err)
	}
	online_match.OnlineMatchAskedQuizzes, err = online_match.GetAskedQuizByOnlineMatch()

	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(online_match)
	w.Write(res)
}
