package services

import (
	"sync"
	"testing"

	"github.com/niroopreddym/go-llimitedfilesizelog/enums"
	"github.com/niroopreddym/go-llimitedfilesizelog/testutil"
	"github.com/stretchr/testify/assert"
)

func TestLogData(t *testing.T) {
	tmpFile := testutil.CreateTempLogFile(t)
	defer func() {
		tmpFile.Remove(t)
	}()

	// controller := gomock.NewController(t)
	// defer controller.Finish()

	logService := LoggerService{
		logWriter: &LogFileMeta{
			VerbosityLevel: int(enums.Error),
			FileLimitInKB:  10,
			Lock:           sync.Mutex{},
			LogFile:        string(tmpFile.File.Name()),
		},
		username: "user name from auth & auth",
	}

	logLevel := enums.Error
	err := logService.Log(logLevel, "some error message")
	assert.Nil(t, err)
}
