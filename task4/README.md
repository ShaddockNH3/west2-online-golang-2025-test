# Task4 - 简易短视频后端

基于 CloudWeGo Hertz 框架实现的短视频平台后端服务。

## 项目说明

完成了 West2-Online 大作品的基础要求，实现了用户、视频、互动和社交四个核心模块。

## 主要功能

### 基础功能
- 用户注册登录、信息查询、头像上传
- 视频发布、列表查询、热门推荐、关键词搜索
- 点赞（视频/评论）、评论发布/查询/删除
- 关注/取消关注、关注列表、粉丝列表、好友列表

### 额外实现
- **MFA 双因素认证**: 登录时支持 TOTP 验证，提高账户安全性
- **Redis 缓存点赞**: 点赞操作先写入 Redis，提升响应速度
- **视频流推荐**: 实现获取最新视频列表的接口
- **以图搜图**: 基于文字描述搜索相关视频（暂未实现真正的图像识别）

### 技术实现
- JWT 身份认证
- bcrypt 密码加密
- Redis 热门排行榜和缓存
- Docker Compose 一键部署
- GORM + MySQL 数据持久化


## 项目结构

```
task4/
├── biz/                    # 业务逻辑层
│   ├── dal/db/            # 数据库操作
│   ├── handler/           # 请求处理器
│   ├── model/             # 数据模型
│   ├── mw/                # 中间件（JWT、Redis）
│   ├── router/            # 路由注册
│   └── service/           # 业务逻辑
├── idl/                    # Thrift 接口定义
├── pkg/                    # 工具和配置
│   ├── configs/           # 配置文件和初始化脚本
│   ├── constants/         # 常量定义
│   ├── errno/             # 错误码
│   └── data/              # 文件存储（头像、视频、封面）
├── docker-compose.yml      # Docker 编排
└── main.go                 # 入口文件
```

## 技术栈

- Go 1.23.5
- Hertz（HTTP 框架）
- GORM + MySQL（数据库）
- Redis（缓存）
- JWT（认证）
- Docker & Docker Compose


## 快速开始

### 使用 Docker（推荐）

```bash
# 1. 进入项目目录
cd task4

# 2. 启动所有服务
docker-compose up --build

# 3. 服务访问地址
# 应用: http://localhost:8080
# MySQL: localhost:9910
# Redis: localhost:9911
```

### 本地开发

```bash
# 1. 启动数据库
docker-compose up mysql redis -d

# 2. 修改配置
# 编辑 pkg/constants/constants.go
# 取消注释本地开发的配置项（127.0.0.1 相关）

# 3. 运行服务
go mod download
go run .
```

### 验证

```bash
curl http://localhost:8080/ping
# 返回: {"message":"pong"}
```


## API 接口

### 用户模块
- `POST /v1/user/register` - 用户注册
- `POST /v1/user/login` - 用户登录（支持 MFA code 参数）
- `GET /v1/user/info` - 获取用户信息
- `PUT /v1/user/avatar/upload` - 上传用户头像
- `GET /v1/auth/mfa/qrcode` - 获取 MFA 二维码（用于开启 MFA）
- `POST /v1/auth/mfa/bind` - 绑定 MFA
- `POST /v1/user/image/search` - 以图搜图（文字搜索）

### 视频模块
- `GET /v1/video/feed/` - 视频流（获取最新视频）
- `POST /v1/video/publish/` - 发布视频
- `GET /v1/video/list/` - 获取用户视频列表
- `GET /v1/video/popular/` - 获取热门视频
- `POST /v1/video/search/` - 搜索视频

### 互动模块
- `POST /v1/like/action` - 点赞/取消（支持 Redis 缓存）
- `GET /v1/like/list` - 点赞列表
- `POST /v1/comment/publish` - 发布评论
- `GET /v1/comment/list` - 评论列表
- `DELETE /v1/comment/delete` - 删除评论

### 社交模块
- `POST /v1/relation/action` - 关注/取消关注
- `GET /v1/following/list` - 关注列表
- `GET /v1/follower/list` - 粉丝列表
- `GET /v1/friends/list` - 好友列表

大部分接口需要 JWT 认证（Header: `Authorization: Bearer <token>`）

## 开发说明

### 配置修改

如果需要本地开发，主要修改 `pkg/constants/constants.go`：

```go
// Docker 部署使用（默认）
MySQLDefaultDSN = "gorm:gorm@tcp(mysql:3306)/gorm?..."
RedisAddr = "redis:6379"

// 本地开发使用（取消注释）
// MySQLDefaultDSN = "gorm:gorm@tcp(127.0.0.1:9910)/gorm?..."
// RedisAddr = "127.0.0.1:9911"
```

### 新功能说明

1. **MFA 验证**: 使用 TOTP 算法，密钥在 `constants.go` 中的 `MfaSecretKey`
2. **Redis 缓存点赞**: 点赞数据先存入 Redis（key: `:video_like_auth` 和 `:comment_like_auth`）
3. **视频流**: `/v1/video/feed/` 接口返回最新发布的视频列表
4. **以图搜图**: 目前实现为文字搜索，图像识别功能待完善

### 数据库表

- `users` - 用户信息
- `videos` - 视频信息
- `likes` - 点赞记录
- `comments` - 评论
- `follows` - 关注关系

初始化脚本在 `pkg/configs/sql/init.sql`

## 未来计划（Task5）

计划在下一个任务中进行以下优化和重构：

### 架构升级
- **微服务化**: 使用 Kitex 重构为微服务架构
- **服务注册与发现**: 接入 Etcd3 实现服务注册
- **架构重构**: 优化当前的代码组织结构，提升可维护性

### 功能增强
- **WebSocket**: 实现实时通信功能
- **聊天系统**: 
  - 实现安全的聊天功能
  - 优化聊天性能和并发处理
- **文件上传优化**:
  - 分片上传大文件
  - 分布式存储方案

### 性能优化
- 并发性能优化
- 数据库查询优化
- 缓存策略优化

### 工程化
- 引入 Git 工作流规范
- 使用 golangci-lint 等工具优化代码质量
- 完善单元测试覆盖

## 注意事项

1. Docker 部署时确保端口 8080、9910、9911 未被占用
2. 首次启动需要等待 MySQL 初始化完成
3. 上传的文件存储在 `pkg/data/` 目录
4. 生产环境请修改 JWT 密钥和数据库密码

## 参考

- [West2-Online 大作品文档](https://github.com/west2-online/learn-go/blob/main/docs/4-%E5%A4%A7%E4%BD%9C%E5%93%81.md)
- [CloudWeGo Hertz 文档](https://www.cloudwego.io/zh/docs/hertz/)
- [GORM 文档](https://gorm.io/zh_CN/docs/)