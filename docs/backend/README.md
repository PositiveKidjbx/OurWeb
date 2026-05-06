# Backend Overview

本文档用于说明当前后端项目的整体结构、启动入口、路由组织和主要模块职责。

## 1. Purpose

后端目前承担三类职责：

- 渲染公开页面入口
- 提供前端表单提交所需的公开 API
- 提供后台管理所需的管理员 API
- 初始化并访问 SQLite 数据库

当前后端既服务 HTML 模板和静态资源，也提供 JSON API。公开页面与管理接口都由同一个 Gin 服务承载。

## 2. Technology Stack

- Go
- Gin
- SQLite
- HTML templates
- Static assets served from `public`

## 3. Directory Structure

```text
backend/
  main.go
  internal/
    handler/
    service/
    repository/
    store/
  contact_api_test.go

templates/
public/
docs/backend/
```

关键目录说明：

- `backend/main.go`：应用启动入口，负责初始化依赖、注册路由、启动 HTTP 服务。
- `backend/internal/handler`：HTTP 层，处理请求解析、鉴权、响应状态码和响应格式。
- `backend/internal/service`：业务层，处理输入校验、分页规则和应用逻辑编排。
- `backend/internal/repository`：数据访问层，封装 SQLite 查询与写入。
- `backend/internal/store`：数据库初始化层，负责 SQLite 连接和 schema 初始化。
- `templates`：服务端渲染的 HTML 页面模板。
- `public`：CSS、JS、图片等静态资源。
- `docs/backend`：后端相关文档。

## 4. Application Entry Point

后端入口是 `backend/main.go`。

启动时会依次完成：

1. 读取运行配置
2. 初始化 SQLite 数据库
3. 创建 repository、service、handler
4. 注册 HTML template 和 static asset 路由
5. 注册公开 API 和 admin API
6. 启动 Gin HTTP server

依赖初始化方向是：

```text
store -> repository -> service -> handler -> routes
```

代码依赖方向是：

```text
handler -> service -> repository
```

`store` 只在启动阶段负责提供数据库连接，不直接参与 HTTP 请求处理。

## 5. Route Organization

当前路由分为四类：公开页面、静态资源、公开 API、后台 API。

```text
/                     public page
/contact              public page
/sustainability       public page
/awards               public page
/career               public page
/terms                public page
/privacy              public page
/news                 public page

/static/*             static assets

/api/contact          public API

/api/admin/login      admin login API
/api/admin/messages   admin-only API
```

组织原则：

- 公开页面路由通过 `registerPageRoute` 注册。
- 静态资源统一挂载在 `/static`。
- 公开业务接口统一放在 `/api` 下。
- 管理后台接口统一放在 `/api/admin` 下。
- 除 `/api/admin/login` 外，admin API 需要经过管理员鉴权中间件。

## 6. Backend Layering

后端采用简单分层结构，避免把 HTTP、业务逻辑和数据库查询混在一起。

```text
handler
  HTTP request/response
  auth middleware
  status code mapping

service
  business validation
  pagination defaults and limits
  application logic

repository
  SQL queries
  database read/write models

store
  SQLite connection
  schema initialization
```

这种分层让页面、API、数据库结构可以相对独立地演进。

## 7. Data Storage

当前后端使用 SQLite。

默认数据库路径：

```text
data/contact_messages.sqlite3
```

可通过环境变量覆盖：

```text
OW_DATABASE_PATH
```

当前核心表是 `contact_messages`，用于保存联系表单提交。表中包含 `status` 字段，用于后台管理提交信息的处理状态，例如新提交、已查看或后续扩展出的其它状态。

当前项目尚未上线，因此数据库初始化只负责创建当前 schema，暂未引入正式迁移机制。

## 8. Configuration

后端通过环境变量配置运行行为：

```text
PORT
OW_DATABASE_PATH
OW_FRONTEND_DIR
OW_ADMIN_USERNAME
OW_ADMIN_PASSWORD
OW_ADMIN_SESSION_SECRET
```

说明：

- `PORT`：HTTP 服务监听地址，默认 `:8080`。
- `OW_DATABASE_PATH`：SQLite 数据库路径。
- `OW_FRONTEND_DIR`：模板和静态资源所在的前端根目录，默认 `..`。
- `OW_ADMIN_USERNAME`：管理员登录用户名。
- `OW_ADMIN_PASSWORD`：管理员登录密码。
- `OW_ADMIN_SESSION_SECRET`：管理员 session cookie 签名密钥。

## 9. Admin Authentication

后台登录由 `/api/admin/login` 处理。登录成功后，后端会写入签名后的 HttpOnly cookie。

管理员数据接口挂在 `/api/admin` 下，并通过 admin auth middleware 保护。当前受保护的后台数据接口是 `/api/admin/messages`。

## 10. Testing

后端测试目前集中在：

```text
backend/contact_api_test.go
```

覆盖范围包括：

- 联系表单提交
- 无效提交拒绝
- 管理员登录
- 管理员鉴权失败
- 管理员分页读取提交信息

运行方式：

```bash
cd backend
GOCACHE=/tmp/ow-go-build-cache go test ./...
```

这里显式设置 `GOCACHE` 是为了避免某些开发环境中默认 Go build cache 目录不可写。

## 11. Related Docs

- `docs/backend/contact.md`：联系表单后端链路与接口细节说明。

后续如果后台功能继续扩展，可以新增：

- `docs/backend/admin.md`
- `docs/backend/api.md`
- `docs/backend/deployment.md`
