# 顶部导航对照规范（多页面结构）

## 范围

本项目顶部主导航采用多页面链接，不再以首页锚点作为主要栏目入口。

公共导航统一维护在：

- `partials/header.html`

公共页脚统一维护在：

- `partials/footer.html`

页面通过 `data-include` 加载公共片段，并使用 `body[data-root]` 适配不同目录层级的相对路径。

## 栏目定义与页面映射

### 1) HOME
- 导航文本：`Home`
- 页面：`index.html`
- 展示内容：首页首屏与企业概览。

### 2) SUSTAINABILITY
- 导航文本：`Sustainability`
- 页面：`pages/sustainability.html`
- 展示内容：可持续实践内容。

### 3) OKIA VIETNAM FACTORY
- 导航文本：`OKIA Vietnam Factory`
- 页面：`pages/factory.html`
- 展示内容：工厂定位、制造能力、经验与交付能力。

### 4) OUR AWARDS
- 导航文本：`Our Awards`
- 页面：`pages/awards.html`
- 展示内容：荣誉与认证展示。

### 5) CAREER
- 导航文本：`Career`
- 页面：`pages/career.html`
- 展示内容：招聘方向、雇主品牌、岗位申请入口。

### 6) NEWS
- 导航文本：`News`
- 页面：`pages/news/index.html`
- 展示内容：新闻列表。

### 7) CONTACT US
- 导航文本：`Contact Us`
- 页面：`pages/contact.html`
- 展示内容：电话、邮箱、地址等联系信息。
- 桌面端作为 CTA 展示；移动端作为菜单项展示。

## 页脚链接

- `Terms`：`pages/terms.html`
- `Privacy`：`pages/privacy.html`

## 验收要点

- 所有顶部导航链接必须命中有效页面。
- Header / Footer 不应在每个页面重复手写。
- 新增页面时需要设置正确的 `body data-root`：
  - 根目录页面：`./`
  - `pages/` 下页面：`../`
  - `pages/news/` 下页面：`../../`
- 由于公共片段通过 `fetch()` 加载，建议通过本地静态服务或正式 Web 服务访问，不建议直接用 `file://` 打开。
