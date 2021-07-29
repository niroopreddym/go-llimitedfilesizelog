package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//LogFileMeta provides required methods to create and maintain dynamic log files
type LogFileMeta struct {
	Lock           sync.Mutex
	VerbosityLevel int
	LogFile        string
	FileLimitInKB  int
	errLog         *log.Logger
}

// //TestGetFileSize is just a test method
// func (w *FileSizeLimitter) TestGetFileSize() error {
// 	file, err := os.Open(w.LogFile + "/" + fileName + ".log")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fi, err := file.Stat()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(fi.Size())
// 	return nil
// }

//CreateNewLogFile creates a file inside given log location
func CreateNewLogFile(dir, serviceName string, truncateDataOnLog bool) (*LogFileMeta, error) {
	instance := LogFileMeta{}
	parentPath := dir

	if !truncateDataOnLog {
		parentPath = filepath.Join(dir, serviceName)
	}

	filePath := filepath.Join(parentPath, serviceName+"_logdata.log")

	// check if file exists
	info, err := os.Stat(filePath)

	fmt.Println(info)
	// create file if not exists
	if os.IsNotExist(err) {
		err := os.MkdirAll(parentPath, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		defer file.Close()
	}

	instance.LogFile = filePath
	instance.Lock = sync.Mutex{}
	instance.FileLimitInKB = 10
	fmt.Println("File Created Successfully", filePath)

	return &instance, nil
}

//Write writes the given message param to the log file pointed by the logFileMeta Struct
func (w *LogFileMeta) Write(message string) {
	w.Lock.Lock()

	file, err := os.OpenFile(w.LogFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("Unable to open log file: %s\n%s\n", err, message)
		return
	}

	w.errLog = log.New(file, "", 0)

	defer func() {
		w.Lock.Unlock()
		file.Close()
	}()

	w.errLog.Println(message)

	fi, err := os.Stat(w.LogFile)
	if err != nil {
		fmt.Println("Unable to stat log file: " + err.Error())
		return
	}

	fileLimitInBytes := w.FileLimitInKB * 1024
	fileSize := int(fi.Size())
	if fileSize > fileLimitInBytes {
		err := w.writeOldDataToNewTSFile(fileLimitInBytes, file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

//writeOldDataToNewTSFile holds the logic to split the log to 10kb size
func (w *LogFileMeta) writeOldDataToNewTSFile(fileLimitInBytes int, oldFileHandler *os.File) error {
	b, err := ioutil.ReadFile(w.LogFile)
	if err != nil {
		fmt.Println("Unable to read log file" + err.Error())
		return err
	}

	directoryPath := filepath.Dir(w.LogFile)
	fileName := serviceLogNameWithoutExtSliceNotation(filepath.Base(w.LogFile))

	fileMeta, err := CreateNewLogFile(directoryPath, fileName+"_"+time.Now().Format("20060102150405"), true)
	newFile, err := os.OpenFile(fileMeta.LogFile, os.O_APPEND|os.O_WRONLY, 0600)

	oldData := b[0 : fileLimitInBytes-1]
	_, err = newFile.Write(oldData)
	if err != nil {
		fmt.Println("Unable to write olddata to new log file" + err.Error())
		return err
	}

	newData := b[fileLimitInBytes:]
	err = oldFileHandler.Close()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	oldDataFile, err := os.OpenFile(w.LogFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Unable to open temp log file" + err.Error())
		return err
	}

	err = oldDataFile.Truncate(0)
	if err != nil {
		fmt.Println("Unable to truncate temp log file" + err.Error())
		return err
	}

	_, err = oldDataFile.Write(newData)
	if err != nil {
		fmt.Println("Unable to write to temp log file" + err.Error())
		return err
	}

	return nil
}

func serviceLogNameWithoutExtSliceNotation(fileName string) string {
	completeName := fileName[:len(fileName)-len(filepath.Ext(fileName))]
	serviceName := strings.Split(completeName, "_")
	return serviceName[0]
}
