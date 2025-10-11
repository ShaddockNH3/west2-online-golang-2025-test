package errno

import (
	"errors"
	"fmt"
)

// ========================== Error Codes ==========================

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10001 // 服务级错误
	ParamErrCode                  // 参数错误

	UserAlreadyExistErrCode // 用户已存在
)

// ========================== Error Messages ==========================

const (
	SuccessMsg   = "Success"
	ServiceErrMsg = "Service is unable to start successfully"
	ParamErrMsg  = "Wrong Parameter has been given"
)

// ErrNo defines a custom error type.
type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

// Error makes it compatible with the `error` interface.
func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

// NewErrNo creates a new ErrNo.
func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

// WithMessage allows chaining to modify the error message.
func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

// ========================== Predefined Errors ==========================

var (
	// Common errors
	Success    = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr = NewErrNo(ServiceErrCode, ServiceErrMsg)
	ParamErr   = NewErrNo(ParamErrCode, ParamErrMsg)

	// User-related errors
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	// ... 之后可以继续在这里添加其他错误哦 ♪
)

// ConvertErr converts a generic error to ErrNo.
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	// If the error is already an ErrNo, return it directly.
	if errors.As(err, &Err) {
		return Err
	}

	// Otherwise, wrap it as a generic service error.
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
