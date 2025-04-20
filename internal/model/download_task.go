package model

import (
	"time"
	"database/sql"
)

type DownloadType string
type DownloadStatus string

const (
	DownloadTypeHTTP    DownloadType = "http"
	DownloadTypeFTP     DownloadType = "ftp"
	DownloadTypeTorrent DownloadType = "torrent"

	DownloadStatusQueued     DownloadStatus = "queued"
	DownloadStatusInProgress DownloadStatus = "in_progress"
	DownloadStatusCompleted  DownloadStatus = "completed"
	DownloadStatusFailed     DownloadStatus = "failed"
)

type DownloadTask struct {
	ID           int64          `db:"id"`
	UserID       int64          `db:"user_id"`
	DownloadType DownloadType   `db:"download_type"`
	URL          string         `db:"url"`
	Metadata     sql.NullString `db:"metadata"`    // store as raw JSON string
	Status       DownloadStatus `db:"status"`
	FilePath     *string        `db:"file_path"`   // nullable
	CreatedAt    time.Time      `db:"created_at"`
	UpdatedAt    time.Time      `db:"updated_at"`
}