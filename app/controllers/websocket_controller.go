package controllers

import (
	"fmt"
	"log"
	"net/http"
	"winning_fish_backend/app/models"

	"github.com/gorilla/websocket"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	wsChan  = make(chan models.WsPayload)
	clients = make(map[models.WebSocketConnection]models.WsUser)
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("ok client connectiong")

	var response models.WsJsonResponse

	conn := models.WebSocketConnection{Conn: ws}
	clients[conn] = models.WsUser{}

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

		if err != nil {
			// log.Println(err)
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func broadcastToAll(response models.WsJsonResponse) {
	for client := range clients {
		fmt.Println(response, client)
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
		case "start_online_match":
			users := getUserList()
			response.Users = users
			response.Action = "start_online_match"
		case "fetch_joined_user":
			users := getUserList()
			response.Action = "list_users"
			response.Users = users
		case "finished_online_match":
			var userID string
			userID = string(e.UserID)
			if !include(clients, userID) {
				clients[e.Conn] = models.WsUser{ID: userID, Name: e.UserName, RemainedTime: &e.RemainedTime}
			}
			users := getUserList()
			response.Action = "finished_online_match"
			response.Users = users
		case "join_online_match":
			var userID string
			userID = string(e.UserID)

			if !include(clients, userID) {
				clients[e.Conn] = models.WsUser{ID: userID, Name: e.UserName, Icon: e.UserIcon}
			}
			users := getUserList()
			response.Action = "list_users"
			response.Users = users
		case "left":
			delete(clients, e.Conn)
			user_ids := getUserList()
			response.Action = "list_users"
			response.Users = user_ids
		}

		broadcastToAll(response)
	}
}

func getUserList() []models.WsUser {
	var clientList []models.WsUser
	for _, client := range clients {
		fmt.Printf("p :%+v\n", client)
		if client.ID != "" {
			clientList = append(clientList, client)
		}
	}

	return clientList
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

func include(array map[models.WebSocketConnection]models.WsUser, target string) bool {
	for _, item := range array {
		if item.ID == target {
			return true
		}
	}

	return false
}
