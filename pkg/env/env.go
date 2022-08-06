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

var InvalidEnvironmentVariables = errors.New("some environment variables are missing")

func LoadEnv() error {
	_, own := os.LookupEnv(Owner)
	_, secret := os.LookupEnv(AwsSecretAccessKey)
	_, access := os.LookupEnv(AwsAccessKeyId)
	_, bucket := os.LookupEnv(Bucket)
	if !own || !secret || !access || !bucket {
		return InvalidEnvironmentVariables
	}
	return nil
}
