package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)
var (
	WSPort=":322"
)
type Client struct{
	ID string
	mu *sync.RWMutex
	conn *websocket.Conn

}
func newClient(conn *websocket.Conn)*Client{
	ID:=rand.Text()[:9]
	return &Client{
		ID:ID, 
		mu: new(sync.RWMutex),
		conn:conn, 

	}


}
type Server struct{
	clients []*Client
	mu *sync.RWMutex
}
//TODO
//[]Http server
//[]upgrade it to ws once client connects
//[]add newly connected ws to server
//[]add client to server
//[]remove client on disconnect
//send broadcast msg -> no race condition
func NewServer()*Server{
	return &Server{
		clients: []*Client{},
		mu:new(sync.RWMutex), 
	}
}
func (s *Server) handleWs( w http.ResponseWriter, r *http.Request){
	upgrader:=websocket.Upgrader{
		ReadBufferSize: 512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		fmt.Printf("Error on HTTP conn upgrade %v\n", err)
			return 
	}
	//add conn to server
	client:=newClient(conn)
	s.clients=append(s.clients, client)
	//read msg loop



}

func main(){
	s:=NewServer()
http.HandleFunc("/",s.handleWs)
log.Fatal(http.ListenAndServe(WSPort, nil))
}