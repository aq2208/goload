package file

import (
	"bufio"
	"context"
	"io"
	"os"
	"path"
)

type Client interface {
	Write(ctx context.Context, filePath string) (io.WriteCloser, error)
	Read(ctx context.Context, filePath string) (io.ReadCloser, error)
}

type LocalFileClient struct {
	Directory string
}

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

func (l *LocalFileClient) Read(ctx context.Context, filePath string) (io.ReadCloser, error) {
	absolutePath := path.Join(l.Directory, filePath)
	file, err := os.Open(absolutePath)
	if err != nil {
		return nil, err
	} 

	return newBufferedFileReader(file), nil
}

func (l *LocalFileClient) Write(
	ctx context.Context, 
	filePath string,
) (io.WriteCloser, error) {
	absolutePath := path.Join(l.Directory, filePath)
	file, err := os.Create(absolutePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewLocalFileClient(directory string) Client {
	return &LocalFileClient{Directory: directory}
}
