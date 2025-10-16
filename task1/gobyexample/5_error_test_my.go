package gobyexample

import (
	"errors"
	"fmt"
)

// ==================================================
// 1. 自定义错误类型
// ==================================================

// ArgError 自定义错误类型，包含参数和错误信息
type ArgError struct {
	Arg     int
	Message string
}

// Error 实现error接口
func (e *ArgError) Error() string {
	return fmt.Sprintf("参数错误 [%d]: %s", e.Arg, e.Message)
}

// ==================================================
// 2. 预定义错误变量
// ==================================================

var (
	ErrOutOfTea = errors.New("茶叶用完了")
	ErrNoPower  = errors.New("停电了，无法烧水")
	ErrNoWater  = errors.New("没有水了")
)

// ==================================================
// 3. 基础错误处理函数
// ==================================================

// calculateWithError 基础函数，演示简单错误处理
func calculateWithError(arg int) (int, error) {
	if arg == 42 {
		return -1, &ArgError{arg, "不能处理数字42"}
	}
	if arg < 0 {
		return -1, errors.New("参数不能为负数")
	}
	return arg + 3, nil
}

// divide 除法函数，演示数学错误
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// makeTea 制茶函数，演示错误包装
func makeTea(teaType int) error {
	switch teaType {
	case 1:
		return ErrOutOfTea
	case 2:
		return ErrNoPower
	case 3:
		return ErrNoWater
	case 4:
		// 错误包装示例
		return fmt.Errorf("制茶失败: %w", ErrNoPower)
	default:
		return nil
	}
}

// ==================================================
// 4. 错误处理演示函数
// ==================================================

// demonstrateBasicErrors 演示基础错误处理
func demonstrateBasicErrors() {
	fmt.Println("=== 基础错误处理演示 ===")

	// 测试不同的输入
	testValues := []int{5, 42, -3, 10}
	for _, val := range testValues {
		result, err := calculateWithError(val)
		if err != nil {
			fmt.Printf("calculateWithError(%d) 失败: %v\n", val, err)
		} else {
			fmt.Printf("calculateWithError(%d) = %d\n", val, result)
		}
	}
	fmt.Println()
}

// demonstrateCustomErrors 演示自定义错误类型
func demonstrateCustomErrors() {
	fmt.Println("=== 自定义错误类型演示 ===")

	_, err := calculateWithError(42)

	// 使用 errors.As 检查特定错误类型
	var argErr *ArgError
	if errors.As(err, &argErr) {
		fmt.Printf("捕获到自定义错误: 参数=%d, 消息=%s\n", argErr.Arg, argErr.Message)
	} else {
		fmt.Println("不是自定义错误类型")
	}
	fmt.Println()
}

// demonstrateErrorChecking 演示错误检查和包装
func demonstrateErrorChecking() {
	fmt.Println("=== 错误检查和包装演示 ===")

	// 测试制茶功能
	for i := 0; i <= 5; i++ {
		err := makeTea(i)
		if err != nil {
			// 使用 errors.Is 检查特定错误
			if errors.Is(err, ErrOutOfTea) {
				fmt.Printf("制茶类型%d: 需要购买新茶叶!\n", i)
			} else if errors.Is(err, ErrNoPower) {
				fmt.Printf("制茶类型%d: 现在停电了\n", i)
			} else if errors.Is(err, ErrNoWater) {
				fmt.Printf("制茶类型%d: 需要加水\n", i)
			} else {
				fmt.Printf("制茶类型%d: 未知错误: %v\n", i, err)
			}
		} else {
			fmt.Printf("制茶类型%d: 茶准备好了!\n", i)
		}
	}
	fmt.Println()
}

// demonstrateMathErrors 演示数学运算错误
func demonstrateMathErrors() {
	fmt.Println("=== 数学运算错误演示 ===")

	// 测试除法
	testCases := [][2]float64{{10, 2}, {10, 0}, {7, 3}}
	for _, tc := range testCases {
		result, err := divide(tc[0], tc[1])
		if err != nil {
			fmt.Printf("%.1f ÷ %.1f: 错误 - %v\n", tc[0], tc[1], err)
		} else {
			fmt.Printf("%.1f ÷ %.1f = %.2f\n", tc[0], tc[1], result)
		}
	}
	fmt.Println()
}

// testErrorHandling 综合错误处理测试
func testErrorHandling() {
	fmt.Println("========== Go错误处理演示 ==========")

	demonstrateBasicErrors()
	demonstrateCustomErrors()
	demonstrateErrorChecking()
	demonstrateMathErrors()

	fmt.Println("错误处理演示完成!")
}

// main 主函数 - 取消注释可直接运行
// func main() {
// 	testErrorHandling()
// }
