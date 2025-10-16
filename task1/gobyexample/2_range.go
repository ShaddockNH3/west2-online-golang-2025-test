package gobyexample

import (
	"fmt"
)

// ==================================================
// 1. 指针操作
// ==================================================

// demonstratePointers 演示指针的基本操作
func demonstratePointers() {
	fmt.Println("=== 指针操作演示 ===")

	// 基本指针操作
	value := 42
	ptr := &value // 获取value的地址

	fmt.Printf("原始值: %d\n", value)
	fmt.Printf("变量地址: %p\n", &value)
	fmt.Printf("指针值(地址): %p\n", ptr)
	fmt.Printf("指针指向的值: %d\n", *ptr)

	// 通过指针修改值
	*ptr = 100
	fmt.Printf("通过指针修改后的值: %d\n", value)

	// 指针比较
	if &value == ptr {
		fmt.Println("✓ 指针地址验证成功")
	}

	// 空指针
	var nullPtr *int
	fmt.Printf("空指针: %v (是否为nil: %t)\n", nullPtr, nullPtr == nil)

	fmt.Println()
}

// ==================================================
// 2. 字符串遍历
// ==================================================

// demonstrateStringTraversal 演示字符串的不同遍历方式
func demonstrateStringTraversal() {
	fmt.Println("=== 字符串遍历演示 ===")
	text := "Hello世界"

	fmt.Printf("字符串: %s (长度: %d字节)\n", text, len(text))

	// 按字节遍历
	fmt.Println("\n按字节遍历:")
	for i := 0; i < len(text); i++ {
		fmt.Printf("  字节%d: %d ('%c')\n", i, text[i], text[i])
	}

	// 按Unicode字符遍历(推荐方式)
	fmt.Println("\n按Unicode字符遍历:")
	for i, char := range text {
		fmt.Printf("  位置%d: Unicode %d, 字符 '%c'\n", i, char, char)
	}

	// 字符串切片
	fmt.Println("\n字符串切片:")
	fmt.Printf("前5个字节: %s\n", text[:5])
	fmt.Printf("从第6个字节开始: %s\n", text[5:])

	// 字符串到字节切片和反向转换
	bytes := []byte(text)
	fmt.Printf("转为字节切片: %v\n", bytes)
	fmt.Printf("从字节切片还原: %s\n", string(bytes))

	fmt.Println()
}

// ==================================================
// 3. 接口定义
// ==================================================

// Shouter 接口定义 - 所有能发声的动物都要实现这个接口
type Shouter interface {
	Shout() string
}

// Behavior 扩展接口 - 动物的行为接口
type Behavior interface {
	Shouter
	GetInfo() string
}

// ==================================================
// 4. 结构体定义 - 猫
// ==================================================

// Cat 猫的结构体定义
type Cat struct {
	Name  string
	Age   int
	Breed string
}

// NewCat 猫的构造函数
func NewCat(name, breed string) *Cat {
	return &Cat{
		Name:  name,
		Age:   1, // 默认1岁
		Breed: breed,
	}
}

// Meow 猫的叫声方法
func (c *Cat) Meow() {
	fmt.Printf("%s(%s): 喵喵喵~\n", c.Name, c.Breed)
}

// Shout 实现Shouter接口
func (c *Cat) Shout() string {
	return fmt.Sprintf("%s: 喵喵!", c.Name)
}

// GetInfo 实现Behavior接口
func (c *Cat) GetInfo() string {
	return fmt.Sprintf("猫咪: %s, 品种: %s, 年龄: %d岁", c.Name, c.Breed, c.Age)
}

// CelebrateBirthday 猫咪过生日
func (c *Cat) CelebrateBirthday() {
	c.Age++
	fmt.Printf("%s现在%d岁了!\n", c.Name, c.Age)
}

// ==================================================
// 5. 结构体定义 - 狗
// ==================================================

// Dog 狗的结构体定义
type Dog struct {
	Name   string
	Breed  string
	Weight float64
}

// NewDog 狗的构造函数
func NewDog(name, breed string, weight float64) *Dog {
	return &Dog{
		Name:   name,
		Breed:  breed,
		Weight: weight,
	}
}

// Bark 狗的叫声方法
func (d *Dog) Bark() {
	fmt.Printf("%s(%s): 汪汪汪!\n", d.Name, d.Breed)
}

// Shout 实现Shouter接口
func (d *Dog) Shout() string {
	return fmt.Sprintf("%s: 汪汪!", d.Name)
}

// GetInfo 实现Behavior接口
func (d *Dog) GetInfo() string {
	return fmt.Sprintf("狗狗: %s, 品种: %s, 体重: %.1fkg", d.Name, d.Breed, d.Weight)
}

// ==================================================
// 6. 接口使用函数
// ==================================================

// makeSound 接受任何实现了Shouter接口的类型
func makeSound(s Shouter) {
	fmt.Println("动物叫声:", s.Shout())
}

// showAnimalInfo 展示动物信息（需要实现Behavior接口）
func showAnimalInfo(b Behavior) {
	fmt.Println("动物信息:", b.GetInfo())
	fmt.Println("动物声音:", b.Shout())
}

// ==================================================
// 7. 综合演示函数
// ==================================================

// test2 演示结构体、指针和接口的使用
func test2() {
	fmt.Println("========== 结构体和接口演示 ==========")

	// 演示指针
	demonstratePointers()

	// 演示字符串处理
	demonstrateStringTraversal()

	fmt.Println("=== 动物创建演示 ===")
	// 创建不同的动物
	cat1 := NewCat("小花", "波斯猫")
	cat2 := NewCat("小黑", "英短")
	dog1 := NewDog("旺财", "金毛", 25.5)
	dog2 := NewDog("小白", "拉布拉多", 28.0)

	fmt.Printf("创建了猫咪: %s 和 %s\n", cat1.Name, cat2.Name)
	fmt.Printf("创建了狗狗: %s 和 %s\n", dog1.Name, dog2.Name)

	fmt.Println("\n=== 动物行为演示 ===")
	cat1.Meow()
	cat2.Meow()
	dog1.Bark()
	dog2.Bark()

	// 生日演示
	fmt.Println("\n=== 生日演示 ===")
	fmt.Printf("%s过生日前: %d岁\n", cat1.Name, cat1.Age)
	cat1.CelebrateBirthday()

	fmt.Println("\n=== 接口演示 ===")
	// 通过Shouter接口调用不同类型的方法
	animals := []Shouter{cat1, cat2, dog1, dog2}
	for _, animal := range animals {
		makeSound(animal)
	}

	fmt.Println("\n=== Behavior接口演示 ===")
	// 通过Behavior接口展示详细信息
	behaviorAnimals := []Behavior{cat1, dog1}
	for _, animal := range behaviorAnimals {
		showAnimalInfo(animal)
		fmt.Println()
	}

	fmt.Println("演示完成!")
}

// main 主函数 - 取消注释可直接运行
// func main() {
// 	fmt.Println("=== Go语言学习 - 结构体和接口 ===")
// 	test2()
// }
