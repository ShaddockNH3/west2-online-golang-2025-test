package utils

import (
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ShaddockNH3/west2-online-golang-2025-test/task4/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func CreateDirIfNotExist(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

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

	savePath := filepath.Join(savePathDir, currentUserID+filename)

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

func UploadVideo(c *app.RequestContext, currentUserID string) (string, *multipart.FileHeader, error) {
	fileHeader, err := c.FormFile("data")
	if err != nil {
		return "", nil, err
	}
	// 检查文件是否存在
	if fileHeader == nil {
		return "", nil, errno.FileUploadErr
	}

	// 检查文件是否为视频
	file, err := fileHeader.Open()
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", nil, err
	}

	contentType := http.DetectContentType(buffer)
	allowedTypes := []string{"video/mp4", "video/quicktime", "video/webm", "video/x-msvideo"}
	isAllowed := false
	for _, t := range allowedTypes {
		if contentType == t {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return "", nil, errno.FileTypeErr
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", nil, errno.FileSeekErr
	}

	// 保存文件到本地
	// 需要确保有 ./data/videos 目录
	_, currentFilePath, _, _ := runtime.Caller(0)
	projectRoot := ""
	if idx := strings.LastIndex(currentFilePath, "task4"); idx != -1 {
		projectRoot = currentFilePath[:idx+len("task4")]
	}
	if projectRoot == "" {
		return "", nil, errno.UnableFindPathErr
	}

	filename := fileHeader.Filename
	savePathDir := filepath.Join(projectRoot, "pkg", "data", "videos", currentUserID)

	if err := os.MkdirAll(savePathDir, 0755); err != nil {
		return "", nil, err
	}

	savePath := filepath.Join(savePathDir, currentUserID+"_"+filename)

	if err = c.SaveUploadedFile(fileHeader, savePath); err != nil {
		return "", nil, err
	}

	return savePath, fileHeader, nil
}

// 处理封面文件，默认为视频第一帧
func GenerateVideoCover(videoSavePath, currentUserID, videoFilename string) (string, string, int64, error) {
	// 确定封面文件的保存路径 ---
	_, currentFilePath, _, _ := runtime.Caller(0)
	projectRoot := ""
	if idx := strings.LastIndex(currentFilePath, "task4"); idx != -1 {
		projectRoot = currentFilePath[:idx+len("task4")]
	}
	if projectRoot == "" {
		return "", "", 0, errno.UnableFindPathErr
	}

	coverFilename := strings.TrimSuffix(videoFilename, filepath.Ext(videoFilename)) + ".jpg"
	coverSaveDir := filepath.Join(projectRoot, "pkg", "data", "covers", currentUserID)
	if err := os.MkdirAll(coverSaveDir, 0755); err != nil {
		return "", "", 0, err
	}
	coverSavePath := filepath.Join(coverSaveDir, coverFilename)

	cmd := exec.Command("ffmpeg",
		"-i", videoSavePath, // 正确的视频输入路径
		"-ss", "00:00:01", // 从第1秒开始
		"-vframes", "1", // 只截取1帧
		"-y",          // 如果文件已存在，则覆盖
		coverSavePath, // 正确的封面输出路径
	)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", 0, err
	}

	fileInfo, err := os.Stat(coverSavePath)
	if err != nil {
		return "", "", 0, err
	}

	coverSize := fileInfo.Size()

	return coverFilename, coverSavePath, coverSize, nil
}
