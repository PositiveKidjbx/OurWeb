# 需求文档：okia-vietnam-factory-delivery

**定义日期：** 2026-04-29  
**核心价值：** 在单页中清晰传达企业信息，同时保持素材替换成本低、风险可控。

## v1 需求

### 结构与内容

- [ ] **STR-01**：页面包含约定模块（Header、Hero、About/Mission、Why Choose Us、Sustainability、Certification、Footer）。
- [ ] **STR-02**：导航锚点与区块映射正确，桌面与移动端均可清晰使用。
- [ ] **STR-03**：联系入口在首屏 CTA 与页脚均可触达。

### 占位素材机制

- [ ] **AST-01**：占位图目录与命名符合 PRD 规范。
- [ ] **AST-02**：替换图片不需要改 HTML 结构。
- [ ] **AST-03**：卡片图与认证图在常见尺寸波动下不破坏布局。

### 交互与响应式

- [ ] **UIX-01**：移动端菜单可正常开关，并在锚点点击后自动关闭。
- [ ] **UIX-02**：桌面/平板/手机断点下无横向滚动条。
- [ ] **UIX-03**：锚点跳转与 CTA 行为一致、可预期。

### 可访问性与质量

- [ ] **A11Y-01**：内容图片具有有意义的 `alt` 文本。
- [ ] **A11Y-02**：主导航与 CTA 支持键盘访问与激活。
- [ ] **A11Y-03**：文字与背景具备基础可读性对比。
- [ ] **QLT-01**：Lighthouse Performance >= 75（开发验证）。
- [ ] **QLT-02**：常规交互下无阻断级控制台报错。

## v2 需求（延期）

- **V2-01**：多语言切换框架。
- **V2-02**：CMS 内容管理能力。
- **V2-03**：表单后端提交与通知。
- **V2-04**：深度 SEO 策略。

## 范围外

| 功能 | 原因 |
|------|------|
| 后端服务层 | 本期为静态页面交付 |
| 多页站点架构 | 当前只做单页官网 |
| 高级埋点体系 | 非本期验收要求 |
| 设计系统迁移 | 对单页项目投入产出不高 |

## 可追踪矩阵

| 需求 | 阶段 | 状态 |
|------|------|------|
| STR-01 | Phase 1 | Pending |
| STR-02 | Phase 1 | Pending |
| STR-03 | Phase 1 | Pending |
| AST-01 | Phase 2 | Pending |
| AST-02 | Phase 2 | Pending |
| AST-03 | Phase 2 | Pending |
| UIX-01 | Phase 1 | Pending |
| UIX-02 | Phase 1 | Pending |
| UIX-03 | Phase 1 | Pending |
| A11Y-01 | Phase 3 | Pending |
| A11Y-02 | Phase 3 | Pending |
| A11Y-03 | Phase 3 | Pending |
| QLT-01 | Phase 3 | Pending |
| QLT-02 | Phase 3 | Pending |

**覆盖统计：**
- v1 需求总数：14
- 已映射阶段：14
- 未映射：0

---
*定义日期：2026-04-29*  
*最后更新：2026-04-29（初始化）*
