//main.go
package main

import (
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

	wsServer := NewWebsocketServer(&repository.RoomRepository{Db: db}, &repository.UserRepository{Db: db})
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))

}
