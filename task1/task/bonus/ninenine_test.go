package bonus

import (
	"os"
	"testing"
)

func TestGenerateMultiplicationTable(t *testing.T) {
	filename := "multiplication_table.txt"

	// 生成九九乘法表并保存到文件
	err := GenerateMultiplicationTable(filename)
	if err != nil {
		t.Fatalf("生成乘法表失败: %v", err)
	}

	t.Logf("九九乘法表已保存到: %s", filename)

	// 读取文件内容并输出到日志
	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("读取文件失败: %v", err)
	}

	t.Log("九九乘法表内容:")
	t.Log("\n" + string(content))
}

func TestGetMultiplicationTable(t *testing.T) {
	// 直接获取乘法表字符串
	table := GetMultiplicationTable()
	t.Log("九九乘法表:")
	t.Log("\n" + table)
}
