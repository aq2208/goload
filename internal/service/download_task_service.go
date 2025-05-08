package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aq2208/goload/internal/constant"
	"github.com/aq2208/goload/internal/dataaccess/file"
	"github.com/aq2208/goload/internal/dataaccess/mq/producer"
	"github.com/aq2208/goload/internal/model"
	"github.com/aq2208/goload/internal/repository"
	"github.com/aq2208/goload/utils"
)

type DownloadTask struct {
	Id             uint64
	OfAccountId    uint64
	DownloadType   constant.DownloadType
	Url            string
	DownloadStatus constant.DownloadStatus
}

type CreateDownloadTaskRequest struct {
	Token        string                `json:"token"`
	DownloadType constant.DownloadType `json:"download_type"`
	URL          string                `json:"url"`
}

type CreateDownloadTaskResponse struct {
	Data *model.DownloadTask `json:"data"`
}

type GetDownloadTaskListRequest struct {
	Token  string
	Offset uint64
	Limit  uint64
}

type GetDownloadTaskListResponse struct {
	Data       []model.DownloadTask `json:"data"`
	TotalItems uint64               `json:"total_items"`
}

type UpdateDownloadTaskRequest struct {
	Token          string
	DownloadTaskID uint64
	URL            string
}

type UpdateDownloadTaskResponse struct {
	DownloadTask *model.DownloadTask
}

type DeleteDownloadTaskRequest struct {
	Token          string
	DownloadTaskID uint64
}

type GetDownloadTaskFileRequest struct {
	Token          string
	DownloadTaskID uint64
}

type DownloadTaskService interface {
	CreateDownloadTask(ctx context.Context, req *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(ctx context.Context, req *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error)
	ProcessDownload(ctx context.Context, id uint64) error
	// GetDownloadTask(ctx context.Context, req *GetDownloadTaskRequest) (*GetDownloadTaskResponse, error)
	// UpdateDownloadTask(ctx context.Context, req *handler.UpdateDownloadTaskRequest) (*handler.UpdateDownloadTaskResponse, error)
	// DeleteDownloadTask(ctx context.Context, req *handler.DeleteDownloadTaskRequest) (*handler.DeleteDownloadTaskResponse, error)
	// GetDownloadFile(*GetDownloadFileRequest, grpc.ServerStreamingServer[GetDownloadFileResponse]) error
}

type downloadTaskService struct {
	db         *sql.DB
	repo       repository.DownloadTaskRepository
	tokenUtil  utils.Token
	producer   producer.Producer
	fileClient file.Client
}

// CreateDownloadTask implements DownloadTaskService.
func (d *downloadTaskService) CreateDownloadTask(ctx context.Context, req *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	// validate jwt
	accountId, err := d.tokenUtil.GetAccountIdAndExpireTime(ctx, req.Token)
	if err != nil {
		return &CreateDownloadTaskResponse{}, err
	}

	newDownloadTask := model.DownloadTask{
		UserID:       accountId,
		Status:       model.DownloadStatusQueued,
		DownloadType: model.DownloadTypeHTTP,
		URL:          req.URL,
	}

	// begin transaction
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return &CreateDownloadTaskResponse{}, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// insert new download_task record into DB
	newTaskId, err := d.repo.CreateDownloadTask(ctx, tx, newDownloadTask)
	if err != nil {
		return &CreateDownloadTaskResponse{}, err
	}
	newDownloadTask.ID = newTaskId

	// push new event to MQ
	msg, err := json.Marshal(newTaskId)
	if err != nil {
		return &CreateDownloadTaskResponse{}, fmt.Errorf("marshal error: %w", err)
	}

	if err := d.producer.SendMessage(msg); err != nil {
		return &CreateDownloadTaskResponse{}, fmt.Errorf("push message kafka error: %w", err)
	}

	tx.Commit()

	return &CreateDownloadTaskResponse{
		Data: &newDownloadTask,
	}, nil
}

// GetDownloadTaskList implements DownloadTaskService.
func (d *downloadTaskService) GetDownloadTaskList(ctx context.Context, req *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error) {
	panic("unimplemented")
}

func (d *downloadTaskService) ProcessDownload(ctx context.Context, id uint64) error {
	// TODO: handle in transaction, lock the data when update status to in-progress
	downloadTask, err := d.repo.GetDownloadTaskById(ctx, id)
	if err != nil {

	}

	if downloadTask.Status != model.DownloadStatusQueued {
		return nil
	}

	if err := d.repo.UpdateStatusDownloadTask(ctx, downloadTask.ID, string(model.DownloadStatusInProgress)); err != nil {
		return err
	}

	// execute download
	var downloader utils.Downloader
	switch downloadTask.DownloadType {
	case model.DownloadTypeHTTP:
		downloader = utils.NewHttpDownloader(downloadTask.URL)

	default:
		log.Default().Printf("Unsupported download type %s", downloadTask.DownloadType)
		return nil
	}

	fileClosure, err := d.fileClient.Write(ctx, "download_file.txt")
	if err != nil {
		return err
	}

	defer fileClosure.Close()

	err = downloader.Download(ctx, fileClosure)
	if err != nil {
		log.Default().Printf("Failed to download file: %v", err)
		return err
	}

	log.Default().Println("Download task executed successfully")

	// update status to completed
	if err := d.repo.UpdateStatusDownloadTask(ctx, downloadTask.ID, string(model.DownloadStatusCompleted)); err != nil {
		log.Default().Printf("Failed to update download task to success: %v", err)
		return err
	}

	return nil
}

// // UpdateDownloadTask implements DownloadTaskService.
// func (d *downloadTaskService) UpdateDownloadTask(ctx context.Context, req *handler.UpdateDownloadTaskRequest) (*handler.UpdateDownloadTaskResponse, error) {
// 	panic("unimplemented")
// }

// // DeleteDownloadTask implements DownloadTaskService.
// func (d *downloadTaskService) DeleteDownloadTask(ctx context.Context, req *handler.DeleteDownloadTaskRequest) (*handler.DeleteDownloadTaskResponse, error) {
// 	panic("unimplemented")
// }

func NewDownloadTaskService(
	repo repository.DownloadTaskRepository,
	tokenUtil utils.Token,
	producer producer.Producer,
	db *sql.DB,
	fileClient file.Client,
) DownloadTaskService {
	return &downloadTaskService{
		repo:      repo,
		tokenUtil: tokenUtil,
		producer:  producer,
		db:        db,
		fileClient: fileClient,
	}
}
