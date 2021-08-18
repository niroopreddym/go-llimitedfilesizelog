package services

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/niroopreddym/go-llimitedfilesizelog/enums"
)

//LoggerService implementsthe Ilogger.go service
type LoggerService struct {
	logWriter *LogFileMeta
	username  string
}

type messageStructure struct {
	LogLevel string
	Time     string
	Message  string
	UserName string
}

//NewLoggerService is the constructor for LoggerService Struct the default verbose level is 3
func NewLoggerService(userName *string, serviceName *string, logLocationBaseDir *string) *LoggerService {
	if strings.Contains(*serviceName, "_") {
		fmt.Println("can't include _ as part of serviceName.. Please use - instead")
		return nil
	}

	fileMeta, err := CreateNewLogFile(*logLocationBaseDir, *serviceName, false)
	if err != nil {
		fmt.Println("Unable to instantiate FileMeta struct" + err.Error())
		return nil
	}

	return &LoggerService{
		logWriter: fileMeta,
		username:  *userName,
	}
}

//SetLogLevel sets the custom level apart from default
func (lwService *LoggerService) SetLogLevel(logLevel enums.VerbosityLevel) {
	lwService.logWriter.VerbosityLevel = int(logLevel)
}

//Log method logs the data onto the local file
func (lwService *LoggerService) Log(logLevel enums.VerbosityLevel, message string) error {
	if logLevel >= enums.VerbosityLevel(lwService.logWriter.VerbosityLevel) {

		message := messageStructure{
			LogLevel: logLevel.String(),
			Time:     time.Now().Format("2006-01-02 15:04:05.000000"),
			Message:  message,
			UserName: lwService.username,
		}

		byteData, err := json.Marshal(message)
		if err != nil {
			fmt.Println("error occured while marshalling the message data")
			return err
		}

		lwService.logWriter.Write(string(byteData))
		if logLevel == enums.Fatal {
			return fmt.Errorf("Fatal error")
		}
	}

	return nil
}
