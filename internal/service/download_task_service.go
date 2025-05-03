package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aq2208/goload/internal/constant"
	"github.com/aq2208/goload/internal/dataaccess/mq"
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
	// GetDownloadTask(ctx context.Context, req *GetDownloadTaskRequest) (*GetDownloadTaskResponse, error)
	// UpdateDownloadTask(ctx context.Context, req *handler.UpdateDownloadTaskRequest) (*handler.UpdateDownloadTaskResponse, error)
	// DeleteDownloadTask(ctx context.Context, req *handler.DeleteDownloadTaskRequest) (*handler.DeleteDownloadTaskResponse, error)
	// GetDownloadFile(*GetDownloadFileRequest, grpc.ServerStreamingServer[GetDownloadFileResponse]) error
}

type downloadTaskService struct {
	db        *sql.DB
	repo      repository.DownloadTaskRepository
	tokenUtil utils.Token
	producer  mq.Producer
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

	// TODO: push new event to MQ
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
	producer mq.Producer,
	db *sql.DB,
) DownloadTaskService {
	return &downloadTaskService{
		repo:      repo,
		tokenUtil: tokenUtil,
		producer:  producer,
		db: db,
	}
}
