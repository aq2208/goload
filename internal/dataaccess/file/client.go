package file

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/minio/minio-go"
)

type bufferedFileReader struct {
	file           *os.File
	bufferedReader io.Reader
}

func newBufferedFileReader(file *os.File) io.ReadCloser {
	return &bufferedFileReader{
		file:           file,
		bufferedReader: bufio.NewReader(file),
	}
}

func (b bufferedFileReader) Close() error {
	return b.file.Close()
}

func (b bufferedFileReader) Read(p []byte) (int, error) {
	return b.bufferedReader.Read(p)
}

type Client interface {
	Write(ctx context.Context, filePath string) (io.WriteCloser, error)
	Read(ctx context.Context, filePath string) (io.ReadCloser, error)
	GetPresignedUrl(ctx context.Context, id uint64) (*url.URL, error)
}

type LocalClient struct {
	Directory string
}

// GetPresignedUrl implements Client.
func (l *LocalClient) GetPresignedUrl(ctx context.Context, id uint64) (*url.URL, error) {
	panic("unimplemented")
}

func NewLocalClient(directory string) Client {
	return &LocalClient{Directory: directory}
}

func (l *LocalClient) Read(ctx context.Context, filePath string) (io.ReadCloser, error) {
	absolutePath := path.Join(l.Directory, filePath)
	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	}

	return newBufferedFileReader(file), nil
}

func (l *LocalClient) Write(
	ctx context.Context,
	filePath string,
) (io.WriteCloser, error) {
	absolutePath := path.Join(l.Directory, filePath)
	file, err := os.Create(absolutePath)
	if err != nil {
		log.Default().Printf("Error creating new file: %v", err)
		return nil, err
	}

	return file, nil
}

type S3Client struct {
	MinioClient *minio.Client
	BucketName  string
}

func (s *S3Client) Read(ctx context.Context, filePath string) (io.ReadCloser, error) {
	object, err := s.MinioClient.GetObjectWithContext(ctx, s.BucketName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Default().Printf("Error get s3 object: %v", err)
		return nil, err
	}

	return object, nil
}

func (s *S3Client) Write(ctx context.Context, filePath string) (io.WriteCloser, error) {
	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeReader.Close()
		_, err := s.MinioClient.PutObjectWithContext(
			ctx,
			s.BucketName,
			filePath,
			pipeReader,
			-1, // unknown size
			minio.PutObjectOptions{},
		)
		if err != nil {
			log.Printf("Failed to upload to S3: %v", err)
			_ = pipeReader.CloseWithError(err) // closes reader side with error
		} else {
			log.Println("Upload file to S3 successfully")
		}
	}()

	return pipeWriter, nil // return the writer end to the caller (e.g., Download())
}

func (s *S3Client) GetPresignedUrl(ctx context.Context, id uint64) (*url.URL, error) {
	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment")

	objectName := fmt.Sprint("download_file_" + fmt.Sprint(id))

	// Gernerate presigned get object url.
	presignedURL, err := s.MinioClient.PresignedGetObject(
		s.BucketName,
		objectName,
		time.Duration(300)*time.Second,
		reqParams,
	)
	if err != nil {
		log.Printf("Error get S3 presigned url: %v", err)
		return nil, err
	}

	log.Printf("Presigned URL: %s", presignedURL)

	return presignedURL, nil
}

func NewS3Client(
	Bucket string,
	Address string,
	Username string,
	Password string,
) Client {
	log.Printf("MinIO Client config address: %s", Address)
	minioClient, err := minio.New(Address, Username, Password, false)
	if err != nil {
		log.Default().Printf("Error create minio client: %v", err)
		return nil
	}

	log.Print("Start MinIO Client successfully")

	return &S3Client{MinioClient: minioClient, BucketName: Bucket}
}
