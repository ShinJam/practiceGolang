package main

import (
	"chat/auth"
	"chat/config"
	"chat/repository"
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	db := config.InitDB()
	defer db.Close()
	flag.Parse()
	config.CreateRedisClient()

	// Define the userRepo here, to use it in bothe the wsServer & the API
	userRepository := &repository.UserRepository{Db: db}

	wsServer := NewWebsocketServer(&repository.RoomRepository{Db: db}, userRepository)
	go wsServer.Run()

	api := &API{UserRepository: userRepository}
	// Add the login route
	http.HandleFunc("/api/login", api.HandleLogin)
	http.HandleFunc("/ws", auth.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	}))

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))

}
