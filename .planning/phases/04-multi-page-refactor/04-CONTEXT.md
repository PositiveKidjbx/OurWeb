# Phase 4: 多页面架构重构 - CONTEXT

**Gathered:** 2026-04-30  
**Status:** Ready for planning  
**Source:** $gsd-discuss-phase（人工讨论结论）

<domain>
## Phase Boundary

将当前单页（`index.html`）重构为多个页面文件，围绕已确定的导航：
- Home
- Sustainability
- Our Awards
- Career
- Contact Us

`OKIA Vietnam Factory` 的页面内容已并入 Home，不再作为独立导航项或独立页面。

本阶段不引入后端服务，仅做前端静态架构升级。

</domain>

<decisions>
## Implementation Decisions

### 页面拆分策略（锁定）
- 每个导航项使用独立 HTML 页面承载（至少：`index.html`、`sustainability.html`、`awards.html`、`career.html`、`contact.html`）。
- 导航从“页内锚点”逐步改为“跨页面链接”。

### 公共结构策略（锁定）
- 头部导航、页脚联系方式与基础样式保持统一。
- 继续沿用 `styles.css` 与 `content/images.js`，避免重复维护。

### 风险控制策略（锁定）
- 单页保留为 `index.html`（Home）作为兼容入口。
- 每个页面先完成基础内容落位，再逐页增强细节，避免一次性大改回归风险。
- 继续执行“每个小功能先验证再提交推送”。

### the agent's Discretion
- 是否引入简单构建工具（如模板 include）可在实现时评估；默认先不引入。
- 公共片段复用方式（复制基线 vs 小脚本生成）可按成本选择。

</decisions>

<canonical_refs>
## Canonical References

- `.planning/ROADMAP.md` - Phase 4 目标与计划
- `.planning/REQUIREMENTS.md` - 当前需求与追踪矩阵
- `docs/NAVIGATION_MAPPING.md` - 顶部导航定义与映射
- `docs/ASSET_REPLACEMENT_GUIDE.md` - 素材替换策略
- `content/images.js` - 图片映射源

</canonical_refs>

<specifics>
## Specific Ideas

- 先做信息架构层面的拆分，不急于改视觉。
- 第一轮先保证“每个页面可访问 + 导航一致 + 联系方式可达”。
- 第二轮再考虑细分每页独有内容与 SEO 元信息。

</specifics>

<deferred>
## Deferred Ideas

- CMS 接入
- 多语言
- 表单后端提交
- 深度 SEO 自动化

</deferred>

---

*Phase: 04-multi-page-refactor*  
*Context gathered: 2026-04-30 via discuss summary*

