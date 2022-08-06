package main

import (
	"context"
	"github.com/sonyamoonglade/lambda-file-service/pkg/env"
	"github.com/sonyamoonglade/lambda-file-service/pkg/file"
	"github.com/sonyamoonglade/lambda-file-service/pkg/types"
	"github.com/sonyamoonglade/s3-yandex-go/s3yandex"
	"log"
	"os"
)

func LambdaInput(ctx context.Context, input []byte) (*types.Response, error) {

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

	service := file.NewFileService(logger, client)
	transport := file.NewTransport(logger, service)
	log.Println("deps are ready")

	return transport.Router(ctx, input)
}
