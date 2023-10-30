package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 2048,
  WriteBufferSize: 2048,
  CheckOrigin: func(r *http.Request) bool { return true },
}

var conns []*websocket.Conn

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":42069", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  socket, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  conns = append(conns, socket)

  for {
    msgType, msg, err := socket.ReadMessage()
    if err != nil {
      fmt.Println(err)
      return
    }

    command, message := string(msg[:bytes.IndexByte(msg, ' ')]), string(msg[bytes.IndexByte(msg, ' ')+1:])

    if (command == "message") {
      for _, conn := range conns {
        if err := conn.WriteMessage(msgType, []byte(message)); err != nil {
          fmt.Println(err)
          return
        }
      }
    }
  }
}

