package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/abdullohsattorov/todo-service/genproto"
)

type todoRepo struct {
	db *sqlx.DB
}

// NewTodoRepo ...
func NewTodoRepo(db *sqlx.DB) *todoRepo {
	return &todoRepo{db: db}
}

func (r *todoRepo) Create(todo pb.TodoFunc) (pb.Todo, error) {
	var id string
	err := r.db.QueryRow(`
        INSERT INTO todos(id, assignee, title, summary, deadline, todo_status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id`, todo.Id, todo.Assignee, todo.Title, todo.Summary, todo.Deadline, todo.Status, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.Todo{}, err
	}

	var NewTodo pb.Todo

	NewTodo, err = r.Get(id)

	if err != nil {
		return pb.Todo{}, err
	}

	return NewTodo, nil
}

func (r *todoRepo) Get(id string) (pb.Todo, error) {
	var todo pb.Todo
	err := r.db.QueryRow(`
        SELECT id, assignee, title, summary, deadline, todo_status, created_at, updated_at FROM todos
        WHERE id=$1 and deleted_at is null`, id).Scan(&todo.Id, &todo.Assignee, &todo.Title, &todo.Summary, &todo.Deadline, &todo.Status, &todo.Created_At, &todo.Updated_At)
	if err != nil {
		return pb.Todo{}, err
	}

	return todo, nil
}

func (r *todoRepo) ListOverdue(req time.Time, page, limit int64) ([]*pb.Todo, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(`
				SELECT id, assignee, title, summary, deadline, todo_status, created_at, updated_at
				FROM todos WHERE deadline >= $1 and deleted_at is null order by id LIMIT $2 OFFSET $3`, req, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		todos []*pb.Todo
		count int64
	)
	for rows.Next() {
		var todo pb.Todo
		err = rows.Scan(
			&todo.Id,
			&todo.Assignee,
			&todo.Title,
			&todo.Summary,
			&todo.Deadline,
			&todo.Status,
			&todo.Created_At,
			&todo.Updated_At)
		if err != nil {
			return nil, 0, err
		}
		todos = append(todos, &todo)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM todos WHERE deadline >= $1 and deleted_at is null`, req).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}

func (r *todoRepo) List(page, limit int64) ([]*pb.Todo, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, todo_status, created_at, updated_at FROM todos WHERE deleted_at is null order by id LIMIT $1 OFFSET $2`,
		limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:err check

	var (
		todos []*pb.Todo
		count int64
	)
	for rows.Next() {
		var todo pb.Todo
		err = rows.Scan(&todo.Id, &todo.Assignee, &todo.Title, &todo.Summary, &todo.Deadline, &todo.Status, &todo.Created_At, &todo.Updated_At)
		if err != nil {
			return nil, 0, err
		}
		todos = append(todos, &todo)
	}

	err = r.db.QueryRow(`SELECT count(*) FROM todos where deleted_at is null`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return todos, count, nil
}

func (r *todoRepo) Update(todo pb.TodoFunc) (pb.Todo, error) {
	result, err := r.db.Exec(`UPDATE todos SET assignee=$1, title=$2, summary=$3, deadline=$4, todo_status=$5, updated_at = $6 WHERE id=$7 and deleted_at is null`,
		todo.Assignee, todo.Title, todo.Summary, todo.Deadline, todo.Status, time.Now().UTC(), todo.Id,
	)
	if err != nil {
		return pb.Todo{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Todo{}, sql.ErrNoRows
	}

	var NewTodo pb.Todo

	NewTodo, err = r.Get(todo.Id)

	if err != nil {
		return pb.Todo{}, err
	}

	return NewTodo, nil
}

func (r *todoRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE todos SET deleted_at = $1 WHERE id=$2`, time.Now().UTC(), id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}
