# 顶部导航对照规范（多页面结构）

## 范围

当前项目顶部导航采用 5 个核心入口（不含 News）：

1. HOME
2. SUSTAINABILITY
3. OUR AWARDS
4. CAREER
5. CONTACT US

说明：`OKIA Vietnam Factory` 已并入 Home 页面内容，不再保留独立页面。

## 公共结构位置

- Header：`partials/header.html`
- Footer：`partials/footer.html`

页面通过 `data-include` 加载公共片段，并使用 `body[data-root]` 适配目录层级路径。

## 栏目定义与页面映射

### 1) HOME
- 导航文本：`Home`
- 页面：`index.html`
- 展示内容：首屏、工厂定位、制造能力概览（已承接原 Factory 页面核心内容）。

### 2) SUSTAINABILITY
- 导航文本：`Sustainability`
- 页面：`pages/sustainability.html`
- 展示内容：可持续实践与行动。

### 3) OUR AWARDS
- 导航文本：`Our Awards`
- 页面：`pages/awards.html`
- 展示内容：荣誉与认证展示。

### 4) CAREER
- 导航文本：`Career`
- 页面：`pages/career.html`
- 展示内容：招聘方向、雇主品牌、岗位入口。

### 5) CONTACT US
- 导航文本：`Contact Us`
- 页面：`pages/contact.html`
- 展示内容：电话、邮箱、地址等联系信息。
- 桌面端作为绿色主 CTA 展示。

## 页脚链接

- `Terms`：`pages/terms.html`
- `Privacy`：`pages/privacy.html`

## 验收要点

- 所有顶部导航链接命中有效页面。
- Header / Footer 不在每个页面重复手写。
- 新增页面时设置正确的 `body data-root`：
  - 根目录页面：`./`
  - `pages/` 下页面：`../`
