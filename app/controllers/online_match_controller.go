package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"winning_fish_backend/app/models"

	"github.com/gorilla/websocket"
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

func HandleOnlineMatchUpdateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))

	switch r.Method {
	case http.MethodGet:
		showOnlineMatch(w, r)
	case http.MethodPut:
		updateOnlineMatch(w, r)
	case http.MethodOptions:
		// handle preflight here
	default:
		w.WriteHeader(405)
	}
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	wsChan  = make(chan models.WsPayload)
	clients = make(map[models.WebSocketConnection]string)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("ok client connectiong")

	var response models.WsJsonResponse

	conn := models.WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenFowWs(&conn)
}

func ListenFowWs(conn *models.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload models.WsPayload

	for {
		err := conn.ReadJSON(&payload)
		fmt.Println(payload, "payload")

		if err != nil {
			log.Println(err)
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func broadcastToAll(response models.WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println(err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func ListenToWsChannel() {
	var response models.WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "fetch_joined_user":
			user_ids := getUserList()
			response.Action = "list_users"
			response.JoinedOnlineMatchUserIDs = user_ids
		case "join_online_match":
			var userID string
			if e.UserID == "" {
				// var user models.User
				// user.UUID = models.CreateUUID().String()
				// user.Name = "guests"
				// user.Role = 2
				// user.CreateUser()
				// userID = string(user.ID)
			} else {
				userID = string(e.UserID)
			}

			if !include(clients, userID) {
				clients[e.Conn] = userID
			}
			user_ids := getUserList()
			response.Action = "list_users"
			response.JoinedOnlineMatchUserIDs = user_ids
		case "left":
			delete(clients, e.Conn)
			user_ids := getUserList()
			response.Action = "list_users"
			response.JoinedOnlineMatchUserIDs = user_ids
		}
		fmt.Println(response, "resposne")

		broadcastToAll(response)
	}
}

func getUserList() []string {
	var clientList []string
	for _, client := range clients {
		if client != "" {
			clientList = append(clientList, client)
		}
	}

	sort.Strings(clientList)
	return clientList
}

func indexOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func joinOrCreateOnlineMatch(w http.ResponseWriter, r *http.Request) {
	user_id := r.FormValue("user_id")
	onlineMatch, err := models.GetJoinableOnlineMatch()
	var online_match_joined_user models.OnlineMatchJoinedUser
	if err == nil {
		online_match_joined_user.CreateOnlineMatchJoinedUser(user_id, onlineMatch.ID)
	} else {
		err = onlineMatch.CreateOnlineMatch()
		online_match_joined_user.CreateOnlineMatchJoinedUser(user_id, onlineMatch.ID)
	}
	onlineMatch.OnlineMatchJoinedUsers, err = onlineMatch.GetJoinedUsersByOnlineMatch()

	defer r.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(onlineMatch)
	w.Write(res)
}

func showOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func updateOnlineMatch(w http.ResponseWriter, r *http.Request) {

}

func startOnlineMatch(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/quiz")
	_, id := filepath.Split(sub)

	online_match_id, _ := strconv.Atoi(id)
	online_match, err := models.GetOnlineMatch(online_match_id)
	if err != nil {
		log.Fatalln(err)
	}
	online_match.Status = "processing"
	err = online_match.UpdateOnlineMatch()

	if err != nil {
		log.Fatalln(err)
	}

	res, _ := json.Marshal(online_match)
	w.Write(res)
}

func uniq(target []string) (result []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}

	return result
}

func include(array map[models.WebSocketConnection]string, target string) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}

	return false
}
