package errno

import (
	"fmt"
	"reflect"
)

// Errno 返回错误码和消息的结构体
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	if reflect.ValueOf(err).IsValid() {
		if reflect.TypeOf(err).Kind() != reflect.Struct {
			if reflect.ValueOf(err).IsNil() {
				return OK.Code, OK.Message
			}
		}

	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
