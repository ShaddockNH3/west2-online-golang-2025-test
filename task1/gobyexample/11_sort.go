package gobyexample

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

// ==================================================
// 1. 基础排序演示
// ==================================================

// demonstrateBasicSort 演示基础排序功能
func demonstrateBasicSort() {
	fmt.Println("=== 基础排序演示 ===")

	// 字符串排序
	fruits := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("排序前: %v\n", fruits)
	slices.Sort(fruits)
	fmt.Printf("排序后: %v\n", fruits)

	// 整数排序
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("数字排序前: %v\n", numbers)
	slices.Sort(numbers)
	fmt.Printf("数字排序后: %v\n", numbers)

	// 检查是否已排序
	fmt.Printf("是否已排序: %t\n", slices.IsSorted(numbers))
	fmt.Println()
}

// ==================================================
// 2. 自定义排序规则
// ==================================================

// demonstrateCustomSort 演示自定义排序规则
func demonstrateCustomSort() {
	fmt.Println("=== 自定义排序演示 ===")

	// 按字符串长度排序
	words := []string{"Go", "Programming", "Language", "Sort", "Algorithm"}
	fmt.Printf("按长度排序前: %v\n", words)

	// 自定义比较函数：按长度排序
	lengthCompare := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(words, lengthCompare)
	fmt.Printf("按长度排序后: %v\n", words)

	// 按字母顺序逆序排序
	cities := []string{"北京", "上海", "广州", "深圳", "杭州"}
	fmt.Printf("逆序排序前: %v\n", cities)
	slices.SortFunc(cities, func(a, b string) int {
		return cmp.Compare(b, a) // 注意参数顺序相反
	})
	fmt.Printf("逆序排序后: %v\n", cities)
	fmt.Println()
}

// ==================================================
// 3. 结构体排序
// ==================================================

// Person 人员结构体
type Person struct {
	Name string
	Age  int
	City string
}

// String 实现Stringer接口，方便打印
func (p Person) String() string {
	return fmt.Sprintf("{%s, %d岁, %s}", p.Name, p.Age, p.City)
}

// demonstrateStructSort 演示结构体排序
func demonstrateStructSort() {
	fmt.Println("=== 结构体排序演示 ===")

	people := []Person{
		{"张三", 25, "北京"},
		{"李四", 30, "上海"},
		{"王五", 20, "广州"},
		{"赵六", 35, "深圳"},
		{"钱七", 28, "杭州"},
	}

	fmt.Println("原始顺序:")
	for i, p := range people {
		fmt.Printf("  %d. %s\n", i+1, p)
	}

	// 按年龄排序
	fmt.Println("\n按年龄排序:")
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
	for i, p := range people {
		fmt.Printf("  %d. %s\n", i+1, p)
	}

	// 按姓名排序
	fmt.Println("\n按姓名排序:")
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Name, b.Name)
	})
	for i, p := range people {
		fmt.Printf("  %d. %s\n", i+1, p)
	}

	// 多级排序：先按城市，再按年龄
	fmt.Println("\n多级排序(城市->年龄):")
	slices.SortFunc(people, func(a, b Person) int {
		if cityCompare := cmp.Compare(a.City, b.City); cityCompare != 0 {
			return cityCompare
		}
		return cmp.Compare(a.Age, b.Age)
	})
	for i, p := range people {
		fmt.Printf("  %d. %s\n", i+1, p)
	}
	fmt.Println()
}

// ==================================================
// 4. 字符串处理演示
// ==================================================

