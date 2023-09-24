package main

import (
  "fmt"
  "context"
  "flag"
  "log"
  "time"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "auth/proto"
)

const (
  username = "gabriel"
  password = "123456"
)

var (
  PORT = flag.String("port", "50001", "the port to connect to")
  host = fmt.Sprintf("localhost:%s", *PORT)
  address = flag.String("address", host, "the address to connect to")
)

func main () {
 flag.Parse()

 conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
 if err != nil {
  log.Fatalf("did not connect: %v", err)
 }
 defer conn.Close()

 cLogin := pb.NewLoginServiceClient(conn)

 loginCtx, loginCancel := context.WithTimeout(context.Background(), time.Second)
 defer loginCancel()

 r, err := cLogin.Login(loginCtx, &pb.LoginRequest{Username: username, Password: password})
 if err != nil {
  log.Fatalf("could not login: %v", err)
 }
 log.Printf("Login result: %s", r.GetToken())

 cSession := pb.NewSessionDataServiceClient(conn)

 fmt.Println("cSession ok")

 sessionCtx, sessionCancel := context.WithTimeout(context.Background(), time.Second)
 defer sessionCancel()

 fmt.Println("sessionCtx ok")

 sessionResponse, err := cSession.GetSessionData(sessionCtx, &pb.GetSessionDataRequest{Token: r.GetToken()})
 fmt.Printf("sessionResponse %v", sessionResponse)
 if err != nil {
  log.Fatalf("could not get session data: %v", err)
 }
 log.Printf("GetSessionData result: %s", sessionResponse.GetUsername())
}
