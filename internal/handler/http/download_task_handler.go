package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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

	log.Default().Printf("CreateAccountRequest: %v", req)

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
