package service

import (
	"context"

	goload "github.com/aq2208/goload/internal/generated"
	"github.com/aq2208/goload/internal/model"
	"github.com/aq2208/goload/internal/repository"
	"github.com/aq2208/goload/utils"
)

type CreateDownloadTaskRequest struct {
	Token        string
	DownloadType goload.DownloadType
	Url          string
}

type CreateDownloadTaskResponse struct {
	DownloadTask goload.DownloadTask
}

type GetDownloadTaskListRequest struct {
}

type GetDownloadTaskListResponse struct {
}

type UpdateDownloadTaskRequest struct {
}

type UpdateDownloadTaskResponse struct {
}

type DeleteDownloadTaskRequest struct {
}

type DeleteDownloadTaskResponse struct {
}

type DownloadTaskService interface {
	CreateDownloadTask(ctx context.Context, req *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(ctx context.Context, req *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error)
	// GetDownloadTask(ctx context.Context, req *GetDownloadTaskRequest) (*GetDownloadTaskResponse, error)
	UpdateDownloadTask(ctx context.Context, req *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error)
	DeleteDownloadTask(ctx context.Context, req *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error)
	// GetDownloadFile(*GetDownloadFileRequest, grpc.ServerStreamingServer[GetDownloadFileResponse]) error
}

type downloadTaskService struct {
	repo      repository.DownloadTaskRepository
	tokenUtil utils.Token
}

// CreateDownloadTask implements DownloadTaskService.
func (d *downloadTaskService) CreateDownloadTask(ctx context.Context, req *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	accountId, err := d.tokenUtil.GetAccountIdAndExpireTime(ctx, req.Token)
	if err != nil {
		return &CreateDownloadTaskResponse{}, err
	}

	newDownloadTask := model.DownloadTask {
		UserID: accountId,
	}

	newTaskId, err := d.repo.CreateDownloadTask(ctx, newDownloadTask)
	if err != nil {
		return &CreateDownloadTaskResponse{}, err
	}

	// TODO: push new event to MQ

	return &CreateDownloadTaskResponse{
		DownloadTask: goload.DownloadTask{
			Id:             newTaskId,
			OfUser:         nil,
			DownloadType:   req.DownloadType,
			Url:            req.Url,
			DownloadStatus: goload.DownloadStatus_queued,
		},
	}, nil
}

// DeleteDownloadTask implements DownloadTaskService.
func (d *downloadTaskService) DeleteDownloadTask(ctx context.Context, req *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error) {
	panic("unimplemented")
}

// GetDownloadTaskList implements DownloadTaskService.
func (d *downloadTaskService) GetDownloadTaskList(ctx context.Context, req *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error) {
	panic("unimplemented")
}

// UpdateDownloadTask implements DownloadTaskService.
func (d *downloadTaskService) UpdateDownloadTask(ctx context.Context, req *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error) {
	panic("unimplemented")
}

func NewDownloadTaskService(
	repo repository.DownloadTaskRepository,
	tokenUtil utils.Token,
) DownloadTaskService {
	return &downloadTaskService{
		repo:      repo,
		tokenUtil: tokenUtil,
	}
}
