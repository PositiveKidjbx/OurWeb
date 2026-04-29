# Technology Stack

**Analysis Date:** 2026-04-29

## Summary
This repository is a static frontend landing page with plain HTML, CSS, and vanilla JavaScript. Runtime execution is browser-native with no detected package manager, build pipeline, or server runtime.

## Key Findings
- Core app assets are `index.html`, `styles.css`, and `main.js` with direct browser loading (`<link rel="stylesheet" href="./styles.css" />`, `<script src="./main.js"></script>`) in `index.html`.
- Styling is hand-authored CSS using custom properties and responsive media queries in `styles.css`.
- Interaction logic is minimal DOM scripting for mobile menu toggle and anchor navigation behavior in `main.js`.
- Google Fonts is used as an external dependency via `fonts.googleapis.com` and `fonts.gstatic.com` in `index.html`.
- Project planning document `docs/PRD.md` defines scope as single-page delivery and explicitly excludes backend form submission for current phase.

## Unknowns/Assumptions
- Version pinning for browser compatibility is not encoded in config files; compatibility is inferred from `docs/PRD.md` browser target statements.
- No CI/CD or deployment manifest is present; hosting assumptions are based on static-site structure.
- No lockfile or package metadata exists; dependency tracking is implicit rather than tool-managed.

## Languages

**Primary:**
- HTML5 - Page structure and section composition in `index.html`.
- CSS3 - Theming, layout, and responsiveness in `styles.css`.
- JavaScript (ES6+) - DOM behavior and menu interactions in `main.js`.

**Secondary:**
- Markdown - Product requirements and implementation scope in `docs/PRD.md`.

## Runtime

**Environment:**
- Web browser runtime (Chrome/Edge/Safari class targets stated in `docs/PRD.md`).

**Package Manager:**
- Not detected.
- Lockfile: missing.

## Frameworks

**Core:**
- Not detected (no React/Vue/Angular or backend framework files).

**Testing:**
- Not detected (no test runner config or test files).

**Build/Dev:**
- Not detected (no bundler, transpiler, or task runner config).

## Key Dependencies

**Critical:**
- Google Fonts CSS endpoint (`https://fonts.googleapis.com/css2?...`) - Supplies `Manrope` and `Barlow` font families for typography in `index.html`.

**Infrastructure:**
- Static local assets under `assets/images/placeholders/` - Provide demo visuals referenced by `index.html` sections.

## Configuration

**Environment:**
- No `.env`-driven runtime integration detected.
- Content and section structure are configured directly in `index.html` and described by `docs/PRD.md`.

**Build:**
- No build config files detected (`package.json`, `vite.config.*`, `webpack.*`, `tsconfig.json`, `go.mod`, `pyproject.toml` absent at repo root).

## Platform Requirements

**Development:**
- Any static file-capable environment and modern browser for preview.
- No local service dependencies detected.

**Production:**
- Static hosting target (CDN, object storage static website, or basic web server) is compatible with current artifact layout (`index.html` + static assets).

---

*Stack analysis: 2026-04-29*
