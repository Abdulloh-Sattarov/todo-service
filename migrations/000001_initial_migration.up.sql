CREATE TABLE IF NOT EXISTS todos(
    id SERIAL Primary Key,
    assignee VARCHAR(50),
    title VARCHAR(50),
    summary VARCHAR(50),
    deadline  timestamp not null,
    todo_status VARCHAR(50)
);