package services

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/niroopreddym/go-llimitedfilesizelog/internal/enums"
)

//LoggerService implementsthe Ilogger.go service
type LoggerService struct {
	LogWriter *LogFileMeta
}

//NewLoggerService is the constructor for LoggerService Struct
func NewLoggerService(userName *string, serviceName *string, logLocationBaseDir *string, verbosityLevel enums.VerbosityLevel) *LoggerService {
	if strings.Contains(*serviceName, "_") {
		fmt.Println("can't include _ as part of serviceName.. Please use - instead")
		return nil
	}
	flag.Set("v", strconv.Itoa(int(verbosityLevel)))
	flag.Set("logtostderr", "true")
	flag.Parse()

	fileMeta, err := CreateNewLogFile(*logLocationBaseDir, *serviceName, int(verbosityLevel), false)
	if err != nil {
		fmt.Println("Unable to instantiate FileMeta struct" + err.Error())
		return nil
	}

	return &LoggerService{
		LogWriter: fileMeta,
	}
}

//Log method logs the data onto the local file
func (service *LoggerService) Log(message string) {
	fmt.Println("________")
	service.LogWriter.Write(message)
	fmt.Println("________")
	service.LogWriter.Write("Love You Mom")
}
