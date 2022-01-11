package main

import (
	"chat/models"

	"github.com/google/uuid"
)

const PubSubGeneralChannel = "general"

type WsServer struct {
	clients        map[*Client]bool
	register       chan *Client
	unregister     chan *Client
	broadcast      chan []byte
	rooms          map[*Room]bool
	users          []models.User
	roomRepository models.RoomRepository
	userRepository models.UserRepository
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer(roomRepository models.RoomRepository, userRepository models.UserRepository) *WsServer {
	wsServer := &WsServer{
		clients:        make(map[*Client]bool),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		rooms:          make(map[*Room]bool),
		roomRepository: roomRepository,
		userRepository: userRepository,
	}

	// Add users from database to server
	wsServer.users = userRepository.GetAllUsers()

	return wsServer
}

// Run our websocket server, accepting various requests
func (server *WsServer) Run() {
	go server.listenPubSubChannel()

	for {
		select {

		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)
		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *WsServer) registerClient(client *Client) {
	// First check if the user does not exist yet.
	if user := server.findUserByID(client.ID.String()); user == nil {
		// Add user to the repo
		server.userRepository.AddUser(client)
	}

	// Publish user in PubSub
	server.publishClientJoined(client)

	server.listOnlineClients(client)
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)

		// Publish user left in PubSub
		server.publishClientLeft(client)
	}
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *WsServer) findRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}

	// NEW: if there is no room, try to create it from the repo
	if foundRoom == nil {
		// Try to run the room from the repository, if it is found.
		foundRoom = server.runRoomFromRepository(name)
	}

	return foundRoom
}

// NEW: Try to find a room in the repo, if found Run it.
func (server *WsServer) runRoomFromRepository(name string) *Room {
	var room *Room
	dbRoom := server.roomRepository.FindRoomByName(name)
	if dbRoom != nil {
		room = NewRoom(dbRoom.GetName(), dbRoom.GetPrivate())
		room.ID, _ = uuid.Parse(dbRoom.GetId())

		go room.RunRoom()
		server.rooms[room] = true
	}

	return room
}

func (server *WsServer) createRoom(name string, private bool) *Room {
	room := NewRoom(name, private)
	// NEW: Add room to repo
	server.roomRepository.AddRoom(room)

	go room.RunRoom()
	server.rooms[room] = true

	return room
}

func (server *WsServer) listOnlineClients(client *Client) {
	// Find unique users instead of returning all users.
	var uniqueUsers = make(map[string]bool)
	for _, user := range server.users {
		if ok := uniqueUsers[user.GetId()]; !ok {
			message := &Message{
				Action: UserJoinedAction,
				Sender: user,
			}
			uniqueUsers[user.GetId()] = true
			client.send <- message.encode()
		}
	}
}

func (server *WsServer) findRoomByID(ID string) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetId() == ID {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (server *WsServer) findClientsByID(ID string) []*Client {
	// Find all clients for given user ID.
	var foundClients []*Client
	for client := range server.clients {
		if client.GetId() == ID {
			foundClients = append(foundClients, client)
		}
	}

	return foundClients
}
