package model

import (
	"database/sql"
	"time"
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
	ID           uint64
	UserID       uint64
	DownloadType DownloadType
	URL          string
	Metadata     sql.NullString
	Status       DownloadStatus
	FilePath     *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
