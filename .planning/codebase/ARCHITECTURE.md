# Architecture

**Analysis Date:** 2026-04-29

## Summary

This project uses a static, single-page architecture with three runtime layers: HTML structure in `index.html`, presentation in `styles.css`, and minimal behavior in `main.js`. The browser loads the page directly with no server-side application layer, no build pipeline, and no package-managed module system.

## Pattern Overview

**Overall:** Static SPA-like landing page (no framework routing)

**Key Characteristics:**
- Content sections are anchored and navigated by fragment links in `index.html`.
- Styling is centralized with CSS custom properties and responsive media queries in `styles.css`.
- Interactivity is limited to mobile navigation toggle behavior in `main.js`.

## Layers

**Markup Layer:**
- Purpose: Defines all semantic sections, navigation, and content placeholders.
- Location: `index.html`
- Contains: Header/nav, hero, content sections, footer, static asset references.
- Depends on: `styles.css`, `main.js`, image assets under `assets/images/placeholders/`.
- Used by: Browser render engine and JS DOM selectors.

**Presentation Layer:**
- Purpose: Controls design tokens, layout grids, typography, and responsive breakpoints.
- Location: `styles.css`
- Contains: `:root` variables, component classes (`.card`, `.tile`, `.logo-wall`), media queries.
- Depends on: Class/id hooks in `index.html`.
- Used by: All visible UI on desktop/mobile.

**Behavior Layer:**
- Purpose: Manages menu open/close state and accessibility state updates.
- Location: `main.js`
- Contains: Event handlers for `#menuToggle` and in-page anchor clicks.
- Depends on: `#menuToggle`, `#nav`, and anchor links in `index.html`.
- Used by: Mobile/tablet navigation interactions.

## Data Flow

**Navigation Interaction Flow:**
1. User clicks menu button `#menuToggle` in `index.html`.
2. `main.js` toggles class `open` on `#nav` and updates `aria-expanded`.
3. CSS rule `.nav.open` in `styles.css` controls visible/hidden state.

**In-page Link Flow:**
1. User clicks any link matching `a[href^="#"]`.
2. `main.js` removes `open` from `#nav` and resets `aria-expanded` to `false`.
3. Browser handles native anchor scroll to target section IDs in `index.html`.

**State Management:**
- UI state is DOM-class based (`open` on `#nav`) with no external store.

## Key Abstractions

**Section-as-Module Pattern:**
- Purpose: Each business block is a self-contained section with its own class structure.
- Examples: `#about`, `#technology`, `#sustainability`, `#certification` in `index.html`.
- Pattern: Semantic section boundaries + reusable utility classes from `styles.css`.

**Tokenized Styling Pattern:**
- Purpose: Keep visual consistency through shared color tokens.
- Examples: `:root` variables (`--bg`, `--brand`, `--line`) in `styles.css`.
- Pattern: CSS custom properties consumed across components.

## Entry Points

**Document Entry:**
- Location: `index.html`
- Triggers: Direct browser load.
- Responsibilities: Bootstraps UI, links CSS/JS, provides all page content and anchors.

**Script Entry:**
- Location: `main.js`
- Triggers: Script tag at end of `index.html` body.
- Responsibilities: Binds event listeners for mobile nav behavior.

## Error Handling

**Strategy:** Implicit browser-level only.

**Patterns:**
- No explicit try/catch in `main.js`.
- No fallback guards if expected DOM nodes are missing.

## Cross-Cutting Concerns

**Logging:** Not implemented (no logging calls in `main.js`).
**Validation:** Not implemented (static content only).
**Authentication:** Not applicable (no user/session flow).

## Key Findings

- Architecture is intentionally minimal and suitable for static hosting; no backend or API coupling detected from `index.html` and `main.js`.
- Responsive behavior is CSS-first, with JS used only for toggling mobile nav visibility (`styles.css` + `main.js`).
- External dependency surface is limited to Google Fonts referenced in `index.html`.

## Unknowns/Assumptions

- Assumes deployment target serves static files with preserved relative paths used in `index.html`.
- Assumes no additional runtime logic exists outside `main.js` since no other scripts are present in repository root.
- Assumes content management/replacement is manual because no CMS integration is referenced in code.

---

*Architecture analysis: 2026-04-29*
