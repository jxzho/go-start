# 学习路径（含实践）

这是一条面向前端工程师的 Go 学习路径，按阶段推进，强调“可落地实践”。

## 0. 准备（1–2 天）
目标：能写、能跑、能测。

- 安装 Go 与 VS Code（Go 插件）
- 了解 `go mod`、`go run`、`go test`

实践：
- 写一个 `hello` 程序，支持参数（`os.Args`），输出 `Hello, {name}`

## 1. 语法与基础（1–2 周）
重点：
- 变量/常量、切片/数组、map、结构体
- 方法、接口、包与模块
- 错误处理与 `defer`
- 指针基础（会用即可）

实践：
1) CLI 词频统计：读取文本，统计词频 Top N
2) 待办清单工具：添加/删除/完成/列表

## 2. 标准库与工程化（1–2 周）
重点：
- `net/http`、`encoding/json`、`context`
- 文件读写、时间处理
- `go fmt`、`go test`、`go vet`

实践：
- 做一个简易 HTTP JSON API：
  - `GET /todos`
  - `POST /todos`
  - `DELETE /todos/{id}`
- 数据先存内存，后续可扩展为文件或数据库

## 3. 并发与性能思维（1 周）
重点：
- goroutine 与 channel
- worker pool、`sync.Mutex`、`WaitGroup`

实践：
1) 并发爬取器：给一组 URL，并发请求并汇总状态码/耗时
2) 并发图片下载器：限制并发数并统计下载结果

## 4. Web 后端实战（2–4 周）
重点：
- 路由、中间件、日志、配置
- 数据库（SQLite 或 Postgres）
- 统一响应结构与错误处理

实践：
- 小型博客 API：
  - `POST /posts`
  - `GET /posts`
  - `GET /posts/{id}`
- 增加分页、排序、搜索参数

## 5. 进阶与部署（持续）
重点：
- 依赖注入（轻量）
- gRPC / OpenAPI
- Docker 与简单 CI

实践：
- 用 OpenAPI 生成接口文档
- Docker 化并部署到小服务器

## 学习节奏建议
- 每天 1–2 小时，坚持 6–8 周可达到“可用后端能力”
- 每个阶段都写 README 与 API 文档，训练工程表达
