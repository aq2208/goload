package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	constant "github.com/aq2208/goload/internal/constant"
	"github.com/aq2208/goload/internal/service"
)

type DownloadTaskHandler struct {
	service service.DownloadTaskService
}

func NewDownloadTaskHandler(service service.DownloadTaskService) *DownloadTaskHandler {
	return &DownloadTaskHandler{service: service}
}

func (h *DownloadTaskHandler) CreateDownloadTaskHandler(w http.ResponseWriter, r *http.Request) {
	// parse request
	var req service.CreateDownloadTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	log.Default().Printf("CreateDownloadTaskRequest: %v", req)

	// validate request
	if (req.Token == "" || req.URL == "" || req.DownloadType != constant.DownloadType_DOWNLOAD_TYPE_HTTP) {
		http.Error(w, "Missing required field(s)", http.StatusBadRequest)
		return
	}

	// service
	resp, err := h.service.CreateDownloadTask(context.TODO(), &req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *DownloadTaskHandler) GetDownloadFile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("Download task ID: %s", id)

	taskID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        http.Error(w, "invalid task ID", http.StatusBadRequest)
        return
    }

	resp, err := h.service.GetDownloadFilePresignedUrl(context.TODO(), r.Header.Get("Authorization"), taskID)
	if err != nil {
		log.Printf("Error downloading file from S3: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(resp))
}

func (h *DownloadTaskHandler) GetListDownloadTasks(w http.ResponseWriter, r *http.Request) {
	
}
