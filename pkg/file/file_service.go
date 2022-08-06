package file

import (
	"context"
	"fmt"
	"github.com/sonyamoonglade/lambda-file-service/pkg/file/dto"
	"github.com/sonyamoonglade/lambda-file-service/pkg/file/out"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
	"log"
	"strings"
	"time"
)

type Service interface {
	Put(ctx context.Context, dto dto.PutFileDto) (out.PutFileOut, error)
	PseudoDelete(ctx context.Context, dto dto.PseudoDeleteFileDto) error
	GetAll(ctx context.Context) (*s3yandex.Storage, error)
	GenerateName(original string) string
}

type fileService struct {
	logger *log.Logger
	client *s3yandex.YandexS3Client
}

func NewFileService(logger *log.Logger, client *s3yandex.YandexS3Client) Service {
	return &fileService{
		logger: logger,
		client: client,
	}
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

func (f *fileService) PseudoDelete(ctx context.Context, dto dto.PseudoDeleteFileDto) error {
	//TODO implement me
	panic("implement me")
}

func (f *fileService) GetAll(ctx context.Context) (*s3yandex.Storage, error) {
	//TODO implement me
	panic("implement me")
}

func (f *fileService) GenerateName(original string) string {
	spl := strings.Split(original, ".")
	l := len(spl)
	ext := spl[:l-1]
	name := spl[0]
	seed := time.Now().Unix()
	return fmt.Sprintf("%s_%d.%s", name, seed, ext)
}
