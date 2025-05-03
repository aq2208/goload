package http

import (
	"encoding/json"
	"log"
	"net/http"

	constant "github.com/aq2208/goload/internal/constant"
	"github.com/aq2208/goload/internal/model"
	"github.com/aq2208/goload/internal/service"
)

type DownloadTask struct {
	Id             uint64
	OfAccountId    uint64
	DownloadType   constant.DownloadType
	Url            string
	DownloadStatus constant.DownloadStatus
}

type CreateDownloadTaskRequest struct {
	Token        string
	DownloadType constant.DownloadType
	URL          string
}

type CreateDownloadTaskResponse struct {
	Data model.DownloadTask `json:"data"`
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

type DownloadTaskHandler struct {
	service service.DownloadTaskService
}

func NewDownloadTaskHandler(service service.DownloadTaskService) *DownloadTaskHandler {
	return &DownloadTaskHandler{service: service}
}

func (h *DownloadTaskHandler) CreateDownloadTaskHandler(w http.ResponseWriter, r *http.Request) {
	// parse request
	var req CreateDownloadTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Default().Printf("CreateAccountRequest: %v", req)

	// validate

	// service
}
