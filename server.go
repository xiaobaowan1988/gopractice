package main

import (
  "log"
  "net"
  "github.com/xiaobaowan1988/gopractice2/chat"
  "google.golang.org/grpc"
  "fmt"
)

func main() {
  fmt.Println("GoGO")

  lis, err := net.Listen("tcp",fmt.Sprintf(":%d",9000))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }

  s := chat.Server{}

  grpcServer := grpc.NewServer()
  chat.RegisterChatServiceServer(grpcServer, &s)

  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to server: %s", err)
  }
}
