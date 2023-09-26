package data

import (
	"context"
	"fmt"
	"path"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/go-kratos/kratos/v2/log"

	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
)

type storageRepo struct {
	conf *conf.Storage
	data *Data
	log  *log.Helper
}

func NewStorageRepo(conf *conf.Storage, data *Data, logger log.Logger) biz.StorageRepo {
	return &storageRepo{
		conf: conf,
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *storageRepo) UploadFile(ctx context.Context, file *biz.File) error {
	// create a client for the specified storage account
	fmt.Println(r.conf.AzureBlob.ConnectionString)
	client, err := azblob.NewClientFromConnectionString(r.conf.AzureBlob.ConnectionString, nil)
	if err != nil {
		r.log.Errorf("azblob.NewClientFromConnectionString, err:%v", err)
		return err
	}

	blobName := path.Join(file.Path, file.Name)

	var header *blob.HTTPHeaders
	if file.IsPDF() {
		applicationType := "application/pdf"
		header = &blob.HTTPHeaders{
			BlobContentType: &applicationType,
		}
	}

	// upload the file to the specified container with the specified blob name
	resp, err := client.UploadStream(ctx, r.conf.AzureBlob.ContainerName, blobName, file.Stream, &blockblob.UploadStreamOptions{
		HTTPHeaders: header,
	})
	if err != nil {
		r.log.Errorf("client.UploadStream, err:%v", err)
		return err
	}

	file.URL = r.conf.AzureBlob.CdnHost + path.Join(r.conf.AzureBlob.ContainerName, blobName)

	r.log.WithContext(ctx).Debugf("UploadStream file: %s response: %v", file.Name, resp)
	return nil
}
