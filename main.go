package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Data struct {
	Message string `json:"message,string"`
}

func Handler() (Data, error) {
	message := Data{Message: "hello"}
	return message, nil
}

func main() {
	lambda.Start(Handler)
}
