package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/golang/glog"
	"github.com/niroopreddym/go-llimitedfilesizelog/internal/enums"
	"github.com/niroopreddym/go-llimitedfilesizelog/internal/services"
)

func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func main() {
	answer, err := divide(5, 0)

	logLocationBaseDir := "C:/Personal/logs"
	logger := services.NewLoggerService(aws.String("user_from_auth_&_auth"), aws.String("HB"), &logLocationBaseDir, enums.Error)
	logger.Log("Hi this is a test")

	if err != nil {
		// Handle the error based on log levels
		if glog.V(1) {
			glog.Warning(err)
			logger.LogWriter.Write("glog v1 level " + err.Error())
		}

		if glog.V(2) {
			glog.Error(err)
			logger.LogWriter.Write("glog v2 level " + err.Error())
		}

		if glog.V(3) {
			glog.Fatal(err)
			logger.LogWriter.Write("glog v3 level " + err.Error())
		}

		if glog.V(0) {
			glog.Info(err)
			logger.LogWriter.Write("glog v0 level " + err.Error())
		}

		fmt.Println(err)
	}

	// No errors!
	fmt.Println(answer)

}

//Log logs as part of func main
func Log(message string) {
	logLocationBaseDir := "C:/Personal/logs"
	logger := services.NewLoggerService(aws.String("user_from_auth_&_auth"), aws.String("HB"), &logLocationBaseDir, enums.Error)
	fmt.Println("________")
	logger.Log("Hi this is a test from func main")
	fmt.Println("________")
}
