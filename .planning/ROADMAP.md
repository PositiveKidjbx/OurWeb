# Roadmap: okia-vietnam-factory-delivery

## Overview

This roadmap takes the current brownfield landing page from functional implementation to a stable delivery baseline by tightening structure fidelity, replacement workflows, and quality verification. Work is split into three coarse phases so we can ship quickly while still enforcing acceptance standards.

## Phases

- [ ] **Phase 1: IA & Responsive Hardening** - Lock page structure, anchors, navigation behavior, and responsive stability.
- [ ] **Phase 2: Asset Replacement Workflow** - Standardize and validate placeholder-to-real-asset replacement mechanics.
- [ ] **Phase 3: Accessibility & Delivery Quality Gate** - Close a11y and performance gaps and prepare acceptance-ready output.

## Phase Details

### Phase 1: IA & Responsive Hardening
**Goal**: Ensure all required sections and interactions behave consistently across target viewport ranges.
**Depends on**: Nothing (first phase)
**Requirements**: STR-01, STR-02, STR-03, UIX-01, UIX-02, UIX-03
**Success Criteria** (what must be TRUE):
1. User can navigate to every major section through top navigation on desktop and mobile.
2. User sees no horizontal overflow on common mobile/tablet/desktop viewport widths.
3. User can use CTA paths to reach contact content from hero and footer routes.
**Plans**: 3 plans

Plans:
- [ ] 01-01: Audit current section/anchor contract and normalize IDs, nav targets, and CTA destinations.
- [ ] 01-02: Refine responsive breakpoints/layout rules and verify no overflow regressions.
- [ ] 01-03: Validate interaction flows (menu toggle, anchor close behavior) and polish edge cases.

### Phase 2: Asset Replacement Workflow
**Goal**: Make image and content replacement deterministic without structural edits.
**Depends on**: Phase 1
**Requirements**: AST-01, AST-02, AST-03
**Success Criteria** (what must be TRUE):
1. User can replace placeholders using documented naming/path conventions without layout breakage.
2. User can inspect a single mapping source for image references instead of editing scattered markup.
3. User sees consistent rendering behavior for hero, cards, and certification assets after replacement.
**Plans**: 2 plans

Plans:
- [ ] 02-01: Introduce or refine centralized asset mapping and wire page references through it.
- [ ] 02-02: Add replacement guide/checklist and verify with sample swap scenarios.

### Phase 3: Accessibility & Delivery Quality Gate
**Goal**: Reach acceptance-ready baseline for accessibility, console health, and performance.
**Depends on**: Phase 2
**Requirements**: A11Y-01, A11Y-02, A11Y-03, QLT-01, QLT-02
**Success Criteria** (what must be TRUE):
1. User can navigate primary controls with keyboard and understand image context via alt text.
2. User encounters no blocking console errors during common interactions.
3. User gets Lighthouse performance at or above target threshold in validation run.
**Plans**: 2 plans

Plans:
- [ ] 03-01: Apply accessibility fixes (alt semantics, focus order, control labels, contrast adjustments).
- [ ] 03-02: Run quality pass (console/performance), optimize critical assets, and finalize handoff notes.

## Progress

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. IA & Responsive Hardening | 0/3 | Not started | - |
| 2. Asset Replacement Workflow | 0/2 | Not started | - |
| 3. Accessibility & Delivery Quality Gate | 0/2 | Not started | - |
