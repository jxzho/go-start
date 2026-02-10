# 实战项目：书签管理器 API（Bookmark API）

核心价值：贴近前端日常需求，易扩展（标签、搜索、收藏）。

## 需求范围
- 书签：标题、URL、描述、标签、创建时间
- 基础接口：
  - `POST /bookmarks`
  - `GET /bookmarks`
  - `GET /bookmarks/{id}`
  - `DELETE /bookmarks/{id}`
- 支持分页、排序、标签过滤

## 里程碑拆解

### M1：基础 API
- 完成路由与 JSON 结构
- 内存存储 + 统一响应格式

### M2：工程化
- `go test` 覆盖 handler 与 service
- 增加日志与错误码

### M3：数据库
- 使用 SQLite/Postgres
- 增加迁移脚本与 Repo 层

### M4：文档与部署
- OpenAPI 文档
- Dockerfile + README
