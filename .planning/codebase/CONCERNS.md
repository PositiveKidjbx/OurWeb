# Codebase Concerns

**Analysis Date:** 2026-04-29

## Summary

This repository is a static landing page implementation with minimal runtime logic and no build/test pipeline. The dominant risks are delivery-readiness gaps (placeholder content and legal/contact placeholders), maintainability constraints from hardcoded content, and quality risk from missing automated checks.

## Key Findings

### 1. Delivery-readiness gap: placeholder content is still in user-facing sections
- Issue: Multiple visible sections still contain demo placeholders and generic contact/legal links.
- Files: `index.html`
- Evidence:
  - Footer uses placeholder contact fields and generic email (`Address: Placeholder, Vietnam`, `hello@example.com`).
  - Legal links (`Terms`, `Privacy`) point to `#` rather than actual policy pages.
  - Several body paragraphs explicitly mark demo text for later replacement.
- Impact: Production publishing can expose incomplete or incorrect business information and legal compliance gaps.
- Fix approach: Replace placeholder text/links before release and gate release with a content-complete checklist.

### 2. Maintainability debt: content and assets are hardcoded in markup
- Issue: All section copy and image paths are embedded directly in `index.html`.
- Files: `index.html`, `docs/PRD.md`
- Evidence:
  - `index.html` contains direct `<img src="./assets/images/placeholders/...">` references and inline section copy.
  - `docs/PRD.md` specifies a mapping approach via `content/images.ts`, but that configuration layer is not present.
- Impact: Every content update requires HTML edits, increasing regression risk and making non-developer updates difficult.
- Fix approach: Introduce a content/config mapping module (per PRD direction), and generate section content from structured data.

### 3. Quality gap: no automated testing or verification tooling
- Issue: No test files, test runner config, or package manifest commands are present.
- Files: `index.html`, `main.js`, `styles.css` (only source files in root)
- Evidence:
  - Repository root contains only static source files plus docs/assets.
  - No `*.test.*`/`*.spec.*` files or common test config files detected.
- Impact: Responsive behavior, menu interaction, and future content updates can regress silently.
- Fix approach: Add a minimal frontend quality baseline (lint + formatting + smoke tests, e.g., Playwright interaction checks for nav/anchors and mobile menu).

### 4. External dependency fragility: runtime dependency on Google Fonts CDN
- Issue: Typography depends on remote Google Fonts at page load.
- Files: `index.html`
- Evidence:
  - `preconnect` and stylesheet links target `fonts.googleapis.com` and `fonts.gstatic.com`.
- Impact: Restricted-network users may see degraded typography or layout shifts; this can also affect predictable visual QA.
- Fix approach: Self-host required font files or define robust local/system fallback strategy with acceptance criteria.

### 5. Accessibility and interaction fragility in navigation toggle
- Issue: Mobile menu state is toggled via JS without broader interaction safeguards.
- Files: `main.js`, `index.html`
- Evidence:
  - Toggle only updates class and `aria-expanded`; no Escape-to-close or outside-click close behavior.
  - Anchor click handler closes nav globally for all hash links, coupling behavior to page structure.
- Impact: Usability and accessibility can degrade as navigation complexity grows.
- Fix approach: Add explicit keyboard/escape handling, focus management, and targeted event handling for nav context only.

## Unknowns / Assumptions

- Assumption: This repo is currently a delivery skeleton, not the final production deployable.
- Unknown: Deployment target and browser support matrix beyond PRD text are not implemented in repo config.
- Unknown: No Lighthouse reports or performance budgets are stored in repo, so current non-functional compliance is unverified.
- Unknown: No explicit security headers/CSP configuration is present in this repo; this may be handled by external hosting infrastructure.

---

*Concerns audit: 2026-04-29*
