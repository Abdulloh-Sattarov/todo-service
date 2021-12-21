package service

import (
	"context"
	pb "github.com/abdullohsattorov/todo-service/genproto"
	"reflect"
	"testing"
	"time"
)

func TestTodoService_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.TodoFunc
		want  pb.Todo
	}{
		{
			name: "successful",
			input: pb.TodoFunc{
				Assignee: "Test First Assignee",
				Title:    "Some Title",
				Summary:  "Summary",
				Deadline: "2021-10-15",
				Status:   "inactive",
			},
			want: pb.Todo{
				Assignee: "Test First Assignee",
				Title:    "Some Title",
				Summary:  "Summary",
				Deadline: "2021-10-15T00:00:00Z",
				Status:   "inactive",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to create todo", err)
			}
			got.Id = ""
			got.Created_At = ""
			got.Updated_At = ""
			got.Deleted_At = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoService_Get(t *testing.T) {
	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.Todo
	}{
		{
			name: "successful",
			input: pb.ByIdReq{
				Id: "0d512776-60ed-4980-b8a3-6904a2234fd5",
			},
			want: pb.Todo{
				Assignee: "Test First Assignee",
				Title:    "Some Title",
				Summary:  "Summary",
				Deadline: "2021-10-15T00:00:00Z",
				Status:   "inactive",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to get todo", err)
			}
			got.Id = ""
			got.Created_At = ""
			got.Updated_At = ""
			got.Deleted_At = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

type TestInput struct {
	time  time.Time
	page  int64
	limit int64
}

type TestWant struct {
	todos []pb.Todo
	count int64
}

func TestTodoService_List(t *testing.T) {
	tests := []struct {
		name  string
		input pb.ListReq
		want  TestWant
	}{
		{
			name: "successful",
			input: pb.ListReq{
				Page:  1,
				Limit: 2,
			},
			want: TestWant{
				todos: []pb.Todo{
					{
						Assignee: "Test First Assignee",
						Title:    "Some Title",
						Summary:  "Summary",
						Deadline: "2021-10-15T00:00:00Z",
						Status:   "inactive",
					},
					{
						Assignee: "Test First Assignee",
						Title:    "Some Title",
						Summary:  "Summary",
						Deadline: "2021-10-15T00:00:00Z",
						Status:   "inactive",
					},
				},
				count: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to get todo", err)
			}

			if got.Count == tc.want.count {
				for i, j := range tc.want.todos {
					if j.Assignee != got.Todos[i].Assignee || j.Title != got.Todos[i].Title || j.Summary != got.Todos[i].Summary || j.Deadline != got.Todos[i].Deadline || j.Status != got.Todos[i].Status {
						t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want.todos, got)
					}
				}
			} else {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want.todos, got)
			}
		})
	}
}

func TestTodoService_Update(t *testing.T) {
	tests := []struct {
		name  string
		input pb.TodoFunc
		want  pb.Todo
	}{
		{
			name: "successful",
			input: pb.TodoFunc{
				Id:       "0d512776-60ed-4980-b8a3-6904a2234fd9",
				Assignee: "Updated1",
				Title:    "Updated",
				Summary:  "Updated",
				Deadline: "2021-12-14",
				Status:   "Updated",
			},
			want: pb.Todo{
				Assignee: "Updated1",
				Title:    "Updated",
				Summary:  "Updated",
				Deadline: "2021-12-14T00:00:00Z",
				Status:   "Updated",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to get todo", err)
			}
			got.Id = ""
			got.Created_At = ""
			got.Updated_At = ""
			got.Deleted_At = ""
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTodoService_Delete(t *testing.T) {
	tests := []struct {
		name  string
		input pb.ByIdReq
		want  error
	}{
		{
			name: "successful",
			input: pb.ByIdReq{
				Id: "cbeb644f-bad4-4dc6-8e23-dd9b9f7bb180",
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.Delete(context.Background(), &tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, err)
			}
		})
	}
}

func TestTodoService_ListOverdue(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Time
		want  TestWant
	}{
		{
			name: "successful",
			input: pb.Time{
				Time:  "2000-12-12",
				Page:  1,
				Limit: 5,
			},
			want: TestWant{
				todos: []pb.Todo{
					{
						Assignee: "Test First Assignee",
						Title:    "Some Title",
						Summary:  "Summary",
						Deadline: "2021-10-15T00:00:00Z",
						Status:   "inactive",
					},
					{
						Assignee: "Test First Assignee",
						Title:    "Some Title",
						Summary:  "Summary",
						Deadline: "2021-10-15T00:00:00Z",
						Status:   "inactive",
					},
					{
						Assignee: "Updated1",
						Title:    "Updated",
						Summary:  "Updated",
						Deadline: "2021-12-14T00:00:00Z",
						Status:   "Updated",
					},
				},
				count: 3,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.ListOverdue(context.Background(), &tc.input)
			if err != nil {
				t.Error("Failed to get todo", err)
			}

			if got.Count == tc.want.count {
				for i, j := range tc.want.todos {
					if j.Assignee != got.Todos[i].Assignee || j.Title != got.Todos[i].Title || j.Summary != got.Todos[i].Summary || j.Deadline != got.Todos[i].Deadline || j.Status != got.Todos[i].Status {
						t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want.todos, got)
					}
				}
			} else {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want.todos, got)
			}
		})
	}
}
