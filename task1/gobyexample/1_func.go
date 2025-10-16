package gobyexample

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"time"
)

// ==================================================
// 1. 基础输出与打印
// ==================================================

// hello_world 演示基本的输出功能
func hello_world() {
	fmt.Println("=== Hello World 演示 ===")
	fmt.Println("Hello World!")
	fmt.Println("Cheer up, ShaddockNH3")
	fmt.Println()
}

// ==================================================
// 2. 基本值类型
// ==================================================

// value_test 演示基本值类型
func value_test() {
	fmt.Println("=== 基本值类型演示 ===")
	fmt.Println("整数:", 42)
	fmt.Println("浮点数:", 3.14)
	fmt.Println("布尔值:", true)
	fmt.Println("字符串:", "Hello Go")
	fmt.Println("除法运算:", 7.0/3.0)
	fmt.Println()
}

// ==================================================
// 3. 变量声明与初始化
// ==================================================

// variables_test 演示变量声明和初始化的不同方式
func variables_test() {
	fmt.Println("=== 变量声明演示 ===")

	// 完整声明并初始化
	var a int = 1
	fmt.Printf("完整声明: var a int = 1, a = %d\n", a)

	// 类型推导
	var b = 2
	fmt.Printf("类型推导: var b = 2, b = %d\n", b)

	// 短声明语法（最常用）
	c := 3
	fmt.Printf("短声明: c := 3, c = %d\n", c)

	// 字符串变量
	var d string = "完整声明字符串"
	var e = "类型推导字符串"
	f := "短声明字符串"
	fmt.Printf("字符串变量: d = %s, e = %s, f = %s\n", d, e, f)

	// 零值变量（未初始化）
	var g int
	var h string
	var i bool
	fmt.Printf("零值变量: int=%d, string='%s', bool=%t\n", g, h, i)

	// 浮点数
	var j float32 = 3.14
	k := 2.718 // 默认为float64
	fmt.Printf("浮点数: float32=%.2f, float64=%.3f\n", j, k)

	fmt.Println()
}

// ==================================================
// 4. 常量定义
// ==================================================

// const_test 演示常量的使用
func const_test() {
	fmt.Println("=== 常量演示 ===")
	const PI = 3.14159
	const GREETING = "Hello"
	const MAX_COUNT = 100

	fmt.Printf("数学常量 PI = %.5f\n", PI)
	fmt.Printf("字符串常量 GREETING = %s\n", GREETING)
	fmt.Printf("整数常量 MAX_COUNT = %d\n", MAX_COUNT)
	fmt.Printf("数学计算: sin(π/2) = %.3f\n", math.Sin(PI/2))
	fmt.Println()
}

// ==================================================
// 5. 循环结构
// ==================================================

// for_test 演示不同的循环语法
func for_test() {
	fmt.Println("=== 循环结构演示 ===")

	// 标准for循环
	fmt.Println("标准for循环 (1-5):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while风格的for循环
	fmt.Println("while风格循环:")
	n := 1
	for n <= 3 {
		fmt.Printf("n = %d ", n)
		n++
	}
	fmt.Println()

	// 无限循环（需要break）
	fmt.Println("带break的无限循环:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("i = %d ", i)
		i++
	}
	fmt.Println()

	// range循环（Go 1.22+）
	fmt.Println("range循环 (0-4):")
	for i := range 5 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// ==================================================
// 6. 条件语句
// ==================================================

// if_else_test 演示条件语句
func if_else_test() {
	fmt.Println("=== 条件语句演示 ===")

	age := 20
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	// 带初始化的if语句
	if score := 85; score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("一般")
	} else {
		fmt.Println("需要努力")
	}

	// 时间相关的条件判断
	hour := time.Now().Hour()
	if hour < 12 {
		fmt.Println("上午好")
	} else if hour < 18 {
		fmt.Println("下午好")
	} else {
		fmt.Println("晚上好")
	}
	fmt.Println()
}

// ==================================================
// 7. Switch语句
// ==================================================

// switch_test 演示switch语句
func switch_test() {
	fmt.Println("=== Switch语句演示 ===")

	// 基础switch
	day := 3
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	default:
		fmt.Println("周末")
	}

	// 多值匹配
	letter := "A"
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println("元音字母")
	default:
		fmt.Println("辅音字母")
	}

	// 表达式switch
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A级")
	case score >= 80:
		fmt.Println("B级")
	case score >= 70:
		fmt.Println("C级")
	default:
		fmt.Println("需要努力")
	}

	// 类型switch
	var x interface{} = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("整数: %d\n", v)
	case string:
		fmt.Printf("字符串: %s\n", v)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
	fmt.Println()
}

// ==================================================
// 8. 数组
// ==================================================

// arr_test 演示数组的使用
func arr_test() {
	fmt.Println("=== 数组演示 ===")

	// 声明固定长度数组
	var a [5]int
	fmt.Printf("零值数组: %v (长度: %d)\n", a, len(a))

	// 修改数组元素
	a[0] = 100
	a[4] = 500
	fmt.Printf("修改后: %v\n", a)

	// 数组初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("初始化数组: %v\n", b)

	// 自动推导长度
	c := [...]string{"apple", "banana", "cherry"}
	fmt.Printf("自动长度数组: %v (长度: %d)\n", c, len(c))

	// 指定索引初始化
	d := [5]int{1: 10, 3: 30}
	fmt.Printf("指定索引初始化: %v\n", d)

	// 数组遍历
	fmt.Println("数组遍历:")
	for i, v := range b {
		fmt.Printf("  索引%d: %d\n", i, v)
	}
	fmt.Println()
}

