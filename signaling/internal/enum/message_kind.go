package enum

type MessageKind uint64

const (
	UploadInitializationMessageKind   MessageKind = 1
	BlockTransmissionMessageKind      MessageKind = 2
	DownloadInitializationMessageKind MessageKind = 3
)
