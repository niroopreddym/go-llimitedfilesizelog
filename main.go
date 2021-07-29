package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/niroopreddym/go-llimitedfilesizelog/enums"
	"github.com/niroopreddym/go-llimitedfilesizelog/services"
)

func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func main() {
	logLocationBaseDir := "C:/Personal/logs"
	logger := services.NewLoggerService(aws.String("user_from_auth_&_auth"), aws.String("HB"), &logLocationBaseDir)
	logger.SetLogLevel(enums.Error)
	logger.Log(enums.Info, "Hi this is Info")
	logger.Log(enums.Warning, "Hi this is Warning")
	logger.Log(enums.Error, "Hi this is Error")
	logger.Log(enums.Fatal, "Hi this is Fatal")

	answer, err := divide(5, 0)
	if err != nil {
		logger.Log(enums.Error, "error occured while diving 5 with 0 "+err.Error())
	}

	// No errors!
	fmt.Println(answer)
}
