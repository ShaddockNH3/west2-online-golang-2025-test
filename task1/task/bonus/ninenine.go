package bonus

import (
	"fmt"
	"os"
)

// GenerateMultiplicationTable 生成九九乘法表并保存到文件
func GenerateMultiplicationTable(filename string) error {
	var content string

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			content += fmt.Sprintf("%d*%d=%-2d ", j, i, i*j)
		}
		content += "\n"
	}

	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	return nil
}

// GetMultiplicationTable 生成九九乘法表字符串（不保存文件）
func GetMultiplicationTable() string {
	var content string

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			content += fmt.Sprintf("%d*%d=%-2d ", j, i, i*j)
		}
		content += "\n"
	}

	return content
}
