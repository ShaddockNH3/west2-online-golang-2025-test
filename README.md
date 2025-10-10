# west2-online-golang-2025-test

task1和task2在其他仓库，写的很抽象，未来有机会再移过来~~（搬史）~~

task3目前已经实现——

- 用户登录
- 用户注册
- 用户查询
- 用户删除
- 增添备忘录
- 删除备忘录
- 更新备忘录
- 查询备忘录

bonus（已完成）

- 考虑了数据库交互的安全性（把密码用了哈希加密）
- 使用了三层架构设计（router -> handler -> service -> dal）

todo

- 自动生成接口文档
- 重新设计json
- 使用redis