package biz

import (
	"context"
	"github.com/google/wire"
	"io"
	"strings"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewLinkUseCase,
	NewRoomWebhookUseCase,
	NewAuthUseCase,
	NewVoiceUseCase,
	NewRoomUseCase,
	NewRoomMsgUseCase,
	NewRoomVodUseCase,
)

type Transaction interface {
	WithTx(context.Context, func(ctx context.Context) error) error
}

type HttpRepo interface {
	Request(ctx context.Context, url string, req interface{}, response interface{}, isRetry bool) error
}

type File struct {
	Name   string
	Path   string
	URL    string
	Stream io.Reader
}

func (f *File) IsPDF() bool {
	return strings.HasSuffix(f.Name, ".pdf")
}

func (f *File) IsMp3() bool {
	return strings.HasSuffix(f.Name, ".mp3")
}

type StorageRepo interface {
	UploadFile(ctx context.Context, file *File) error
}
