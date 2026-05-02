# Backend Internal 逻辑说明

## 1. 目标

`backend/internal` 目录负责处理联系表单提交的完整后端链路，包括：

- 接收前端 `POST /api/contact` 请求
- 校验用户提交内容
- 将合法数据写入 SQLite
- 对外返回安全、简洁的响应
- 在服务端保留内部错误信息

当前实现采用三层结构：

- `handler`：HTTP 接口层
- `service`：业务校验与编排层
- `repository`：数据库访问层
- `store`：数据库初始化层

## 2. 总体请求流程

一次完整的提交流程如下：

1. 客户端向 `POST /api/contact` 发送 JSON 请求
2. `ContactHandler.Create` 解析 JSON
3. `ContactService.Submit` 校验字段
4. `ContactRepository.Create` 将数据写入 SQLite
5. Handler 返回最终响应给客户端

## 3. 目录职责

### 3.1 `internal/handler`

文件：

- `internal/handler/contact_handler.go`

职责：

- 接收 HTTP 请求
- 解析 JSON 请求体
- 调用 service 层
- 控制对外返回的状态码和响应体
- 处理校验错误和内部错误的分流

当前对外行为：

- JSON 非法，返回 `400`
- 字段校验失败，返回 `400`
- 服务端内部错误，返回 `500`
- 成功提交后，返回 `201`

### 3.2 `internal/service`

文件：

- `internal/service/contact_service.go`

职责：

- 定义请求数据结构 `ContactInput`
- 做字段级业务校验
- 组织提交逻辑
- 调用 repository 层完成数据库写入

当前校验规则：

- `name` 必填
- `email` 必填
- `email` 必须是合法邮箱格式
- `message` 必填
- `company` 和 `phone` 允许为空

错误类型：

- `ValidationError` 用于表示用户输入不合法

### 3.3 `internal/repository`

文件：

- `internal/repository/contact_repository.go`

职责：

- 只负责数据库写入
- 不处理 HTTP，不处理业务校验
- 将数据插入 `contact_messages` 表

当前写入字段：

- `name`
- `email`
- `company`
- `phone`
- `message`

返回值：

- 当前返回插入后的 `LastInsertId`
- 该值只用于内部链路，不应直接暴露给客户端

### 3.4 `internal/store`

文件：

- `internal/store/sqlite.go`

职责：

- 初始化 SQLite 数据库连接
- 在数据库文件不存在时创建目录和数据库文件
- 建立表结构
- 提供可复用的数据库初始化入口

当前表结构：

- `contact_messages`
- 字段：
  - `id`
  - `name`
  - `email`
  - `company`
  - `phone`
  - `message`
  - `created_at`

## 4. 接口说明

### `POST /api/contact`

请求头：

```http
Content-Type: application/json
```
请求体示例：

```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "company": "OKIA",
  "phone": "123456789",
  "message": "Hello from contact form"
}
```
成功响应：

```json
{
  "message": "contact message saved"
}
```
失败响应示例：

```json
{
  "error": "email is invalid"
}
```
内部错误响应示例：

```json
{
  "error": "failed to save contact message"
}
```
## 5. 数据库说明
默认数据库路径：

```text
data/contact_messages.sqlite3
```
可通过环境变量覆盖：

```text
OW_DATABASE_PATH
```
例如：

```bash
OW_DATABASE_PATH=/tmp/contact.sqlite3
```
初始化规则：

- 如果数据库路径不存在，自动创建
- 如果父目录不存在，自动创建父目录
- 初始化时自动创建 contact_messages 表
- 重复启动不会重复建表，因为使用了 CREATE TABLE IF NOT EXISTS

## 6. 错误处理原则
当前实现遵循以下原则：

- 用户输入错误只返回简洁的校验信息
- 内部错误不直接暴露底层实现细节
- 真实错误会记录在服务端日志中
- handler 负责对外安全输出，service 负责业务校验，repository 负责底层操作

## 7. 依赖关系
依赖方向为：

```text
handler -> service -> repository -> store
```
解释：

- handler 只知道 service
- service 只知道 repository
- repository 只知道数据库连接
- store 只负责初始化数据库

## 8. 未来扩展建议
后续可以继续扩展：

- 增加提交频率限制
- 增加管理员查看列表接口
- 增加邮件通知
- 增加验证码或反垃圾提交机制
- 增加提交来源字段，例如 source_page
- 增加单元测试和集成测试覆盖
