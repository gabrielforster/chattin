package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"net/http"

	"github.com/gorilla/websocket"
  "github.com/go-redis/redis/v8"
)

var configs = map[string]string{
  // change to get it from env
  "REDIS_ADDR": "localhost:6379",
  "PORT": os.Args[1],
}

var upgrader = websocket.Upgrader{
  ReadBufferSize: 2048,
  WriteBufferSize: 2048,
  CheckOrigin: func(r *http.Request) bool { return true },
}

var conns []*websocket.Conn

var ctx = context.Background()

var redisPub = redis.NewClient(&redis.Options{
  Addr: configs["REDIS_ADDR"],
})

var redisSub = redis.NewClient(&redis.Options{
  Addr: configs["REDIS_ADDR"],
})

func main() {
  go func() {
    sub := redisSub.Subscribe(ctx, "message")
    ch := sub.Channel()
    for msg := range ch {
      handleMessageFromRedis(msg)
    }
  }()

  http.HandleFunc("/", handlerWebsocketMessage)
  http.ListenAndServe(":" + configs["PORT"], nil)
}

func handleMessageFromRedis (message *redis.Message) {
    for _, conn := range conns {
      if err := conn.WriteMessage(1, []byte(message.Payload)); err != nil {
        fmt.Println(err)
        return
      }
    }
}

func handlerWebsocketMessage(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  socket, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  conns = append(conns, socket)

  for {
    _, msg, err := socket.ReadMessage()
    if err != nil {
      fmt.Println(err)
      return
    }

    command, message := string(msg[:bytes.IndexByte(msg, ' ')]), string(msg[bytes.IndexByte(msg, ' ')+1:])

    if (command == "message") {
      if err := redisPub.Publish(ctx, "message", message).Err(); err != nil {
        panic(err)
      }
    }
  }
}

