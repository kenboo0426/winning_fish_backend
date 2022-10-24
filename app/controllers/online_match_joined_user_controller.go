package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"winning_fish_backend/app/models"
)

func HandleOnlineMatchRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodPost:
		joinOrCreateOnlineMatch(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

func joinOrCreateOnlineMatch(w http.ResponseWriter, r *http.Request) {
	user_or_guest_id := r.FormValue("user_or_guest_id")
	onlineMatch, err := models.GetJoinableOnlineMatch()
	var online_match_joined_user models.OnlineMatchJoinedUser
	if err == nil {
		online_match_joined_user.CreateOnlineMatchJoinedUser(user_or_guest_id, onlineMatch.ID)
	} else {
		onlineMatch.PersonNumber = 0
		onlineMatch.ParticipantsNumber = 4
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

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
