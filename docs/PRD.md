# PRD - 企业官网落地页（参考 OKIA Vietnam Factory）

## 1. 文档信息
- 项目名称：`okia-vietnam-factory-delivery`
- PRD版本：`v0.1`
- 创建日期：`2026-04-29`
- 目标：交付一个与参考站点版式风格相近的企业展示网页，后续可由业务方自行替换内部图片与文案。

## 2. 背景与目标
外部需求要求快速搭建一个企业展示页面，参考页面：
- https://www.okia.com/okia-vietnam-factory

本期聚焦：
- 复刻信息架构与页面节奏（非逐像素抄写）
- 搭建可维护的前端项目骨架
- 所有图片使用占位资源，确保后续低成本替换

## 3. 范围定义
### 3.1 In Scope（本期范围）
- 单页官网（Landing Page）
- PC + 移动端响应式
- 页面模块：
  - 顶部导航
  - Hero 首屏
  - 品牌/使命区
  - Why Choose Us（核心能力）
  - Sustainability（可持续）卡片区
  - Certification（资质）区
  - Footer 联系方式区
- 图片资产占位机制（统一命名、统一尺寸建议）

### 3.2 Out of Scope（非本期范围）
- CMS 后台
- 多语言切换（可预留）
- 表单后端提交能力
- SEO 深度优化（仅保留基础 meta）

## 4. 目标用户
- B2B 潜在客户
- 合作伙伴与供应链审核方
- 求职者/品牌调研访问者

## 5. 业务目标（可量化）
- 页面首屏 3 秒内传达企业定位
- 关键能力点（经验、创新、认证、服务）清晰可扫读
- 联系入口在首屏和底部均可触达
- 后续替换图片无需改动 HTML 结构

## 6. 页面信息架构
1. Header
- Logo
- 导航锚点：About / Technology / Sustainability / Certification / Contact
- CTA 按钮：Contact Us

2. Hero
- 主标题（企业名或工厂名）
- 副标题（品牌主张）
- 双按钮（Contact / Learn More）
- 背景大图（占位）

3. About / Mission
- 简短品牌描述
- 使命文案
- 配图（占位）

4. Why Choose Us
- 四象限能力点：Experience / Innovation / Certified / Professional Service
- 每项支持标题 + 1 段说明

5. Sustainability
- 5 张卡片：
  - Eco Construction
  - Solar Energy
  - Net Zero Emissions
  - Rain Collection System
  - CSR
- 卡片包含图、标题、摘要文案

6. Certification
- 认证 logo 墙（占位图）
- 可配置徽章数量

7. Footer
- 公司地址
- 电话/邮箱
- 社交链接入口
- 法务链接（Terms / Privacy）

## 7. 图片占位与替换规范
### 7.1 目录规范
- `assets/images/placeholders/hero.jpg`
- `assets/images/placeholders/about.jpg`
- `assets/images/placeholders/why-choose-us.jpg`
- `assets/images/placeholders/sustainability-01.jpg` ~ `05.jpg`
- `assets/images/placeholders/cert-01.png` ~ `06.png`

### 7.2 命名规则
- 全小写 + 中划线
- 同模块图片按两位序号递增（`01`、`02`）

### 7.3 尺寸建议
- Hero 背景：`1920 x 900`
- 内容配图：`1200 x 800`
- 卡片图：`800 x 600`
- 认证 Logo：透明底，建议高度统一 `120px`

### 7.4 替换机制
- 在配置文件中维护图片路径映射（如 `content/images.ts`）
- 替换时仅更新资源文件与映射，不改结构代码

## 8. 功能需求
### 8.1 基础交互
- 顶部导航点击平滑滚动到对应区块
- CTA 点击跳转到 Contact 区
- 移动端折叠菜单

### 8.2 响应式
- 断点建议：
  - `>=1200` 桌面
  - `768-1199` 平板
  - `<768` 手机
- 小屏时卡片自动单列/双列

### 8.3 可访问性
- 所有图片必须有 `alt`
- 键盘可访问导航
- 文本与背景对比度满足基础可读性

## 9. 非功能需求
- 兼容现代浏览器（Chrome/Edge/Safari 最新两个主版本）
- 首屏资源优化（压缩图片、延迟加载非首屏）
- 代码结构清晰，便于二次开发

## 10. 验收标准（UAT）
- 页面模块齐全，结构与参考站点节奏一致
- 任意替换一组占位图后，布局不破坏
- 手机端无横向滚动条
- Lighthouse Performance >= 75（开发环境可接受）
- 无阻断级控制台报错

## 11. 风险与依赖
- 风险：真实业务图片尺寸不统一导致视觉错位
- 缓解：提供图片比例约束与 `object-fit` 策略
- 依赖：甲方提供最终图片与企业文案

## 12. 里程碑建议
1. M1：项目脚手架 + 静态结构（0.5 天）
2. M2：样式与响应式 + 占位图接入（1 天）
3. M3：内容替换、联调、验收（0.5~1 天）

## 13. 交付物
- 前端项目源码（单页）
- 图片占位目录与替换说明
- 本 PRD 文档
