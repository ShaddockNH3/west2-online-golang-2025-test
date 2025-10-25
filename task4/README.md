# Task4 - 丐版抖音 📝

![Go Version](https://img.shields.io/badge/Go-1.23.5-blue.svg)
![Hertz](https://img.shields.io/badge/Hertz-v0.10.2-brightgreen.svg)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED.svg)
![License](https://img.shields.io/badge/License-MIT-lightgrey.svg)

基于 **CloudWeGo Hertz** 框架开发的轻量级短视频社交平台后端服务，实现了类抖音的核心功能。

## 📖 项目说明

- **项目要求**: [West2-Online 大作品文档](https://github.com/west2-online/learn-go/blob/main/docs/4-%E5%A4%A7%E4%BD%9C%E5%93%81.md)
- **接口文档**: [West2-Online API 文档](https://doc.west2.online/)
- **实现范围**: 完成了项目的最低要求，包含用户、视频、互动和社交四大核心模块

---

## ✨ 项目特色

### 🎯 核心功能

- **用户系统**: 用户注册、登录、信息查询、头像上传
- **视频模块**: 视频发布、列表查询、热门视频推荐、关键词搜索
- **互动功能**: 点赞（视频/评论）、评论发布/查询/删除
- **社交网络**: 关注/取消关注、关注列表、粉丝列表、好友列表

### � 技术亮点

- **JWT 认证**: 基于 `hertz-contrib/jwt` 实现的身份验证和授权机制
- **密码加密**: 使用 `bcrypt` 算法对用户密码进行安全加密
- **热门推荐**: 基于 Redis ZSet 实现的视频热度排行榜
- **文件存储**: 支持视频、封面、头像的本地文件上传与管理
- **容器化部署**: 完整的 Docker Compose 编排，一键启动所有服务
- **数据库设计**: 规范的表结构设计，支持软删除、索引优化

---

## �📂 项目结构

```
task4/
├── biz/                          # 业务逻辑层
│   ├── dal/                      # 数据访问层
│   │   ├── db/                   # 数据库操作
│   │   │   ├── init.go          # 数据库初始化 (GORM + MySQL)
│   │   │   ├── user.go          # 用户表操作
│   │   │   ├── video.go         # 视频表操作
│   │   │   ├── interact.go      # 互动表操作 (点赞/评论)
│   │   │   └── social.go        # 社交表操作 (关注关系)
│   ├── handler/                  # HTTP 请求处理器
│   │   ├── ping.go              # 健康检查
│   │   ├── user/                # 用户模块处理器
│   │   ├── video/               # 视频模块处理器
│   │   ├── interact/            # 互动模块处理器
│   │   └── social/              # 社交模块处理器
│   ├── model/                    # 数据模型 (Thrift 生成)
│   │   ├── common/              # 通用模型
│   │   ├── user/                # 用户相关模型
│   │   ├── video/               # 视频相关模型
│   │   ├── interact/            # 互动相关模型
│   │   └── social/              # 社交相关模型
│   ├── mw/                       # 中间件
│   │   ├── jwt/                 # JWT 认证中间件
│   │   └── redis/               # Redis 缓存中间件
│   ├── pack/                     # 数据封装层
│   ├── router/                   # 路由注册
│   └── service/                  # 业务逻辑服务层
│       ├── user_service/        # 用户业务逻辑
│       ├── video_service/       # 视频业务逻辑
│       ├── interact_service/    # 互动业务逻辑
│       └── social_service/      # 社交业务逻辑
├── idl/                          # Thrift IDL 接口定义
│   ├── common.thrift            # 通用数据结构
│   ├── user.thrift              # 用户接口定义
│   ├── video.thrift             # 视频接口定义
│   ├── interact.thrift          # 互动接口定义
│   └── social.thrift            # 社交接口定义
├── pkg/                          # 工具包
│   ├── configs/                 # 配置文件
│   │   ├── sql/init.sql        # 数据库初始化脚本
│   │   └── redis/redis.conf    # Redis 配置
│   ├── constants/               # 常量定义
│   ├── errno/                   # 错误码定义
│   ├── utils/                   # 工具函数 (密码加密等)
│   └── data/                    # 数据存储目录
│       ├── avatars/             # 用户头像
│       ├── covers/              # 视频封面
│       └── videos/              # 视频文件
├── script/                       # 脚本文件
│   └── bootstrap.sh             # 启动脚本
├── docker-compose.yml            # Docker Compose 配置
├── Dockerfile                    # 应用容器构建文件
├── go.mod                        # Go 模块依赖
├── main.go                       # 程序入口
├── router.go                     # 路由定义
└── router_gen.go                 # 路由生成代码
```

---

## 🛠️ 技术栈

### 后端框架与库

| 技术 | 版本 | 用途 |
|------|------|------|
| [Go](https://go.dev/) | 1.23.5 | 编程语言 |
| [Hertz](https://github.com/cloudwego/hertz) | 0.10.2 | 高性能 HTTP 框架 |
| [GORM](https://gorm.io/) | 1.31.0 | ORM 数据库操作 |
| [MySQL](https://www.mysql.com/) | 8.0 | 关系型数据库 |
| [Redis](https://redis.io/) | 7 | 缓存与排行榜 |
| [JWT](https://github.com/golang-jwt/jwt) | 4.5.2 | 身份认证 |
| [Thrift](https://thrift.apache.org/) | 0.13.0 | IDL 接口定义 |
| [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) | - | 密码加密 |
| [UUID](https://github.com/google/uuid) | 1.6.0 | 唯一标识符生成 |

### 基础设施

- **Docker**: 容器化部署
- **Docker Compose**: 多容器编排
- **健康检查**: MySQL 和 Redis 服务健康监控

---

## 🚀 快速开始

### 1. 环境准备

确保已安装以下工具：

- **Docker** (≥ 20.10)
- **Docker Compose** (≥ 2.0)
- **Go** (≥ 1.23.5) - 仅本地开发需要

### 2. 克隆项目

```bash
git clone https://github.com/ShaddockNH3/west2-online-golang-2025-test.git
cd west2-online-golang-2025-test/task4
```

### 3. 使用 Docker Compose 启动（推荐）

一键启动所有服务（MySQL + Redis + 应用）：

```bash
docker-compose up --build
```

容器启动后：
- **应用服务**: `http://localhost:8080`
- **MySQL**: `localhost:9910` (用户名: `gorm`, 密码: `gorm`)
- **Redis**: `localhost:9911` (密码: `shenmidazhi`)

### 4. 本地开发启动

如需本地开发调试：

```bash
# 先启动 MySQL 和 Redis
docker-compose up mysql redis -d

# 修改配置文件中的数据库连接地址
# pkg/constants/constants.go 中取消注释本地地址

# 安装依赖
go mod download

# 运行服务
go run .
```

本地服务启动后访问: `http://localhost:8888`

### 5. 验证服务

```bash
# 健康检查
curl http://localhost:8080/ping

# 预期返回
{"message":"pong"}
```

---

## 📡 API 接口

### 用户模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/v1/user/register` | 用户注册 | ❌ |
| POST | `/v1/user/login` | 用户登录 | ❌ |
| GET | `/v1/user/info` | 获取用户信息 | ✅ |
| PUT | `/v1/user/avatar/upload` | 上传用户头像 | ✅ |

### 视频模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/v1/video/publish/` | 发布视频 | ✅ |
| GET | `/v1/video/list/` | 获取用户视频列表 | ✅ |
| GET | `/v1/video/popular/` | 获取热门视频 | ✅ |
| POST | `/v1/video/search/` | 搜索视频 | ✅ |

### 互动模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/v1/like/action` | 点赞/取消点赞 | ✅ |
| GET | `/v1/like/list` | 获取点赞列表 | ✅ |
| POST | `/v1/comment/publish` | 发布评论 | ✅ |
| GET | `/v1/comment/list` | 获取评论列表 | ✅ |
| DELETE | `/v1/comment/delete` | 删除评论 | ✅ |

### 社交模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/v1/relation/action` | 关注/取消关注 | ✅ |
| GET | `/v1/following/list` | 获取关注列表 | ✅ |
| GET | `/v1/follower/list` | 获取粉丝列表 | ✅ |
| GET | `/v1/friends/list` | 获取好友列表 | ✅ |

### 静态资源

- **静态文件访问**: `http://localhost:8080/static/`
  - 头像: `/static/avatars/`
  - 封面: `/static/covers/`
  - 视频: `/static/videos/`

---

## 🗄️ 数据库设计

### 核心表结构

#### users - 用户表
- `id`: 用户ID (UUID)
- `username`: 用户名 (唯一)
- `password`: 密码 (bcrypt加密)
- `avatar_url`: 头像URL
- 支持软删除

#### videos - 视频表
- `id`: 视频ID (UUID)
- `user_id`: 作者ID
- `video_url`: 视频URL
- `cover_url`: 封面URL
- `title`: 标题
- `description`: 描述
- `visit_count`: 播放次数
- `like_count`: 点赞数
- `comment_count`: 评论数

#### likes - 点赞表
- `id`: 点赞ID (UUID)
- `user_id`: 点赞用户ID
- `likeable_id`: 被点赞对象ID
- `likeable_type`: 类型 (video/comment)

#### comments - 评论表
- `id`: 评论ID (UUID)
- `user_id`: 评论用户ID
- `video_id`: 视频ID
- `parent_id`: 父评论ID (楼中楼)
- `content`: 评论内容
- `like_count`: 点赞数
- `child_count`: 子评论数

#### follows - 关注关系表
- `id`: 关系ID (UUID)
- `follower_id`: 关注者ID
- `followed_id`: 被关注者ID
- 唯一约束: `(follower_id, followed_id)`

---

## 🔐 认证机制

### JWT Token

- **密钥**: `task4-secret-key`
- **访问令牌有效期**: 2小时
- **刷新令牌有效期**: 7天
- **Token传递方式**: HTTP Header
  ```
  Authorization: Bearer <token>
  ```

### 受保护的路由

除了用户注册和登录接口外，其他所有接口均需要携带有效的JWT Token。

---

## 📊 热门推荐算法

基于 **Redis ZSet** 实现的视频热度排行：

- **热度分数**: 综合播放量、点赞数、评论数等指标
- **实时更新**: 用户互动时动态更新热度分数
- **高效查询**: O(log N) 时间复杂度获取Top K热门视频
- **Redis Key**: `:popular_videos`

---

## 🐳 Docker 配置说明

### 服务编排

- **MySQL**: 
  - 端口映射: `9910:3306`
  - 自动执行初始化SQL脚本
  - 健康检查确保服务就绪

- **Redis**: 
  - 端口映射: `9911:6379`
  - 使用自定义配置文件
  - 数据持久化到本地目录

- **应用服务**: 
  - 端口映射: `8080:8888`
  - 依赖MySQL和Redis健康检查通过后启动
  - 多阶段构建优化镜像大小

---

## 🔧 配置文件

### 关键配置项

**pkg/constants/constants.go**:
```go
// JWT配置
JwtSecretKey = "task4-secret-key"
AccessTokenTimeout = 2 * time.Hour

// 数据库配置 (容器内)
MySQLDefaultDSN = "gorm:gorm@tcp(mysql:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

// Redis配置 (容器内)
RedisAddr = "redis:6379"
RedisPassword = "shenmidazhi"

// 服务地址
Host = "http://localhost:8080"
```

**本地开发时需修改为**:
```go
MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:9910)/gorm?..."
RedisAddr = "127.0.0.1:9911"
```

---

## 📝 开发说明

### 代码生成

项目使用 Thrift IDL 定义接口，通过 Hertz 工具生成代码：

```bash
# 安装 hz 工具
go install github.com/cloudwego/hertz/cmd/hz@latest

# 生成代码 (示例)
hz update -idl idl/user.thrift
```

### 添加新接口

1. 在 `idl/` 目录下定义 Thrift 接口
2. 使用 `hz` 工具生成代码
3. 在 `biz/service/` 实现业务逻辑
4. 在 `biz/handler/` 添加处理器
5. 在 `biz/router/` 注册路由

---

## 🔍 调试与测试

### 查看容器日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f task4_tiktok_app
docker-compose logs -f mysql
docker-compose logs -f redis
```

### 停止服务

```bash
# 停止并删除容器
docker-compose down

# 停止并删除容器及数据卷
docker-compose down -v
```

---

## 📌 注意事项

1. **首次启动**: MySQL初始化需要一定时间，应用会等待数据库就绪
2. **端口占用**: 确保本地端口 8080、9910、9911 未被占用
3. **文件上传**: 上传的文件存储在 `pkg/data/` 目录下
4. **密码安全**: 生产环境请修改默认的JWT密钥和数据库密码
5. **最大请求体**: 应用支持最大20MB的文件上传

---

## 🚧 已知限制

- 仅实现了项目的最低要求功能
- 文件存储使用本地文件系统，未集成对象存储服务
- 未实现视频转码和压缩功能
- 热门推荐算法较为简单，可优化为更复杂的推荐策略

---

## 📚 参考资料

- [CloudWeGo Hertz 文档](https://www.cloudwego.io/zh/docs/hertz/)
- [GORM 文档](https://gorm.io/zh_CN/docs/)
- [Docker Compose 文档](https://docs.docker.com/compose/)
- [West2-Online 学习文档](https://doc.west2.online/)

---

## 👤 作者

**ShaddockNH3**
- GitHub: [@ShaddockNH3](https://github.com/ShaddockNH3)
- 项目仓库: [west2-online-golang-2025-test](https://github.com/ShaddockNH3/west2-online-golang-2025-test)

---

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](../LICENSE) 文件

---

<div align="center">
  <sub>Built with ❤️ for West2-Online Golang 2025</sub>
</div>