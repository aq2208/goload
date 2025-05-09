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
	UpdateStatusAndMetadataDownloadTask(ctx context.Context, id uint64, status string, metadata string) error
	DeleteDownloadTask(ctx context.Context, taskId uint64) error
	GetDownloadTaskListOfUser(ctx context.Context, userId, offset, limit uint64) ([]model.DownloadTask, error)
	GetDownloadTaskCountOfUser(ctx context.Context, userId uint64) (uint64, error)
	GetDownloadTaskById(ctx context.Context, id uint64) (model.DownloadTask, error)
	GetPendingTasks(ctx context.Context) ([]model.DownloadTask, error)
}

type downloadTaskRepository struct {
	db *sql.DB
}

// GetPendingTasks implements DownloadTaskRepository.
func (d *downloadTaskRepository) GetPendingTasks(ctx context.Context) ([]model.DownloadTask, error) {
	var records []model.DownloadTask

	rows, err := d.db.Query("SELECT id, user_id, download_type, url, status FROM download_task WHERE status in ('queued, failed')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var record model.DownloadTask

		if err := rows.Scan(&record.ID, &record.UserID, &record.DownloadType, &record.URL, &record.Status); err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

// GetDownloadTaskById implements DownloadTaskRepository.
func (d *downloadTaskRepository) GetDownloadTaskById(ctx context.Context, id uint64) (model.DownloadTask, error) {
	var dt model.DownloadTask

	row := d.db.QueryRow("SELECT id, user_id, download_type, url, status FROM download_task WHERE id = ?", id)
	if err := row.Scan(&dt.ID, &dt.UserID, &dt.DownloadType, &dt.URL, &dt.Status); err != nil {
		log.Default().Printf("Error query download_task: %v", err)
		return model.DownloadTask{}, err
	}

	return dt, nil
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

func (r *downloadTaskRepository) UpdateStatusAndMetadataDownloadTask(ctx context.Context, id uint64, status string, metadata string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE download_task SET status = ?, metadata = ? WHERE id = ?", status, metadata, id)
	return err
}

func NewDownloadTaskRepository(db *sql.DB) DownloadTaskRepository {
	return &downloadTaskRepository{db: db}
}
