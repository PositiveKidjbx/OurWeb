# 集成点

**分析日期：** 2026-04-29

## 总结

当前项目集成点较少，主要是浏览器内原生能力与少量外部静态资源服务。

## 外部集成

- **Google Fonts**：用于加载 `Manrope`、`Barlow` 字体。
- **静态资源路径**：`assets/images/placeholders/` 作为图片输入边界。

## 内部集成关系

- `index.html` 通过 `<link>` 与 `<script>` 集成 `styles.css` 和 `main.js`。
- `main.js` 依赖 `index.html` 中的 DOM 结构（如 `#menuToggle`、`#nav`）。
- `styles.css` 依赖 HTML class/id 钩子。

## 风险与建议

- DOM 选择器与结构强耦合，后续改结构需同步更新脚本。
- 字体依赖外网，建议提供本地字体回退策略。
- 建议后续引入集中化内容/资源映射，降低硬编码分散风险。

---
*集成分析：2026-04-29*
