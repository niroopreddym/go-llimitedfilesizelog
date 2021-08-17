package testutil

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TempDir represents a temporary directory which can easily be removed and is
// a bit more convenient that using ioutil.Tempdir when you are writing tests.
//
// Typical usage:
//
//     func TestSomething(t *testing.T) {
//         tmpDir := testutil.CreateTempDir(t)
//         defer tmpDir.Remove(t)
//         ...
//         // use tmpDir.String() to get the name
//     }
//
type TempDir string

//TempFile is the log file
type TempFile struct {
	File *os.File
}

// CreateTempDir creates a temporary directory.
func CreateTempDir(t *testing.T) TempDir {
	tmp, err := ioutil.TempDir("", "hb-test-")
	assert.Nil(t, err)
	return TempDir(tmp)
}

func (tmp TempDir) String() string {
	return string(tmp)
}

//CreateTempLogFile creates a temp log File
func CreateTempLogFile(t *testing.T) TempFile {
	dirName := CreateTempDir(t)
	tmp, err := ioutil.TempFile(dirName.String(), "temp.log")
	assert.Nil(t, err)
	return TempFile{
		File: tmp,
	}
}

// Remove the temporary directory.
func (tmp TempDir) Remove(t *testing.T) {
	assert.Nil(t, os.RemoveAll(string(tmp)))
}

// Remove the temporary File.
func (tmp TempFile) Remove(t *testing.T) {
	err := tmp.File.Close()
	assert.Nil(t, err)
	err = os.Remove(tmp.File.Name())
	assert.Nil(t, err)
}