// ==================================================
// 9. 切片 (Slice)
// ==================================================

// slice_test 演示切片的使用
func slice_test() {
	fmt.Println("=== 切片演示 ===")

	// 零值切片
	var s []string
	fmt.Printf("零值切片: %v (是否为nil: %t, 长度: %d)\n", s, s == nil, len(s))

	// 使用append添加元素
	s = append(s, "apple", "banana", "cherry")
	fmt.Printf("添加元素后: %v (长度: %d, 容量: %d)\n", s, len(s), cap(s))

	// 使用make创建切片
	nums := make([]int, 3, 5) // 长度3，容量5
	fmt.Printf("make创建: %v (长度: %d, 容量: %d)\n", nums, len(nums), cap(nums))
	nums[0] = 10
	nums = append(nums, 40, 50)
	fmt.Printf("修改后: %v (长度: %d, 容量: %d)\n", nums, len(nums), cap(nums))

	// 切片字面量
	fruits := []string{"apple", "banana", "cherry", "date"}
	fmt.Printf("字面量切片: %v\n", fruits)

	// 切片操作
	fmt.Printf("切片 [1:3]: %v\n", fruits[1:3])
	fmt.Printf("切片 [:2]: %v\n", fruits[:2])
	fmt.Printf("切片 [2:]: %v\n", fruits[2:])

	// 切片比较
	fruits2 := []string{"apple", "banana", "cherry", "date"}
	if slices.Equal(fruits, fruits2) {
		fmt.Println("两个切片相等")
	}

	// 复制切片
	copied := make([]string, len(fruits))
	copy(copied, fruits)
	fmt.Printf("复制的切片: %v\n", copied)
	fmt.Println()
}

// ==================================================
// 10. 映射 (Map)
// ==================================================

// map_test 演示映射(map)的使用
func map_test() {
	fmt.Println("=== 映射演示 ===")

	// 创建空map
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Printf("学生成绩: %v\n", scores)

	// map字面量
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Printf("颜色映射: %v\n", colors)

	// 访问元素
	if score, exists := scores["Alice"]; exists {
		fmt.Printf("Alice的成绩: %d\n", score)
	}

	// 修改元素
	scores["David"] = 89
	fmt.Printf("添加David后: %v\n", scores)

	// 删除元素
	delete(scores, "Bob")
	fmt.Printf("删除Bob后: %v\n", scores)

	// 遍历map
	fmt.Println("遍历成绩:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}

	// 清空map
	clear(scores)
	fmt.Printf("清空后: %v (长度: %d)\n", scores, len(scores))

	// map比较
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"a": 1, "b": 2}
	if maps.Equal(map1, map2) {
		fmt.Println("两个map相等")
	}
	fmt.Println()
}

// ==================================================
// 11. 函数定义与调用
// ==================================================

// add 基础函数 - 两数相加
func add(a, b int) int {
	return a + b
}

// calculate 多参数函数 - 三数运算
func calculate(a, b, c int) (int, float64) {
	sum := a + b + c
	average := float64(sum) / 3
	return sum, average
}

// greet 带命名返回值的函数
func greet(name string) (message string) {
	message = "Hello, " + name + "!"
	return // 裸返回
}

// sum 可变参数函数
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// processor 函数作为参数
func processor(numbers []int, operation func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = operation(num)
	}
	return result
}

// counter 闭包函数
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// factorial 递归函数 - 计算阶乘
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// ==================================================
// 12. 综合测试函数
// ==================================================

// test 主测试函数，演示所有基础语法
func test() {
	fmt.Println("========== Go基础语法演示 ==========")

	// 基础语法演示
	hello_world()
	value_test()
	variables_test()
	const_test()
	for_test()
	if_else_test()
	switch_test()
	arr_test()
	slice_test()
	map_test()

	// 函数功能演示
	fmt.Println("=== 函数功能演示 ===")
	fmt.Printf("加法: add(10, 20) = %d\n", add(10, 20))

	total, avg := calculate(10, 20, 30)
	fmt.Printf("计算结果: 和=%d, 平均值=%.2f\n", total, avg)

	fmt.Printf("问候: %s\n", greet("Go语言"))

	fmt.Printf("可变参数求和: sum(1,2,3,4,5) = %d\n", sum(1, 2, 3, 4, 5))

	// 高阶函数演示
	numbers := []int{1, 2, 3, 4, 5}
	doubled := processor(numbers, func(x int) int { return x * 2 })
	fmt.Printf("数组翻倍: %v -> %v\n", numbers, doubled)

	// 闭包演示
	fmt.Println("\n=== 闭包演示 ===")
	count := counter()
	for i := 0; i < 5; i++ {
		fmt.Printf("闭包调用 %d: %d\n", i+1, count())
	}

	// 递归演示
	fmt.Println("\n=== 递归演示 ===")
	fmt.Printf("阶乘: factorial(5) = %d\n", factorial(5))

	// 匿名函数递归
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	fmt.Printf("斐波那契数列第10项: %d\n", fibonacci(10))

	// 字符串遍历
	fmt.Println("\n=== 字符串遍历 ===")
	text := "Hello世界"
	for i, char := range text {
		fmt.Printf("位置%d: 字符'%c' (Unicode: %d)\n", i, char, char)
	}

	fmt.Println("\n========== 演示完成 ==========")
}

// main 主函数 - 取消注释可直接运行
// func main() {
//     test()
// }
