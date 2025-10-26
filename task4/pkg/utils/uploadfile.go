package utils

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func UploadImage(c *app.RequestContext, currentUserID string) (*multipart.FileHeader, string, error) {
	fileHeader, err := CheckImageFile(c)
	if err != nil {
		return nil, "", err
	}

	// 保存文件到本地
	// 需要确保有 ./data/avatars 目录
	_, currentFilePath, _, _ := runtime.Caller(0)
	projectRoot := ""
	if idx := strings.LastIndex(currentFilePath, "task4"); idx != -1 {
		projectRoot = currentFilePath[:idx+len("task4")]
	}
	if projectRoot == "" {
		return nil, "", errno.UnableFindPathErr
	}

	filename := fileHeader.Filename
	savePathDir := filepath.Join(projectRoot, "pkg", "data", "avatars")

	if err := os.MkdirAll(savePathDir, 0755); err != nil {
		return nil, "", err
	}

	savePath := filepath.Join(savePathDir, filename)

	if err = c.SaveUploadedFile(fileHeader, savePath); err != nil {
		return nil, "", err
	}

	return fileHeader, savePath, nil
}

// 检查是否是图片
func CheckImageFile(c *app.RequestContext) (*multipart.FileHeader, error) {
	fileHeader, err := c.FormFile("data")
	if err != nil {
		return nil, err
	}

	// 检查文件是否存在
	if fileHeader == nil {
		return nil, errno.FileUploadErr
	}

	// 检查文件是否为图片
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	contentType := http.DetectContentType(buffer)
	allowedTypes := []string{"image/jpeg", "image/png", "image/gif"}
	isAllowed := false
	for _, t := range allowedTypes {
		if contentType == t {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return nil, errno.FileTypeErr
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return fileHeader, nil
}
