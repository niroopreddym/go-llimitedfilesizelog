package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/niroopreddym/go-llimitedfilesizelog/enums"
)

//LoggerService implementsthe Ilogger.go service
type LoggerService struct {
	LogWriter *LogFileMeta
}

type messageStructure struct {
	LogLevel string
	Time     time.Time
	Message  string
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
		LogWriter: fileMeta,
	}
}

//SetLogLevel sets the custom level apart from default
func (lwService *LoggerService) SetLogLevel(logLevel enums.VerbosityLevel) {
	lwService.LogWriter.VerbosityLevel = int(logLevel)
}

//Log method logs the data onto the local file
func (lwService *LoggerService) Log(logLevel enums.VerbosityLevel, message string) {
	if logLevel >= enums.VerbosityLevel(lwService.LogWriter.VerbosityLevel) {

		message := messageStructure{
			LogLevel: logLevel.String(),
			Time:     time.Now(),
			Message:  message,
		}

		byteData, err := json.Marshal(message)
		if err != nil {
			fmt.Println("error occured while marshalling the message data")
		}

		lwService.LogWriter.Write(string(byteData))
		if logLevel == enums.Fatal {
			os.Exit(1)
		}
	}
}
