package handler

import (
	"context"
	v1 "faceto-ai/api_gen/voice/v1"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"mime/multipart"
	netHttp "net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	OperationVoiceUpload = "/faceto.v1.Voice/Upload"
)

func convertMultipartFile(fileFromMultipart multipart.File, header *multipart.FileHeader) (*biz.File, error) {
	// if header.Size > FileSizeLimit {
	// 	return nil, errors.New("file size invalid")
	// }

	split := strings.Split(header.Filename, ".")
	if len(split) < 2 {
		return nil, errors.New("file type invalid")
	}
	ext := split[len(split)-1]
	now := time.Now()
	fPath := fmt.Sprintf("upload/%s/%s/", now.Format("2006"), now.Format("01"))
	name := fmt.Sprintf("%s.%s", uuid.New().String(), ext)
	return &biz.File{
		Name:   name,
		Path:   fPath,
		Stream: fileFromMultipart,
	}, nil
}

func VoiceUploadHandler(srv *service.VoiceService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		req := ctx.Request()
		file, header, err := req.FormFile("file")
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := convertMultipartFile(file, header)
		if err != nil {
			return err
		}

		http.SetOperation(ctx, OperationVoiceUpload)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFileToStorage(ctx, req.(*biz.File))
		})

		out, err := h(ctx, fileInfo)
		if err != nil {
			return err
		}
		reply := out.(*v1.FileUploadReply)
		return ctx.Result(netHttp.StatusOK, reply)
	}
}
