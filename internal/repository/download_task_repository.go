package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/aq2208/goload/internal/model"
)

type DownloadTaskRepository interface {
	CreateDownloadTask(ctx context.Context, tx *sql.Tx, task model.DownloadTask) (uint64, error)
	UpdateDownloadTask(ctx context.Context, task model.DownloadTask) error
	UpdateStatusDownloadTask(ctx context.Context, id int64, status string) error
	DeleteDownloadTask(ctx context.Context, taskId uint64) error
	GetDownloadTaskListOfUser(ctx context.Context, userId, offset, limit uint64) ([]model.DownloadTask, error)
	GetDownloadTaskCountOfUser(ctx context.Context, userId uint64) (uint64, error)
}

type downloadTaskRepository struct {
	db *sql.DB
}

// CreateDownloadTask implements DownloadTaskRepository.
func (d *downloadTaskRepository) CreateDownloadTask(ctx context.Context, tx *sql.Tx, task model.DownloadTask) (uint64, error) {
	query := `INSERT INTO download_task (user_id, download_type, url, metadata, status)
	          VALUES (?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, query, task.UserID, task.DownloadType, task.URL, task.Metadata, task.Status)
	if err != nil {
		return 0, err
	}

	insertId, _ := result.LastInsertId()
	log.Default().Printf("Inserted user %d successful", insertId)

	return uint64(insertId), nil
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

func (r *downloadTaskRepository) UpdateStatusDownloadTask(ctx context.Context, id int64, status string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE download_task SET status = ? WHERE id = ?", status, id)
	return err
}

func NewDownloadTaskRepository(db *sql.DB) DownloadTaskRepository {
	return &downloadTaskRepository{db: db}
}
