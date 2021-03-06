// Package errors is a custom error message implementation with file path and file number
package errors

import (
	"fmt"
	"go/build"
	"os"
	"runtime"
	"strings"
)

type errors struct {
	msg      string
	fileInfo string
}

// New will return error with the caller file info from string type
func New(m string) error {
	return errors{
		msg:      m,
		fileInfo: getFileInfo(),
	}
}

// Error will return error string
func (e errors) Error() string {
	return e.fileInfo + " " + e.msg
}

// Wrap will return error wrapped with the file info when the error happened
func Wrap(e error) error {
	if e != nil {
		return &errors{
			msg:      e.Error(),
			fileInfo: getFileInfo(),
		}
	}
	return e
}

// to remove the application folder from file path
func getFileInfo() string {
	_, fn, line, _ := runtime.Caller(2)

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	fn = strings.TrimLeft(fn, gopath+"/src/")
	f := strings.SplitAfterN(fn, "/", 4)
	fn = f[len(f)-1]
	return fmt.Sprintf("[%s:%d]", fn, line)
}
