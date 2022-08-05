package env

import (
	"errors"
	"os"
)

const (
	Owner              = "OWNER"
	AwsAccessKeyId     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

var InvalidEnvironmentVariables = errors.New("some environment variables are missing")

func LoadEnv() error {
	_, own := os.LookupEnv(Owner)
	_, secret := os.LookupEnv(AwsSecretAccessKey)
	_, access := os.LookupEnv(AwsAccessKeyId)
	if !own || !secret || !access {
		return InvalidEnvironmentVariables
	}
	return nil
}
