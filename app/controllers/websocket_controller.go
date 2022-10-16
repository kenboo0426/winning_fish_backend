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
	wsChan  = make(chan models.WsRequest)
	clients = make(map[models.WsConnection]models.WsClient)
)

func startWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("ok client connectiong")

	var response models.WsResponse

	conn := models.WsConnection{Conn: ws}
	clients[conn] = models.WsClient{}

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenFowWs(&conn)
}

func ListenFowWs(conn *models.WsConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var request models.WsRequest

	for {
		err := conn.ReadJSON(&request)

		if err != nil {
			break
		} else {
			request.Conn = *conn
			wsChan <- request
		}
	}
}

func ListenToWsChannel() {
	var response models.WsResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "start_online_match":
			users := getUserList(e.OnlineMatchID)
			response.Users = users
			response.Action = "start_online_match"
		case "fetch_joined_user":
			users := getUserList(e.OnlineMatchID)
			response.Action = "list_users"
			response.Users = users
		case "finished_online_match":
			var userID string
			userID = string(e.UserID)
			if !include(clients, userID) {
				clients[e.Conn] = models.WsClient{
					WsUser: models.WsUser{
						ID:           userID,
						Name:         e.UserName,
						RemainedTime: &e.RemainedTime,
						Icon:         e.UserIcon,
					},
					OnlineMatchID: e.OnlineMatchID,
				}
			}
			users := getUserList(e.OnlineMatchID)
			response.Action = "finished_online_match"
			response.Users = users
		case "join_online_match":
			var userID string
			userID = string(e.UserID)
			if !include(clients, userID) {
				clients[e.Conn] = models.WsClient{
					WsUser: models.WsUser{
						ID:           userID,
						Name:         e.UserName,
						RemainedTime: &e.RemainedTime,
						Icon:         e.UserIcon,
					},
					OnlineMatchID: e.OnlineMatchID,
				}
			}
			users := getUserList(e.OnlineMatchID)
			response.Action = "list_users"
			response.Users = users
		case "left":
			delete(clients, e.Conn)
			user_ids := getUserList(e.OnlineMatchID)
			response.Action = "list_users"
			response.Users = user_ids
		}
		response.OnlineMatchID = e.OnlineMatchID

		broadcastToAll(response)
	}
}

func broadcastToAll(response models.WsResponse) {
	for conn, client := range clients {
		if client.OnlineMatchID != response.OnlineMatchID {
			continue
		}

		err := conn.WriteJSON(response)
		if err != nil {
			log.Println(err)
			_ = conn.Close()
			delete(clients, conn)
		}
	}
}

func getUserList(online_match_id int) []models.WsUser {
	var clientList []models.WsUser
	for _, client := range clients {
		if client.WsUser.ID != "" && client.OnlineMatchID == online_match_id {
			clientList = append(clientList, client.WsUser)
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

func include(array map[models.WsConnection]models.WsClient, target string) bool {
	for _, item := range array {
		if item.WsUser.ID == target {
			return true
		}
	}

	return false
}
