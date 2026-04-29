# okia-vietnam-factory-delivery

## What This Is

This project is a responsive, single-page corporate website delivery modeled on the OKIA Vietnam Factory reference page. It is designed for B2B visitors, partners, and brand researchers to quickly understand company positioning, core capabilities, sustainability practices, and contact channels. The current implementation emphasizes maintainability so business teams can replace assets and copy with minimal engineering effort.

## Core Value

Deliver a clear, trustworthy corporate narrative in one page while keeping content and image replacement low-cost and low-risk.

## Requirements

### Validated

- [x] Single-page structure with key corporate sections is implemented in existing codebase.
- [x] Responsive layout and mobile navigation behavior are implemented.
- [x] Placeholder image strategy and standardized asset paths are already in use.

### Active

- [ ] Improve accessibility compliance consistency (keyboard flow, labels, contrast checks).
- [ ] Harden performance baseline to reliably meet Lighthouse performance target in normal delivery conditions.
- [ ] Add maintainable content/image mapping layer so replacements can be done without touching page structure.

### Out of Scope

- CMS backend - excluded in current PRD scope to keep delivery fast.
- Multi-language switching - explicitly deferred to later iteration.
- Form backend submission - not required for current static delivery phase.
- Deep SEO optimization - only baseline meta support expected in this cycle.

## Context

- Existing implementation is a static frontend stack (`index.html`, `styles.css`, `main.js`) with no build system or backend.
- PRD (`docs/PRD.md`) defines this as a fast landing page delivery with placeholder assets and later business-side replacement.
- Codebase mapping artifacts under `.planning/codebase/` provide architecture, stack, conventions, testing, and concerns baselines.
- Current maturity is closer to "functional draft delivered" than "production-hardened release."

## Constraints

- **Tech stack**: Static HTML/CSS/JS without framework or backend - must preserve low complexity and easy handoff.
- **Compatibility**: Modern browsers (Chrome/Edge/Safari latest two major versions) - required by PRD.
- **Asset dependency**: Final enterprise assets/copy provided by client - delivery quality depends on incoming material quality.
- **Scope/timeline**: Corporate landing page only - avoid expanding into CMS or multi-page platform work.

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Keep static architecture for v1 | Aligns with fast delivery and easy replacement requirement from PRD | - Pending |
| Treat existing implemented sections as validated baseline | Brownfield repo already contains expected page modules | - Pending |
| Use coarse-grained roadmap phases | Project is small and focused; broad phases reduce overhead | - Pending |

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `$gsd-transition`):
1. Requirements invalidated? -> Move to Out of Scope with reason
2. Requirements validated? -> Move to Validated with phase reference
3. New requirements emerged? -> Add to Active
4. Decisions to log? -> Add to Key Decisions
5. "What This Is" still accurate? -> Update if drifted

**After each milestone** (via `$gsd-complete-milestone`):
1. Full review of all sections
2. Core Value check - still the right priority?
3. Audit Out of Scope - reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-04-29 after initialization*
