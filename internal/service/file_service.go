package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/sonyamoonglade/lambda-file-service/internal/lambda_errors"
	"github.com/sonyamoonglade/lambda-file-service/internal/transport/dto"
	"github.com/sonyamoonglade/lambda-file-service/internal/transport/out"
	"github.com/sonyamoonglade/lambda-file-service/internal/types"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
)

type FileService interface {
	Put(ctx context.Context, dto dto.PutFileDto) (out.PutFileOut, error)
	FindOldestByRoot(items []*s3yandex.File, root types.Root) (*s3yandex.File, error)
	GetAll(ctx context.Context) (*s3yandex.Storage, error)
	Delete(ctx context.Context, dto dto.DeleteFileDto) error
	GenerateName(original string) string
}

type fileService struct {
	logger *log.Logger
	client *s3yandex.YandexS3Client
}

func NewFileService(logger *log.Logger, client *s3yandex.YandexS3Client) FileService {
	return &fileService{
		logger: logger,
		client: client,
	}
}

func (f *fileService) FindOldestByRoot(items []*s3yandex.File, root types.Root) (*s3yandex.File, error) {

	//Leave only root-included
	var sorted []*s3yandex.File

	//Sort it
	for _, item := range items {
		spl := strings.Split(item.Name, "_")
		initialRoot := spl[0]
		if types.Root(initialRoot) == root {
			sorted = append(sorted, item)
		}
	}

	//If only one or zero siblings were found
	if len(sorted) == 1 || len(sorted) == 0 {
		return nil, lambda_errors.UnableToDeleteFile
	}

	min := sorted[0]

	for _, item := range sorted[1:] {
		mod := item.LastModified.UnixNano()
		minx := min.LastModified.UnixNano()
		if minx > mod {
			min = item
		}
	}

	return min, nil
}

func (f *fileService) Put(ctx context.Context, dto dto.PutFileDto) (out.PutFileOut, error) {

	nameWithSeed := f.GenerateName(dto.Filename)

	err := f.client.PutFileWithBytes(ctx, &s3yandex.PutFileWithBytesInput{
		ContentType: dto.ContentType,
		FileName:    nameWithSeed,
		Destination: dto.Destination,
		FileBytes:   &dto.Bytes,
	})
	if err != nil {
		return out.PutFileOut{}, err
	}

	return out.PutFileOut{
		Filename: nameWithSeed,
	}, nil
}

func (f *fileService) Delete(ctx context.Context, dto dto.DeleteFileDto) error {
	return f.client.DeleteFile(ctx, &s3yandex.DeleteFileInput{
		FileName:    dto.Filename,
		Destination: dto.Destination,
	})
}

func (f *fileService) GetAll(ctx context.Context) (*s3yandex.Storage, error) {
	return f.client.GetFiles(ctx)
}

func (f *fileService) GenerateName(original string) string {
	spl := strings.Split(original, ".")
	l := len(spl)
	ext := spl[l-1]      //get last el after '.'
	name := spl[0 : l-1] // the rest from 0 to ext
	joinedName := strings.Join(name, ".")
	seed := time.Now().Unix()
	return fmt.Sprintf("%s_%d.%s", joinedName, seed, ext)
}
