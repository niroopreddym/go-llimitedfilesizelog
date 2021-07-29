package services

import "github.com/niroopreddym/go-llimitedfilesizelog/enums"

//LoggerIface exposes the log package defintions to other modules
type LoggerIface interface {
	Log(logLevel enums.VerbosityLevel, message string)
	// Log2Db(message string)
	// Log2Kinesis(message string)
}
