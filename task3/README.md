# Task3 - 备忘录后端项目 📝

![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)
![Framework](https://img.shields.io/badge/Framework-Hertz-orange.svg)
![ORM](https://img.shields.io/badge/ORM-Gorm-brightgreen.svg)
![License](https://img.shields.io/badge/License-MIT-lightgrey.svg)

这是一个基于 Go 语言和 Hertz 框架构建的备忘录应用后端服务。项目实现了用户的注册、登录、信息管理以及核心的待办事项（ToDoList）增删改查等功能，并使用 JWT 进行用户认证。

---

## ✨ 项目特色

*   **高性能框架**: 基于字节跳动开源的高性能 Go HTTP 框架 [Hertz](https://github.com/cloudwego/hertz)。
*   **接口契约先行**: 使用 Thrift 定义接口，通过 `hz` 工具生成代码，保证前后端协作规范。
*   **清晰的业务分层**: 采用 `biz` 目录对业务逻辑进行清晰分层（Handler, Service, DAL），易于维护和扩展。
*   **强大的数据库支持**: 使用 [Gorm](https://gorm.io/)作为 ORM，简化数据库操作。
*   **安全的用户认证**: 通过 [hertz-jwt](https://github.com/hertz-contrib/jwt) 中间件实现 JWT 认证，保护用户接口。
*   **完善的接口文档**: 使用 `swag` 自动生成交互式的 [Swagger](https://swagger.io/) API 文档。

---

## 📂 项目结构

```task3
 ┣ biz
 ┃ ┣ dal
 ┃ ┃ ┣ mw
 ┃ ┃ ┃ ┗ jwt.go
 ┃ ┃ ┣ mysql
 ┃ ┃ ┃ ┣ init.go
 ┃ ┃ ┃ ┣ todo_list.go
 ┃ ┃ ┃ ┗ user.go
 ┃ ┃ ┗ init.go
 ┃ ┣ handler
 ┃ ┃ ┣ task3
 ┃ ┃ ┃ ┣ to_do_list_service.go
 ┃ ┃ ┃ ┗ user_service.go
 ┃ ┃ ┗ ping.go
 ┃ ┣ model
 ┃ ┃ ┣ sql
 ┃ ┃ ┃ ┣ todo_list.sql
 ┃ ┃ ┃ ┗ user.sql
 ┃ ┃ ┣ task3
 ┃ ┃ ┃ ┗ api.go
 ┃ ┃ ┣ todo_list.go
 ┃ ┃ ┗ user.go
 ┃ ┣ mw
 ┃ ┃ ┗ jwt.go
 ┃ ┣ pack
 ┃ ┃ ┣ todo_list.go
 ┃ ┃ ┗ user.go
 ┃ ┣ router
 ┃ ┃ ┣ task3
 ┃ ┃ ┃ ┣ api.go
 ┃ ┃ ┃ ┗ middleware.go
 ┃ ┃ ┗ register.go
 ┃ ┣ service
 ┃ ┃ ┣ todo_list.go
 ┃ ┃ ┗ user_service.go
 ┃ ┗ utils
 ┃ ┃ ┗ password.go
 ┣ docs
 ┃ ┣ docs.go
 ┃ ┣ swagger.json
 ┃ ┗ swagger.yaml
 ┣ idl
 ┃ ┗ api.thrift
 ┣ script
 ┃ ┗ bootstrap.sh
 ┣ .gitignore
 ┣ .hz
 ┣ build.sh
 ┣ docker-compose.yml
 ┣ git.keep
 ┣ go.mod
 ┣ go.sum
 ┣ main.go
 ┣ README.md
 ┣ router.go
 ┗ router_gen.go
```
*   `biz`: 存放核心业务逻辑，进行了清晰的分层。
*   `docs`: 存放由 `swag` 自动生成的 API 文档文件。
*   `idl`: 存放 Thrift 接口定义文件。
*   `main.go`: 项目的入口文件。
*   `router.go`: 路由注册文件。
---

## 🛠️ 技术栈

*   **语言**: Go
*   **框架**: Hertz
*   **数据库**: MySQL
*   **ORM**: Gorm
*   **认证**: JWT
*   **API文档**: Swagger

---

## 🚀 快速开始

### 1. 环境准备

请确保你的电脑上已经安装了以下环境：
*   Go (版本 >= 1.18)
*   MySQL (版本 >= 5.7)

### 2. 克隆项目

```bash
git clone https://github.com/ShaddockNH3/west2-online-golang-2025-test
cd task3
```

### 3. 配置

```bash
docker-compose up -d
```

### 4. 安装依赖

```bash
go mod tidy
```

### 5. 启动项目

```bash
go run .
```
启动成功后，服务将在 `http://localhost:8888` 运行。

---

## 📖 API 文档

项目启动后，可以通过浏览器访问以下地址来查看和测试 API 接口：

[http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html)

---

## 📝 未来计划 (ToDo List)

为了让项目变得更强大、更可靠，计划在未来实现以下功能：

-   [ ] **重新设计统一的 JSON 返回格式**：
    *   目前不同接口的返回结构略有差异，计划设计一套标准化的 `code`, `msg`, `data` 格式，让前端处理起来更方便，也让错误信息更清晰！

-   [ ] **引入 Redis 缓存，提升高频接口性能**：
    *   对于像“查询用户信息”或“获取待办列表”这类可能会被频繁请求的接口，引入 Redis 进行数据缓存，可以大大减轻数据库的压力。

---

## 👤 作者

[ShaddockNH3](https://github.com/ShaddockNH3)