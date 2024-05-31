package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	TaskId      uuid.UUID `json:"id"`
	User        *User     `json:"user"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// функция создания карточки
func (db DB) CreateTask(task *Task) (*Task, error) {
	task.TaskId = uuid.New()

	conn, err := db.pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("невозможно получить соединение с базой данных: %v", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		"INSERT INTO task (taskid, userid, title, description) VALUES ($1, $2, $3, $4)",
		task.TaskId, task.User.UserId, task.Title, task.Description)

	return task, err
}

// функция удаления таски
func (db DB) DeleteTask(task *Task) (*Task, error) {
	conn, err := db.pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("невозможно получить соединение с базой данных: %v", err)
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		"DELETE FROM task WHERE taskid = $1 AND userid = $2",
		task.TaskId, task.User.UserId)

	return task, err
}

// функция получения всех тасок
func (db DB) GetAllTasks() ([]Task, error) {
	conn, err := db.pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("невозможно получить соединение с базой данных: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		`SELECT c.taskid, c.title, c.description, u.userid, u.username, u.avatar FROM task c
		LEFT JOIN users u ON c.userid = u.userid`)

	if err != nil {
		return nil, err
	}

	res := make([]Task, 0)
	for rows.Next() {
		c := Task{}
		u := User{}
		var userAvatar string
		err := rows.Scan(&c.TaskId, &c.Title, &c.Description, &u.UserId, &u.Username, &userAvatar)
		if err != nil {
			break
		}
		u.Avatar = &userAvatar
		c.User = &u
		res = append(res, c)
	}

	return res, err
}
