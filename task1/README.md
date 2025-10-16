# Task1 - Go 语言学习与练习

> Go 语言基础学习记录与实践项目
> 
> 更新日期：2025年10月

## 📁 项目结构

```
task1/
├── gobyexample/          # Go 语言系统学习（基于 Go by Example）
│   ├── 1_func.go         # 函数与基础语法大全
│   ├── 2_range.go        # 指针、字符串与接口
│   ├── 3_struct.go       # 结构体与动物收容所示例
│   ├── 4_after_struct.go # 枚举、泛型、迭代器
│   ├── 5_error_test_my.go # 错误处理机制
│   ├── 6_goroutines.go   # 并发编程：Goroutines & Channels
│   ├── 7_timers.go       # 定时器、协程池、并发控制
│   ├── 8_foundation.go   # 综合练习：学生成绩管理系统
│   ├── 11_sort.go        # 排序、正则、模板、文件操作
│   └── gobyexample_test.go # 测试文件（46个测试用例）
│
├── task/
│   ├── base/             # 基础算法题
│   │   ├── luogu_p1001.go    # A+B Problem
│   │   ├── luogu_p1046.go    # 摘苹果
│   │   ├── luogu_p5737.go    # 闰年判断
│   │   └── base_test.go      # 测试文件
│   │
│   └── bonus/            # 额外练习
│       ├── ninenine.go       # 九九乘法表生成
│       └── ninenine_test.go  # 测试文件
│
└── study/                # 学习笔记与实验（已废弃）
```

## 🎯 学习内容

### 1. gobyexample - 系统学习模块

基于 [Go by Example](https://gobyexample.com/) 的系统学习，所有代码均使用 **测试驱动** 方式组织。

#### 📚 知识点覆盖

| 文件 | 主要内容 | 测试数量 |
|------|---------|---------|
| `1_func.go` | 变量、常量、循环、条件、数组、切片、Map、函数 | 11 |
| `2_range.go` | 指针操作、字符串遍历、接口与多态 | 3 |
| `3_struct.go` | 结构体、接口实现、动物收容所示例 | 1 |
| `4_after_struct.go` | 枚举、泛型、泛型链表、斐波那契生成器 | 5 |
| `5_error_test_my.go` | 错误处理、自定义错误、错误包装 | 4 |
| `6_goroutines.go` | Goroutines、Channels、Select、通道关闭 | 5 |
| `7_timers.go` | Timer、Ticker、Worker Pool、WaitGroup、原子操作、互斥锁 | 8 |
| `8_foundation.go` | 学生成绩管理系统（综合练习） | 1 |
| `11_sort.go` | 排序算法、正则表达式、模板引擎、文件操作 | 8 |

**总计：46 个测试用例**

#### 🚀 运行测试

```bash
# 进入 gobyexample 目录
cd task1/gobyexample

# 运行所有测试（查看学习记录）
go test -v

# 运行特定测试
go test -v -run Test1_HelloWorld
go test -v -run Test6_Goroutines
go test -v -run Test7_WorkerPool
```

#### 📖 核心概念示例

**1. 基础语法（1_func.go）**
- 变量声明的三种方式
- 数组 vs 切片的区别
- Map 的增删改查
- 函数的多返回值、可变参数、闭包、递归

**2. 并发编程（6_goroutines.go, 7_timers.go）**
- Goroutines 的创建与通信
- Channel 的阻塞与非阻塞操作
- Select 多路复用
- Worker Pool 模式
- 原子操作与互斥锁

**3. 错误处理（5_error_test_my.go）**
- 基础错误创建
- 自定义错误类型
- errors.Is 和 errors.As 的使用
- 错误包装与解包

### 2. task - 算法练习模块

#### base 文件夹 - 基础题目

所有题目均已重构为 **可传参测试** 形式，便于单元测试和学习记录。

| 文件 | 题目 | 功能 |
|------|------|------|
| `luogu_p1001.go` | A+B Problem | 两数相加 |
| `luogu_p1046.go` | 摘苹果 | 数组遍历与条件统计 |
| `luogu_p5737.go` | 闰年判断 | 循环与条件判断 |

```bash
# 运行测试
cd task1/task/base
go test -v
```

#### bonus 文件夹 - 额外练习

| 文件 | 功能 | 输出 |
|------|------|------|
| `ninenine.go` | 九九乘法表生成器 | 保存为 txt 文件 |

```bash
# 运行测试并生成乘法表
cd task1/task/bonus
go test -v
# 会在当前目录生成 multiplication_table.txt
```

## 🎓 学习建议

### 新手学习路径

1. **基础语法** → 从 `1_func.go` 开始
   ```bash
   cd gobyexample
   go test -v -run Test1_
   ```

2. **数据结构** → 学习 `2_range.go` 和 `3_struct.go`
   ```bash
   go test -v -run Test2_
   go test -v -run Test3_
   ```

3. **错误处理** → 掌握 `5_error_test_my.go`
   ```bash
   go test -v -run Test5_
   ```

4. **并发编程** → 挑战 `6_goroutines.go` 和 `7_timers.go`
   ```bash
   go test -v -run Test6_
   go test -v -run Test7_
   ```

5. **高级特性** → 探索 `4_after_struct.go` 和 `11_sort.go`
   ```bash
   go test -v -run Test4_
   go test -v -run Test11_
   ```

### 实践练习

1. 修改测试用例中的输入数据
2. 尝试实现相似功能的函数
3. 阅读并理解每个函数的实现逻辑
4. 运行测试查看输出结果

## 📊 测试覆盖

- ✅ **基础语法**：变量、常量、循环、条件、数组、切片、Map
- ✅ **函数**：多返回值、可变参数、闭包、递归
- ✅ **指针**：指针操作、内存地址、指针传递
- ✅ **结构体**：定义、方法、嵌套、组合
- ✅ **接口**：定义、实现、多态、类型断言
- ✅ **并发**：Goroutines、Channels、Select、WaitGroup
- ✅ **同步**：Mutex、原子操作、通道同步
- ✅ **错误处理**：error 接口、自定义错误、错误包装
- ✅ **高级特性**：泛型、枚举、迭代器
- ✅ **标准库**：排序、正则、模板、文件 I/O

## 📝 学习笔记

### 重要概念

1. **切片 vs 数组**
   - 数组长度固定，切片长度可变
   - 切片是对数组的引用
   - 切片有长度（len）和容量（cap）

2. **通道（Channel）**
   - 用于 Goroutines 之间的通信
   - 默认是阻塞的
   - 可以有缓冲区

3. **接口（Interface）**
   - 定义行为规范
   - 隐式实现（无需显式声明）
   - 空接口 `interface{}` 可以接收任何类型

4. **错误处理**
   - Go 使用返回值而非异常
   - `error` 是一个接口类型
   - 推荐使用 `errors.Is` 和 `errors.As` 进行错误检查

## 🔧 开发环境

- **Go 版本**：1.22+
- **包管理**：go.mod
- **测试框架**：标准库 testing

## 📖 参考资料

- [Go by Example](https://gobyexample.com/)
- [Go 官方文档](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go 语言之旅](https://tour.go-zh.org/)

## 💡 改进方向

- [x] 将所有练习改为测试驱动
- [x] 统一包名为 gobyexample
- [x] 删除冗余注释
- [x] 修复类型冲突
- [ ] 添加更多实践项目
- [ ] 增加性能测试（Benchmark）
- [ ] 添加代码覆盖率报告

## 📄 许可

学习项目，仅供参考。

---

**学习记录** - 持续更新中... 🚀