// demonstrateStringProcessing 演示字符串处理功能
func demonstrateStringProcessing() {
	fmt.Println("=== 字符串处理演示 ===")

	text := "Go Programming Language"

	// 基础字符串操作
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("长度: %d\n", len(text))
	fmt.Printf("大写: %s\n", strings.ToUpper(text))
	fmt.Printf("小写: %s\n", strings.ToLower(text))
	fmt.Printf("是否包含'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("替换'Go'->'Python': %s\n", strings.Replace(text, "Go", "Python", 1))

	// 字符串分割和连接
	words := strings.Fields(text)
	fmt.Printf("分割为单词: %v\n", words)
	fmt.Printf("用'-'连接: %s\n", strings.Join(words, "-"))

	// 字符串修剪
	messyText := "   Hello World   "
	fmt.Printf("修剪前: '%s'\n", messyText)
	fmt.Printf("修剪后: '%s'\n", strings.TrimSpace(messyText))
	fmt.Println()
}

// ==================================================
// 5. 正则表达式演示
// ==================================================

// demonstrateRegex 演示正则表达式功能
func demonstrateRegex() {
	fmt.Println("=== 正则表达式演示 ===")

	// 邮箱验证
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRegex := regexp.MustCompile(emailPattern)

	emails := []string{
		"user@example.com",
		"invalid.email",
		"test@domain.org",
		"bad@",
	}

	fmt.Println("邮箱验证:")
	for _, email := range emails {
		isValid := emailRegex.MatchString(email)
		status := "❌"
		if isValid {
			status = "✅"
		}
		fmt.Printf("  %s %s\n", status, email)
	}

	// 提取数字
	text := "我有3个苹果，5个橙子，和10个香蕉"
	numberRegex := regexp.MustCompile(`\d+`)
	numbers := numberRegex.FindAllString(text, -1)
	fmt.Printf("\n从'%s'中提取的数字: %v\n", text, numbers)

	// 替换操作
	phoneText := "联系电话: 138-1234-5678 或 159-8765-4321"
	phoneRegex := regexp.MustCompile(`(\d{3})-(\d{4})-(\d{4})`)
	formatted := phoneRegex.ReplaceAllString(phoneText, "($1) $2-$3")
	fmt.Printf("电话格式化: %s\n", formatted)
	fmt.Println()
}

// ==================================================
// 6. 模板引擎演示
// ==================================================

// demonstrateTemplate 演示模板引擎功能
func demonstrateTemplate() {
	fmt.Println("=== 模板引擎演示 ===")

	// 简单模板
	simpleTemplate := `欢迎 {{.Name}}！
您今年 {{.Age}} 岁，来自 {{.City}}。`

	tmpl := template.Must(template.New("person").Parse(simpleTemplate))

	person := Person{Name: "张三", Age: 25, City: "北京"}

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, person)
	if err != nil {
		fmt.Printf("模板执行错误: %v\n", err)
		return
	}

	fmt.Println("简单模板输出:")
	fmt.Println(buf.String())

	// 复杂模板（带循环）
	listTemplate := `学生名单:
{{range .}}
- {{.Name}} ({{.Age}}岁, {{.City}})
{{end}}`

	listTmpl := template.Must(template.New("list").Parse(listTemplate))

	students := []Person{
		{"小明", 20, "上海"},
		{"小红", 22, "北京"},
		{"小华", 21, "广州"},
	}

	fmt.Println("\n列表模板输出:")
	err = listTmpl.Execute(os.Stdout, students)
	if err != nil {
		fmt.Printf("模板执行错误: %v\n", err)
	}
	fmt.Println()
}

// ==================================================
// 7. 文件操作演示
// ==================================================

// demonstrateFileOps 演示文件操作
func demonstrateFileOps() {
	fmt.Println("=== 文件操作演示 ===")

	fileName := "test_output.txt"
	content := `这是一个测试文件
包含多行内容
用于演示Go语言的文件操作`

	// 写入文件
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 成功写入文件: %s\n", fileName)

	// 读取文件
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)
		return
	}

	fmt.Println("文件内容:")
	fmt.Println(string(data))

	// 清理文件
	err = os.Remove(fileName)
	if err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功删除文件: %s\n", fileName)
	}
	fmt.Println()
}

// ==================================================
// 8. 综合演示函数
// ==================================================

// testSortingAndAdvanced 排序和高级功能综合测试
func testSortingAndAdvanced() {
	fmt.Println("========== Go排序和高级功能演示 ==========")

	demonstrateBasicSort()
	demonstrateCustomSort()
	demonstrateStructSort()
	demonstrateStringProcessing()
	demonstrateRegex()
	demonstrateTemplate()
	demonstrateFileOps()

	fmt.Println("========================================")
	fmt.Println("           演示完成")
	fmt.Println("========================================")
}

// main 主函数 - 取消注释可直接运行
// func main() {
// 	testSortingAndAdvanced()
// }
