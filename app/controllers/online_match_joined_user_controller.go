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
	"winning_fish_backend/app/models"
)

func HandleOnlineMatchRequest(w http.ResponseWriter, r *http.Request) {
	permitCors(w, r)

	switch r.Method {
	case http.MethodPost:
		createOnlineMatch(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func createOnlineMatch(w http.ResponseWriter, r *http.Request) {
	var online_match models.OnlineMatch
	defer r.Body.Close()

	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &online_match); err != nil {
		fmt.Println(err)
	}
	online_match.Status = "opening"
	err := online_match.CreateOnlineMatch()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	res, _ := json.Marshal(online_match)
	w.WriteHeader(201)
	w.Write(res)
}

func joinOrCreateOnlineMatch(w http.ResponseWriter, r *http.Request) {
	permitCors(w, r)
	user_or_guest_id := r.FormValue("user_or_guest_id")
	onlineMatch, err := models.GetJoinableOnlineMatch()
	var online_match_joined_user models.OnlineMatchJoinedUser
	if err == nil {
		online_match_joined_user.CreateOnlineMatchJoinedUser(user_or_guest_id, onlineMatch.ID)
	} else {
		onlineMatch.MaxParticipateNumber = 4
		onlineMatch.Status = "opening"
		err = onlineMatch.CreateOnlineMatch()
		online_match_joined_user.CreateOnlineMatchJoinedUser(user_or_guest_id, onlineMatch.ID)
	}
	onlineMatch.OnlineMatchJoinedUsers, err = onlineMatch.GetJoinedUsersByOnlineMatch()

	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(onlineMatch)
	w.Write(res)
}

func calculateOnlineMatch(w http.ResponseWriter, r *http.Request) {
	permitCors(w, r)

	joined_user_id := r.FormValue("joined_user_id")
	sub := strings.TrimPrefix(r.URL.Path, "/online_match/calculate")
	_, id := filepath.Split(sub)
	online_match_id, _ := strconv.Atoi(id)

	online_match_joined_user, err := models.GetJoinedUsersByOnlineMatchAndUserID(online_match_id, joined_user_id)

	if err != nil {
		log.Fatalln(err)
	}
	err = online_match_joined_user.CalculateRemainedTimeByOnlineMatchID()
	online_match, err := models.GetOnlineMatch(online_match_id)
	if err != nil {
		log.Fatalln(err)
	}
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

func permitCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
}
