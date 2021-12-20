package postgres

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/abdullohsattorov/todo-service/config"
	pb "github.com/abdullohsattorov/todo-service/genproto"
	"github.com/abdullohsattorov/todo-service/pkg/db"
	"github.com/abdullohsattorov/todo-service/storage/repo"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TodoStorageI
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTodoRepo(pgPool)
	suite.CleanupFunc = cleanup
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *UserRepositoryTestSuite) TestTodoCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd4"
	assignee := "Assignee"

	todoFunc := pb.TodoFunc{
		Id:       id,
		Assignee: "Assignee",
		Title:    "Some Title",
		Summary:  "Summary",
		Deadline: "2020-12-10",
		Status:   "Active",
	}

	_ = suite.Repository.Delete(id)

	var todo pb.TodoFunc

	todoTodo, err := suite.Repository.Create(todoFunc)
	suite.Nil(err)

	todo = pb.TodoFunc{
		Id:       todoTodo.Id,
		Title:    todoTodo.Title,
		Assignee: todoTodo.Assignee,
		Summary:  todoTodo.Summary,
		Deadline: todoTodo.Deadline,
		Status:   todoTodo.Status,
	}

	getUser, err := suite.Repository.Get(todo.Id)
	suite.Nil(err)
	suite.NotNil(getUser, "todo must not be nil")
	suite.Equal(assignee, getUser.Assignee, "assignee must match")

	todo.Title = "Updated Title"
	updatedTodo, err := suite.Repository.Update(todo)
	suite.Nil(err)

	getUser, err = suite.Repository.Get(id)
	suite.Nil(err)
	suite.NotNil(getUser)
	suite.Equal(todo.Title, updatedTodo.Title)

	listTodos, _, err := suite.Repository.List(1, 10)
	suite.Nil(err)
	suite.NotEmpty(listTodos)
	suite.Equal(todo.Title, listTodos[0].Title)

	err = suite.Repository.Delete(id)
	suite.Nil(err)
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
