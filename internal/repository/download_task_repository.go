package repository

import (
	"context"
	"database/sql"

	"github.com/aq2208/goload/internal/model"
)

type DownloadTaskRepository interface {
	CreateDownloadTask(ctx context.Context, task model.DownloadTask) (uint64, error)
	UpdateDownloadTask(ctx context.Context, task model.DownloadTask) error
	DeleteDownloadTask(ctx context.Context, taskId uint64) error
	GetDownloadTaskListOfUser(ctx context.Context, userId, offset, limit uint64) ([]model.DownloadTask, error)
	GetDownloadTaskCountOfUser(ctx context.Context, userId uint64) (uint64, error)
}

type downloadTaskRepository struct {
	db *sql.DB
}

// CreateDownloadTask implements DownloadTaskRepository.
func (d *downloadTaskRepository) CreateDownloadTask(ctx context.Context, task model.DownloadTask) (uint64, error) {
	panic("unimplemented")
}

// DeleteDownloadTask implements DownloadTaskRepository.
func (d *downloadTaskRepository) DeleteDownloadTask(ctx context.Context, taskId uint64) error {
	panic("unimplemented")
}

// GetDownloadTaskCountOfUser implements DownloadTaskRepository.
func (d *downloadTaskRepository) GetDownloadTaskCountOfUser(ctx context.Context, userId uint64) (uint64, error) {
	panic("unimplemented")
}

// GetDownloadTaskListOfUser implements DownloadTaskRepository.
func (d *downloadTaskRepository) GetDownloadTaskListOfUser(ctx context.Context, userId uint64, offset uint64, limit uint64) ([]model.DownloadTask, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements DownloadTaskRepository.
func (d *downloadTaskRepository) UpdateDownloadTask(ctx context.Context, task model.DownloadTask) error {
	panic("unimplemented")
}

func NewDownloadTaskRepository(db *sql.DB) DownloadTaskRepository {
	return &downloadTaskRepository{db: db}
}
