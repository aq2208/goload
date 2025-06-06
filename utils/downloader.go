package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	HttpResponseContentType = "Content-Type"
	HttpMetadataKeyContentType = "content-type"
)

type Downloader interface {
	Download(ctx context.Context, writer io.Writer) (map[string]any, error)
}

type HttpDownloader struct {
	URL string
}

func NewHttpDownloader(URL string) Downloader {
	return &HttpDownloader{URL: URL}
}

func (h *HttpDownloader) Download(ctx context.Context, writer io.Writer) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.URL, http.NoBody)
	if err != nil {
		log.Default().Printf("Error creating HTTP request: %v", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Default().Printf("Error making HTTP request: %v", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Default().Printf("Error response HTTP status code not 200")
		return nil, fmt.Errorf("BAD STATUS: %s", resp.Status)
	}

	defer resp.Body.Close()

	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		log.Default().Printf("Error reading response and write to writer: %v", err)
		return nil, err
	}

	metadata := map[string]any {
		HttpMetadataKeyContentType: resp.Header.Get(HttpResponseContentType),
	}

	return metadata, nil
}
