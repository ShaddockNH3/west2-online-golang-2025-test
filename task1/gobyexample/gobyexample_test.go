package gobyexample

import (
	"testing"
)

// 1_func.go tests
func Test1_HelloWorld(t *testing.T) {
	t.Log("=== 测试: Hello World ===")
	hello_world()
}

func Test1_ValueTest(t *testing.T) {
	t.Log("=== 测试: 基本值类型 ===")
	value_test()
}

func Test1_VariablesTest(t *testing.T) {
	t.Log("=== 测试: 变量声明 ===")
	variables_test()
}

func Test1_ConstTest(t *testing.T) {
	t.Log("=== 测试: 常量 ===")
	const_test()
}

func Test1_ForTest(t *testing.T) {
	t.Log("=== 测试: 循环结构 ===")
	for_test()
}

func Test1_IfElseTest(t *testing.T) {
	t.Log("=== 测试: 条件语句 ===")
	if_else_test()
}

func Test1_SwitchTest(t *testing.T) {
	t.Log("=== 测试: Switch语句 ===")
	switch_test()
}

func Test1_ArrTest(t *testing.T) {
	t.Log("=== 测试: 数组 ===")
	arr_test()
}

func Test1_SliceTest(t *testing.T) {
	t.Log("=== 测试: 切片 ===")
	slice_test()
}

func Test1_MapTest(t *testing.T) {
	t.Log("=== 测试: 映射 ===")
	map_test()
}

func Test1_AllBasics(t *testing.T) {
	t.Log("=== 测试: 所有基础语法 ===")
	test()
}

// 2_range.go tests
func Test2_DemonstratePointers(t *testing.T) {
	t.Log("=== 测试: 指针操作 ===")
	demonstratePointers()
}

func Test2_DemonstrateStringTraversal(t *testing.T) {
	t.Log("=== 测试: 字符串遍历 ===")
	demonstrateStringTraversal()
}

func Test2_StructAndInterface(t *testing.T) {
	t.Log("=== 测试: 结构体和接口 ===")
	test2()
}

// 3_struct.go tests
func Test3_Shelter(t *testing.T) {
	t.Log("=== 测试: 动物收容所 ===")
	test3()
}

// 4_after_struct.go tests
func Test4_Enums(t *testing.T) {
	t.Log("=== 测试: 枚举 ===")
	enums()
}

func Test4_Generics(t *testing.T) {
	t.Log("=== 测试: 泛型 ===")
	Generics()
}

func Test4_Generics1(t *testing.T) {
	t.Log("=== 测试: 泛型函数 ===")
	Generics1()
}

func Test4_ListTest(t *testing.T) {
	t.Log("=== 测试: 泛型链表 ===")
	List_Test()
}

func Test4_Fibtest(t *testing.T) {
	t.Log("=== 测试: 斐波那契 ===")
	fibtest()
}

// 5_error_test_my.go tests
func Test5_DemonstrateBasicErrors(t *testing.T) {
	t.Log("=== 测试: 基础错误处理 ===")
	demonstrateBasicErrors()
}

func Test5_DemonstrateCustomErrors(t *testing.T) {
	t.Log("=== 测试: 自定义错误 ===")
	demonstrateCustomErrors()
}

func Test5_DemonstrateErrorChecking(t *testing.T) {
	t.Log("=== 测试: 错误检查 ===")
	demonstrateErrorChecking()
}

func Test5_TestErrorHandling(t *testing.T) {
	t.Log("=== 测试: 错误处理综合 ===")
	testErrorHandling()
}

// 6_goroutines.go tests
func Test6_Goroutines(t *testing.T) {
	t.Log("=== 测试: Goroutines ===")
	test6()
}

func Test6_Directions(t *testing.T) {
	t.Log("=== 测试: 通道方向 ===")
	directions()
}

func Test6_SelectTry(t *testing.T) {
	t.Log("=== 测试: Select语句 ===")
	select_try()
}

func Test6_NonBlocking(t *testing.T) {
	t.Log("=== 测试: 非阻塞通道 ===")
	Non_Blocking()
}

func Test6_CloseTry(t *testing.T) {
	t.Log("=== 测试: 关闭通道 ===")
	close_try()
}

// 7_timers.go tests
func Test7_TimeTry(t *testing.T) {
	t.Log("=== 测试: 定时器 ===")
	time_try()
}

func Test7_TickerTry(t *testing.T) {
	t.Log("=== 测试: Ticker ===")
	ticker_try()
}

func Test7_WorkerPool(t *testing.T) {
	t.Log("=== 测试: Worker Pool ===")
	worker_pool()
}

func Test7_WaitGroup(t *testing.T) {
	t.Log("=== 测试: Wait Group ===")
	wait_group()
}

func Test7_RatingLimit(t *testing.T) {
	t.Log("=== 测试: 速率限制 ===")
	rating_limit()
}

func Test7_AtomicCounters(t *testing.T) {
	t.Log("=== 测试: 原子计数器 ===")
	Atomic_Counters()
}

func Test7_Mutexes(t *testing.T) {
	t.Log("=== 测试: 互斥锁 ===")
	Mutexes()
}

func Test7_StatefulGoroutines(t *testing.T) {
	t.Log("=== 测试: 有状态的Goroutines ===")
	Stateful_Goroutines()
}

// 8_foundation.go tests
func Test8_Foundation(t *testing.T) {
	t.Log("=== 测试: 基础综合练习 ===")
	testFoundation()
}

// 11_sort.go tests
func Test11_DemonstrateBasicSort(t *testing.T) {
	t.Log("=== 测试: 基础排序 ===")
	demonstrateBasicSort()
}

func Test11_DemonstrateCustomSort(t *testing.T) {
	t.Log("=== 测试: 自定义排序 ===")
	demonstrateCustomSort()
}

func Test11_DemonstrateStructSort(t *testing.T) {
	t.Log("=== 测试: 结构体排序 ===")
	demonstrateStructSort()
}

func Test11_DemonstrateStringProcessing(t *testing.T) {
	t.Log("=== 测试: 字符串处理 ===")
	demonstrateStringProcessing()
}

func Test11_DemonstrateRegex(t *testing.T) {
	t.Log("=== 测试: 正则表达式 ===")
	demonstrateRegex()
}

func Test11_DemonstrateTemplate(t *testing.T) {
	t.Log("=== 测试: 模板引擎 ===")
	demonstrateTemplate()
}

func Test11_DemonstrateFileOps(t *testing.T) {
	t.Log("=== 测试: 文件操作 ===")
	demonstrateFileOps()
}

func Test11_SortingAndAdvanced(t *testing.T) {
	t.Log("=== 测试: 排序与高级功能综合 ===")
	testSortingAndAdvanced()
}
