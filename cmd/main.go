package main

import (
	"github.com/sonyamoonglade/lambda-file-service/pkg/env"
	"log"
)

func main() {

	log.Println("starting the execution")

	if err := env.LoadEnv(); err != nil {
		log.Fatalf("could not load env. %s", err.Error())
	}
	
}
