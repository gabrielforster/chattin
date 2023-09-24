package main

import (
  "fmt"
  "context"
  "flag"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "auth/proto"
)

type server struct {
  pb.UnimplementedLoginServiceServer
  pb.UnimplementedSessionDataServiceServer
}

const(
  defaultToken = "thisisthetoken"
  defaultUsername = "gabriel"
  PORT = 50001
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
  listenner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := grpc.NewServer()

  pb.RegisterLoginServiceServer(s, &server{})
  pb.RegisterSessionDataServiceServer(s, &server{})

  log.Printf("Starting server on port %d", listenner.Addr())
  if err := s.Serve(listenner); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
