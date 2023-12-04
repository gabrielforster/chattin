package main

import (
  "fmt"
  "context"
  "flag"
  "log"
  "net"
  "net/http"

  "github.com/labstack/echo/v4"
  "google.golang.org/grpc"
  pb "auth/proto"
  database "auth/database"
)

type server struct {
  pb.UnimplementedLoginServiceServer
  pb.UnimplementedSessionDataServiceServer
}

type RegisterResponse struct {
  Username string `json:"username"`
  Email string `json:"email"`
  Path string `json:"path"`
}

const(
  PORT = 50001
  HTTP_PORT = 50002
)

var (
  port = flag.Int("port", PORT, "the port to connect to")
  host = fmt.Sprintf("localhost:%d", PORT) 
  addr = flag.String("addr", host, "the address to connect to")
)

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
  log.Printf(
    "Received username: %v, Received password %v",
    in.GetUsername(),
    in.GetPassword(),
  )

  return &pb.LoginResponse{Token: defaultToken}, nil
}

func (s *server) GetSessionData (ctx context.Context, in *pb.GetSessionDataRequest) (*pb.GetSessionDataResponse, error) {
  log.Printf(
    "Received token: %v",
    in.GetToken(),
  )

  return &pb.GetSessionDataResponse{Username: defaultUsername, ProfileUrl: "dsahjkdghjhkashdjkha"}, nil
}

func main () {
  flag.Parse()
  
  // creating database connection
  _, err := database.NewDatabase()
  if err != nil {
    log.Fatalf("could not connect to database: %v", err)
  }

  // init gRPC server 
  listenner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := grpc.NewServer()

  pb.RegisterLoginServiceServer(s, &server{})
  pb.RegisterSessionDataServiceServer(s, &server{})

  e := echo.New()

  e.GET("/register", func(c echo.Context) error {
    return c.JSON(http.StatusOK, &RegisterResponse{
      Username: "gabriel",
      Email: "rochafrgabriel@gmail.com",
      Path: "/profile/gabriel",
    })
  })

  // e.GET("/profile/:username", func(c echo.Context) error {
  //   username := c.Param("username")
  //   return c.JSON(http.StatusOK, &profileResponse{
  //     Username: username,
  //   })
  // })

  go func () {
    if err := e.Start(fmt.Sprintf(":%d", HTTP_PORT)); err != nil {
      e.Logger.Info("could not listen, shutting down the server")
    } else {
      e.Logger.Info("starting the server")
    }
  } ()

  log.Printf("Starting gRPC server on port %d", listenner.Addr())
  if err := s.Serve(listenner); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}

