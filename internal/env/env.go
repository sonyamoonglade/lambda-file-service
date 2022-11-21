package env

import (
	"errors"
	"os"
)

const (
	Owner              = "OWNER"
	AwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
	Bucket             = "BUCKET"
)

type Config struct {
	Owner  string
	Bucket string
}

var InvalidEnvironmentVariables = errors.New("some environment variables are missing")

func LoadEnv() (*Config, error) {
	own, ok := os.LookupEnv(Owner)
	if !ok {
		return nil, InvalidEnvironmentVariables
	}
	bucket, ok := os.LookupEnv(Bucket)
	if !ok {
		return nil, InvalidEnvironmentVariables
	}
	_, ok = os.LookupEnv(AwsSecretAccessKey)
	if !ok {
		return nil, InvalidEnvironmentVariables
	}
	_, ok = os.LookupEnv(AwsAccessKeyId)
	if !ok {
		return nil, InvalidEnvironmentVariables
	}

	return &Config{
		Owner:  own,
		Bucket: bucket,
	}, nil
}
