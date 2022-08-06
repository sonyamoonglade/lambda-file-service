package main

import (
	"context"
	"github.com/sonyamoonglade/lambda-file-service/pkg/env"
	"github.com/sonyamoonglade/lambda-file-service/pkg/file"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
	"log"
	"os"
)

func main() {

	log.Println("starting the execution")

	logger := log.New(os.Stdout, "[YANDEX CLOUD FUNCTION]", 0)

	if err := env.LoadEnv(); err != nil {
		log.Fatalf("could not load env. %s", err.Error())
	}

	ctx := context.Background()

	//Must load env before creating this (loads env variables see *env.example)
	provider := s3yandex.NewEnvCredentialsProvider()

	client := s3yandex.NewYandexS3Client(ctx, provider, &s3yandex.YandexS3Config{
		Owner:  os.Getenv(env.Owner),
		Bucket: os.Getenv(env.Bucket),
		Debug:  false,
	})
	log.Println("client is ready")

	service := file.NewFileService(logger, client)
	transport := file.NewTransport(service)
	_ = transport
}
