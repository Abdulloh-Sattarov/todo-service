package repo

import (
	"time"

	pb "github.com/abdullohsattorov/todo-service/genproto"
)

// TodoStorageI ...
type TodoStorageI interface {
	Create(todoFunc pb.TodoFunc) (pb.Todo, error)
	Get(id string) (pb.Todo, error)
	List(page, limit int64) ([]*pb.Todo, int64, error)
	Update(update pb.TodoFunc) (pb.Todo, error)
	Delete(id string) error
	ListOverdue(time time.Time, page, limit int64) ([]*pb.Todo, int64, error)
}
