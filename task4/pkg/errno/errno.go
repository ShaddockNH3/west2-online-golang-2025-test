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

	UserAlreadyExistErrCode         // 用户已存在
	UserNotExistErrCode             // 用户不存在
	PasswordIsNotVerifiedErrCode    // 密码错误
	UnableToRetrieveUserInfoErrCode // 无法获取用户信息

	FileUploadErrCode // 文件上传错误
	FileSaveErrCode   // 文件保存错误
	UnableFindPathErrCode // 无法找到路径
	FileOpenErrCode   // 文件打开错误
	FileReadErrCode   // 文件读取错误
	FileTypeErrCode   // 文件类型错误
	FileSeekErrCode   // 文件大小错误
	FileDirCreateErrCode // 文件目录创建错误
	FileCoverCreateErrCode // 文件封面创建错误
	VideoAlreadyExistErrCode // 视频已存在
)

// ========================== Error Messages ==========================

const (
	SuccessMsg    = "Success"
	ServiceErrMsg = "Service is unable to start successfully"
	ParamErrMsg   = "Wrong Parameter has been given"
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
	UserAlreadyExistErr         = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UserNotExistErr             = NewErrNo(UserNotExistErrCode, "User does not exist")
	PasswordIsNotVerified       = NewErrNo(PasswordIsNotVerifiedErrCode, "Password is not verified")
	UnableToRetrieveUserInfoErr = NewErrNo(UnableToRetrieveUserInfoErrCode, "Unable to retrieve user information")

	// File-related errors
	FileUploadErr = NewErrNo(FileUploadErrCode, "File upload error")
	FileSaveErr   = NewErrNo(FileSaveErrCode, "File save error")
	UnableFindPathErr = NewErrNo(UnableFindPathErrCode, "Unable to find path")
	FileOpenErr   = NewErrNo(FileOpenErrCode, "File open error")
	FileReadErr   = NewErrNo(FileReadErrCode, "File read error")
	FileTypeErr   = NewErrNo(FileTypeErrCode, "File type error")
	FileSeekErr   = NewErrNo(FileSeekErrCode, "File size error")
	FileDirCreateErr = NewErrNo(FileDirCreateErrCode, "File directory creation error")
	FileCoverCreateErr = NewErrNo(FileCoverCreateErrCode, "File cover creation error")
	VideoAlreadyExistErr = NewErrNo(VideoAlreadyExistErrCode, "Video already exists")
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
