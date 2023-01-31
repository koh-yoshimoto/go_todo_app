package service

import (
	"context"
	"fmt"

	"github.com/koh-yoshimoto/go_todo_app/entity"
	"github.com/koh-yoshimoto/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (lt *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := lt.Repo.ListTasks(ctx, lt.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}
