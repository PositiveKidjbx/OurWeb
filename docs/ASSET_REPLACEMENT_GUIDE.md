# 图片替换指南（Phase 2）

## 目标

本项目已接入集中素材映射层。后续替换图片时，只改一个文件：`content/images.js`。

## 操作入口

- 映射文件：`content/images.js`
- 页面引用：`index.html` 中的 `data-image-key`
- 运行时注入：`main.js`

## 替换步骤

1. 将新图片放入 `assets/images/placeholders/`（或你约定的新目录）。
2. 打开 `content/images.js`，按 key 修改路径。
3. 不要改 `index.html` 的结构与 `data-image-key`，除非新增模块。
4. 刷新页面并执行下方验证清单。

## Key 对照表

- `heroBg` -> 首屏背景图
- `aboutImage` -> About 区图片
- `sustainability01` ~ `sustainability05` -> Sustainability 五张卡片图
- `cert01` ~ `cert06` -> Certification 六个 logo 图

## 命名建议

- 小写 + 中划线（例如 `sustainability-06.svg`）
- 同模块按序号递增（`01`、`02`）
- 统一扩展名风格（建议同一模块保持一致）

## 注意事项

- 若某个 key 缺失映射，页面会回退到 HTML 内的 `src` 默认值。
- 若路径错误，浏览器会出现图片加载失败；需在控制台/网络面板检查 404。
- 替换时优先保持与原图接近的宽高比，减少版面波动。
