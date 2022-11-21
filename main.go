package main

import (
	"context"
	"log"
	"os"

	"github.com/sonyamoonglade/lambda-file-service/internal/env"
	"github.com/sonyamoonglade/lambda-file-service/internal/service"
	"github.com/sonyamoonglade/lambda-file-service/internal/transport"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
)

func LambdaInput(ctx context.Context, input []byte) (*transport.Response, error) {

	log.Println("starting the execution")

	logger := log.New(os.Stdout, "[YANDEX CLOUD FUNCTION]", 0)

	envCfg, err := env.LoadEnv()
	if err != nil {
		log.Fatalf("could not load env variables: %s", err.Error())
	}

	//Must load env before creating this (loads env variables see *env.example)
	provider := s3yandex.NewEnvCredentialsProvider()

	client := s3yandex.NewYandexS3Client(ctx, provider, &s3yandex.YandexS3Config{
		Owner:  envCfg.Owner,
		Bucket: envCfg.Bucket,
		Debug:  true,
	})
	log.Println("client is ready")

	fileService := service.NewFileService(logger, client)
	transport := transport.NewTransport(logger, fileService)
	log.Println("deps are ready")

	return transport.Router(ctx, input)
}
