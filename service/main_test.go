package service

import (
	pb "github.com/abdullohsattorov/todo-service/genproto"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

var client pb.TodoServiceClient

func TestMain(m *testing.M) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect %v", err)
	}
	client = pb.NewTodoServiceClient(conn)

	os.Exit(m.Run())
}
