package service

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gofrs/uuid"

	pb "github.com/abdullohsattorov/todo-service/genproto"
	l "github.com/abdullohsattorov/todo-service/pkg/logger"
	"github.com/abdullohsattorov/todo-service/storage"
)

// TodoService ...
type TodoService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTodoService ...
func NewTodoService(storage storage.IStorage, log l.Logger) *TodoService {
	return &TodoService{
		storage: storage,
		logger:  log,
	}
}

func (s *TodoService) Create(ctx context.Context, req *pb.TodoFunc) (*pb.Todo, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}

	req.Id = id.String()

	todo, err := s.storage.Todo().Create(*req)
	if err != nil {
		s.logger.Error("failed to create todo", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create todo")
	}

	return &todo, nil
}

func (s *TodoService) Get(ctx context.Context, req *pb.ByIdReq) (*pb.Todo, error) {
	todo, err := s.storage.Todo().Get(req.GetId())
	if err != nil {
		s.logger.Error("failed to get todo", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get todo")
	}

	return &todo, nil
}

func (s *TodoService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	todos, count, err := s.storage.Todo().List(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list todos", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list todos")
	}

	return &pb.ListResp{
		Todos: todos,
		Count: count,
	}, nil
}

func (s *TodoService) Update(ctx context.Context, req *pb.TodoFunc) (*pb.Todo, error) {
	todo, err := s.storage.Todo().Update(*req)
	if err != nil {
		s.logger.Error("failed to update todo", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update todo")
	}

	return &todo, nil
}

func (s *TodoService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.Todo().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete todo", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete todo")
	}

	return &pb.EmptyResp{}, nil
}

func (s *TodoService) ListOverdue(ctx context.Context, req *pb.Time) (*pb.ListResp, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, req.Time)
	if err != nil {
		s.logger.Error("Failed to convert time!")
		return nil, err
	}
	todos, count, err := s.storage.Todo().ListOverdue(t, req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list todos", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list todos")
	}

	return &pb.ListResp{
		Todos: todos,
		Count: count,
	}, nil
}
