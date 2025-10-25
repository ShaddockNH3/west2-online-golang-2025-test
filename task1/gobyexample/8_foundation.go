package gobyexample

import (
	"fmt"
	"unicode/utf8"
)

// ==================================================
// 数据结构定义
// ==================================================

// studentRoster 学生名册 - 存储学生姓名和分数
var studentRoster = map[string]int{
	"张三": 85,
	"李四": 92,
	"王五": 78,
	"🐱":  100, // 支持Unicode字符
}

// ==================================================
// 核心功能函数
// ==================================================

// printWelcomeMessage 打印欢迎信息
func printWelcomeMessage() {
	fmt.Println("========================================")
	fmt.Println("      欢迎使用学生成绩管理系统")
	fmt.Println("========================================")
}

// showAllStudents 显示所有学生信息
func showAllStudents(roster *map[string]int) {
	fmt.Println("\n=== 当前学生名单 ===")
	if len(*roster) == 0 {
		fmt.Println("暂无学生记录")
		return
	}

	count := 0
	for name, score := range *roster {
		count++
		grade := getGrade(score)
		fmt.Printf("%d. 姓名: %-8s | 分数: %3d | 等级: %s\n", count, name, score, grade)
	}
	fmt.Printf("总计: %d 名学生\n", len(*roster))
}

// getGrade 根据分数返回等级
func getGrade(score int) string {
	switch {
	case score >= 90:
		return "优秀"
	case score >= 80:
		return "良好"
	case score >= 70:
		return "中等"
	case score >= 60:
		return "及格"
	default:
		return "不及格"
	}
}

// addStudent 添加学生到名册
func addStudent(roster *map[string]int, name string, score int) bool {
	if _, exists := (*roster)[name]; exists {
		fmt.Printf("❌ 错误: 学生 %s 已经存在于系统中\n", name)
		return false
	}
	(*roster)[name] = score
	fmt.Printf("✅ 成功添加学生: %s, 分数: %d\n", name, score)
	return true
}

// calculateAverage 计算平均分（可变参数）
func calculateAverage(scores ...int) float64 {
	if len(scores) == 0 {
		return 0.0
	}
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64(total) / float64(len(scores))
}

// makeGreeter 创建问候函数（闭包演示）
func makeGreeter(greeting string) func(string) {
	return func(name string) {
		fmt.Printf("🗣️ %s, %s!\n", greeting, name)
	}
}

// factorialCalc 计算阶乘（递归演示）
func factorialCalc(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialCalc(n-1)
}

// analyzeName 分析姓名的详细信息
func analyzeName(name string) {
	fmt.Printf("📝 分析姓名: %s\n", name)
	fmt.Printf("   字节长度: %d\n", len(name))
	fmt.Printf("   字符数量: %d\n", utf8.RuneCountInString(name))
	fmt.Printf("   字符详情: ")
	for i, char := range name {
		fmt.Printf("[%d]'%c' ", i, char)
	}
	fmt.Println()
}

// calculatePassRate 计算及格率
func calculatePassRate(scores []int, passScore int) float64 {
	if len(scores) == 0 {
		return 0.0
	}

	passCount := 0
	for _, score := range scores {
		if score >= passScore {
			passCount++
		}
	}

	return float64(passCount) / float64(len(scores)) * 100
}

// ==================================================
// 主程序演示
// ==================================================

// testFoundation 基础功能综合测试
func testFoundation() {
	// 定义常量
	const PASS_SCORE = 60
	const EXCELLENT_SCORE = 90

	// 1. 打印欢迎信息
	printWelcomeMessage()

	// 2. 显示初始学生名单
	showAllStudents(&studentRoster)

	// 3. 添加新学生
	fmt.Println("\n=== 添加新学生 ===")
	addStudent(&studentRoster, "小明", 88)
	addStudent(&studentRoster, "小红", 95)
	addStudent(&studentRoster, "小明", 99) // 重复添加测试

	// 4. 批量添加候选学生
	fmt.Println("\n=== 批量添加候选学生 ===")
	candidates := []string{"小华", "小丽", "小强"}
	defaultScore := 75
	for i, name := range candidates {
		score := defaultScore + i*5 // 给不同的分数
		addStudent(&studentRoster, name, score)
	}

	// 5. 显示更新后的名单
	showAllStudents(&studentRoster)

	// 6. 统计分析
	fmt.Println("\n=== 统计分析 ===")
	var scores []int
	excellentCount := 0
	for name, score := range studentRoster {
		scores = append(scores, score)
		if score >= EXCELLENT_SCORE {
			excellentCount++
			fmt.Printf("🌟 优秀学生: %s (分数: %d)\n", name, score)
		}
	}

	average := calculateAverage(scores...)
	fmt.Printf("\n📊 班级统计信息:\n")
	fmt.Printf("- 总人数: %d\n", len(studentRoster))
	fmt.Printf("- 平均分: %.2f\n", average)
	fmt.Printf("- 优秀学生数: %d\n", excellentCount)
	fmt.Printf("- 及格率: %.1f%%\n", calculatePassRate(scores, PASS_SCORE))

	// 7. 闭包函数演示
	fmt.Println("\n=== 问候功能演示 ===")
	sayHello := makeGreeter("你好")
	sayGoodbye := makeGreeter("再见")
	sayHello("张三")
	sayGoodbye("李四")

	// 8. 递归函数演示
	fmt.Println("\n=== 数学计算演示 ===")
	for i := 1; i <= 6; i++ {
		fmt.Printf("%d! = %d\n", i, factorialCalc(i))
	}

	// 9. 字符串分析演示
	fmt.Println("\n=== 姓名分析演示 ===")
	testNames := []string{"张三", "🐱", "Alice", "李小明"}
	for _, name := range testNames {
		analyzeName(name)
	}

	fmt.Println("\n========================================")
	fmt.Println("           程序执行完成")
	fmt.Println("========================================")
}

// main 主函数 - 取消注释可直接运行
// func main() {
// 	testFoundation()
// }
