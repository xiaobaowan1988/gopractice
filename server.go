package main

import (
  "log"
  "net"
  "github.com/xiaobaowan1988/gopractice2/chat"
  "google.golang.org/grpc"
  "fmt"
  "context"
)


type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
  log.Printf("Receive message body from client: %s", in.Body)
  return    &chat.Message{Body: "Hello From the Server!"}, nil
}


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
