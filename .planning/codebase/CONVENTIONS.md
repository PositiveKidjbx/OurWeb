# Coding Conventions

**Analysis Date:** 2026-04-29

## Summary
This repository is a static frontend project with conventions defined by direct HTML/CSS/vanilla-JS usage in `index.html`, `styles.css`, and `main.js`. No formatter, linter, or typed language configuration files are detected, so current conventions are lightweight and file-local.

## Naming Patterns

**Files:**
- Use lowercase root-level filenames for core assets: `index.html`, `main.js`, `styles.css`.
- Use kebab-case for placeholder asset filenames under `assets/images/placeholders/`, e.g. `sustainability-01.svg`, `cert-06.svg`.

**Functions:**
- Use anonymous arrow callbacks for DOM event handlers in `main.js`.
- Prefer verb-first handler intent in usage context, e.g. click handlers attached via `addEventListener("click", ...)` in `main.js`.

**Variables:**
- Use `camelCase` for JavaScript variables: `menuToggle`, `isOpen` in `main.js`.
- Use `const` for DOM references and local values in `main.js`.

**Types:**
- Not applicable; no TypeScript files are present (`rg --files` shows no `.ts`/`.tsx`).

## Code Style

**Formatting:**
- Tool used: Not detected (no `.prettierrc*`, `biome.json`, or equivalent config files found).
- Key settings: Not applicable.
- Current style in `main.js` and `styles.css` indicates 2-space indentation and semicolon-terminated JavaScript statements.

**Linting:**
- Tool used: Not detected (no `.eslintrc*` or `eslint.config.*` found).
- Key rules: Not applicable.

## Import Organization

**Order:**
1. Not applicable (no ES module `import` statements; script loaded via `<script src="./main.js"></script>` in `index.html`)
2. Not applicable
3. Not applicable

**Path Aliases:**
- Not detected.

## Error Handling

**Patterns:**
- No explicit defensive null checks around `document.getElementById(...)` results in `main.js`; script assumes `#menuToggle` and `#nav` exist in `index.html`.
- No `try/catch` or custom error boundary pattern detected.

## Logging

**Framework:** console not used in current code.

**Patterns:**
- No runtime logging statements detected in `main.js`.

## Comments

**When to Comment:**
- No inline comments detected in `index.html`, `main.js`, or `styles.css`; style favors self-descriptive class and variable names.

**JSDoc/TSDoc:**
- Not used.

## Function Design

**Size:**
- JavaScript logic is compact and event-driven in `main.js` (single-purpose callbacks for menu toggle and anchor click behavior).

**Parameters:**
- Callback parameters are omitted when unused (e.g., `() => { ... }` in `main.js`).

**Return Values:**
- Event callbacks rely on side effects (`classList.toggle`, `setAttribute`) and do not return explicit values.

## Module Design

**Exports:**
- Not applicable; no module export/import system in use.

**Barrel Files:**
- Not applicable.

## Key Findings
- `main.js` centralizes interactive behavior and follows consistent `camelCase` + `const` usage.
- `styles.css` uses CSS custom properties in `:root` (`--bg`, `--brand`, etc.) as the styling token pattern.
- Structural naming in `index.html` and `styles.css` is class-based and semantic (`.hero`, `.section-light`, `.footer-grid`).
- Quality tooling (lint/format/test) is not configured in repository root.

## Unknowns/Assumptions
- Assumption: formatting/linting may be enforced outside this repository because no local config is present.
- Unknown: whether this project is intentionally no-build/no-tooling or tooling files were omitted.
- Unknown: browser support baseline for CSS features such as `color-mix()` in `styles.css`.

---

*Convention analysis: 2026-04-29*
