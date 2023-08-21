web框架：Gin 处理API路由和参数
与数据库交互：数据库驱动mysql，database/sql包
组织代码：MVC架构

- Model层：定义表示数据库实体的Go结构体
- Controller层：处理业务逻辑，处理请求逻辑,与Model、Repository层交互。直接返回JSON格式数据给前端
- Repository层：基于Model操作数据库,使用 database/sql 直接执行 SQL 查询,并手动转换结果到 Model 对象。
  主要功能：

1. 查看简历列表

- 后端提供获取所有简历列表的API
- 前端调用API并展示列表,支持分页

2. 查看简历详情

- 后端提供通过ID获取简历详情的API
- 前端调用API并展示简历详情

3. 创建简历

- 后端提供创建简历的API
- 前端表单收集创建信息,调用API创建简历

4. 编辑简历

- 后端提供编辑简历的API
- 前端表单收集编辑信息,调用API编辑简历

5. 删除简历

- 后端提供删除简历的API
- 前端调用API删除简历

6. 导出简历为PDF

- 后端提供导出简历为PDF文件的API
- 前端调用API并在浏览器中展示或下载PDF文件

7. 搜索过滤

- 前端实现搜索和过滤功能
- 后端提供支持搜索和过滤的API

9. 简历模板

- 后端可预置存储几种简历模板
- 前端可选择模板并生成简历

├── main.go
├── routes
│ └── api.go
├── middleware
│ └── auth.go
│ └── logging.go
├── config
├── models
│ └── user.go
├── repositories  
│ └── user.go
├── services
│ └── user.go
├── controllers
│ └── user.go
├── docs
│ └── docs.go
│ └── swagger.json
├── tests
└── boot
└── db.go

- main.go: 程序入口
- routes: 路由配置
- middleware: 中间件
- config: 配置
- models: 数据模型
- repositories: 数据库操作
- services: 业务逻辑
- controllers: 控制器
- docs: API 文档生成
- tests: 测试
- boot: 启动初始化任务
