package handler

import (
	"context"

	"github.com/koh-yoshimoto/go_todo_app/entity"
)

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskServiceinterface interface {
	AddTask(ctx context.Context, title string) (entity.Task, error)
}
