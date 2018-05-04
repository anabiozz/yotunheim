package utility

import (
	"fmt"
	"log"
	"runtime/debug"
)

var (
	// BugMsg ...
	BugMsg = "There was an unexpected issue; please report this as a bug."
)

// CustomError ...
type CustomError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

// WrapError ...
func WrapError(
	err error,
	messagef string,
	msgArgs ...interface{},
) CustomError {
	return CustomError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err CustomError) Error() string {
	return err.Message
}

// HandleError ...
func HandleError(key string, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	log.Printf("[%v] %v", key, message)
}
