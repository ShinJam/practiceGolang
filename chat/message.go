// message.go
package main

import (
	"chat/config"
	"chat/models"
	"encoding/json"
	"log"
)

const JoinRoomPrivateAction = "join-room-private"
const RoomJoinedAction = "room-joined"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Target  *Room       `json:"target"`
	Sender  models.User `json:"sender"`
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}

// UnmarshalJSON custom unmarshel to create a Client instance for Sender
func (message *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	msg := &struct {
		Sender Client `json:"sender"`
		*Alias
	}{
		Alias: (*Alias)(message),
	}
	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}
	message.Sender = &msg.Sender
	return nil
}

// Publish userJoined message in pub/sub
func (server *WsServer) publishClientJoined(client *Client) {

	message := &Message{
		Action: UserJoinedAction,
		Sender: client,
	}

	if err := config.Redis.Publish(ctx, PubSubGeneralChannel, message.encode()).Err(); err != nil {
		log.Println(err)
	}
}

// Publish userleft message in pub/sub
func (server *WsServer) publishClientLeft(client *Client) {

	message := &Message{
		Action: UserLeftAction,
		Sender: client,
	}

	if err := config.Redis.Publish(ctx, PubSubGeneralChannel, message.encode()).Err(); err != nil {
		log.Println(err)
	}
}

// Listen to pub/sub general channels
func (server *WsServer) listenPubSubChannel() {

	pubsub := config.Redis.Subscribe(ctx, PubSubGeneralChannel)
	ch := pubsub.Channel()
	for msg := range ch {

		var message Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("Error on unmarshal JSON message %s", err)
			return
		}

		switch message.Action {
		case UserJoinedAction:
			server.handleUserJoined(message)
		case UserLeftAction:
			server.handleUserLeft(message)
		case JoinRoomPrivateAction:
			server.handleUserJoinPrivate(message)
		}
	}
}
func (server *WsServer) handleUserJoinPrivate(message Message) {
	// Find client for given user, if found add the user to the room.
	// Expect multiple clients for one user now.
	targetClients := server.findClientsByID(message.Message)
	for _, targetClient := range targetClients {
		targetClient.joinRoom(message.Target.GetName(), message.Sender)
	}
}

// Add the findUserByID method used by client.go
func (server *WsServer) findUserByID(ID string) models.User {
	var foundUser models.User
	for _, client := range server.users {
		if client.GetId() == ID {
			foundUser = client
			break
		}
	}

	return foundUser
}

func (server *WsServer) handleUserJoined(message Message) {
	// Add the user to the slice
	server.users = append(server.users, message.Sender)
	server.broadcastToClients(message.encode())
}

func (server *WsServer) handleUserLeft(message Message) {
	for i, user := range server.users {
		if user.GetId() == message.Sender.GetId() {
			server.users[i] = server.users[len(server.users)-1]
			server.users = server.users[:len(server.users)-1]
			break // added this break to only remove the first occurrence
		}
	}

	server.broadcastToClients(message.encode())
}
