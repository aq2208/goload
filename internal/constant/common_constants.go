package constant

type DownloadType int32

const (
	DownloadType_DOWNLOAD_TYPE_UNSPECIFIED DownloadType = 0
	DownloadType_DOWNLOAD_TYPE_HTTP        DownloadType = 1
)

type DownloadStatus int32

const (
	DownloadStatus_DOWNLOAD_STATUS_UNSPECIFIED DownloadStatus = 0
	DownloadStatus_DOWNLOAD_STATUS_PENDING     DownloadStatus = 1
	DownloadStatus_DOWNLOAD_STATUS_DOWNLOADING DownloadStatus = 2
	DownloadStatus_DOWNLOAD_STATUS_FAILED      DownloadStatus = 3
	DownloadStatus_DOWNLOAD_STATUS_SUCCESS     DownloadStatus = 4
)