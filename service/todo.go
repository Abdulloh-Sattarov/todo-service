package service

import (
	"time"
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
func NewTodoService(db *sqlx.DB, log l.Logger) *TodoService {
	return &TodoService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *TodoService) Create(ctx context.Context, req *pb.Todo) (*pb.Todo, error) {
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

func (s *TodoService) Update(ctx context.Context, req *pb.Todo) (*pb.Todo, error) {
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

func (s *TodoService) ListOverdue(ctx context.Context, req *pb.Time) (*pb.Deadline, error) {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout , req.Time)
	if err != nil {
		s.logger.Error("Faild to convert time!")
		return nil, err
	}
	todo, err := s.storage.Todo().ListOverdue(t)
	if err != nil {
		s.logger.Error("failed to get todo", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get todo")
	}

	return &pb.Deadline{
		Todos: todo,
	}, nil
}
