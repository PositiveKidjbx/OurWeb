# Codebase Structure

**Analysis Date:** 2026-04-29

## Summary

The repository is a flat static-site layout with three top-level runtime files and two supporting directories (`assets/`, `docs/`). There is no `src/` split, so placement decisions are made by file role: markup in root HTML, behavior in root JS, and styles in root CSS.

## Directory Layout

```text
okia-vietnam-factory-delivery/
├── .planning/                    # Planning artifacts and generated codebase maps
│   └── codebase/                 # Architecture/structure documentation output
├── assets/                       # Static image resources
│   └── images/placeholders/      # Placeholder SVG assets used by page sections
├── docs/                         # Product/documentation notes
│   └── PRD.md                    # Product requirements document
├── index.html                    # Main page entry and semantic content
├── main.js                       # Client-side interaction logic
└── styles.css                    # Global styles, layout, responsive rules
```

## Directory Purposes

**`.planning/codebase`:**
- Purpose: Stores generated mapping documents for downstream planning/execution flows.
- Contains: Markdown analysis files.
- Key files: `.planning/codebase/ARCHITECTURE.md`, `.planning/codebase/STRUCTURE.md`.

**`assets/images/placeholders`:**
- Purpose: Holds visual placeholder assets referenced by page sections.
- Contains: SVG files for hero, about, sustainability, and certifications.
- Key files: `assets/images/placeholders/hero.svg`, `assets/images/placeholders/about.svg`.

**`docs`:**
- Purpose: Project-level written requirements.
- Contains: Markdown documentation.
- Key files: `docs/PRD.md`.

## Key File Locations

**Entry Points:**
- `index.html`: Browser document entry, section structure, stylesheet/script linking.
- `main.js`: Runtime behavior entry for navigation interactions.

**Configuration:**
- Not detected as dedicated config files (no package/runtime config manifests present at repo root).

**Core Logic:**
- `main.js`: Menu toggle state and anchor-click close behavior.
- `styles.css`: Layout grids, responsive breakpoints, design tokens.

**Testing:**
- Not detected (no test files or test config files in current repository tree).

## Naming Conventions

**Files:**
- Root runtime files use simple lowercase names: `index.html`, `main.js`, `styles.css`.
- Placeholder assets use kebab-case with numeric suffixes: `sustainability-01.svg`, `cert-06.svg`.

**Directories:**
- Content directories use lowercase words, with nesting by asset type: `assets/images/placeholders/`.

## Where to Add New Code

**New Feature:**
- Primary markup/content: `index.html` (add section and anchor IDs).
- Behavior: `main.js` (append DOM event logic, keep selectors ID/class based).
- Styling: `styles.css` (add component classes and media-query rules).

**New Component/Module:**
- Implementation: Add HTML block in `index.html` and corresponding style block in `styles.css`.

**Utilities:**
- Shared helpers: Not currently separated; place small shared JS helpers in `main.js` unless size growth justifies creating `assets/js/` or similar.

## Special Directories

**`.planning`:**
- Purpose: Planning system artifacts.
- Generated: Yes.
- Committed: Yes (present in repository).

**`assets/images/placeholders`:**
- Purpose: Demo visual inventory.
- Generated: No (static assets).
- Committed: Yes.

## Key Findings

- Structure is intentionally simple and optimized for direct static delivery.
- Single-file-per-concern pattern is clear: content (`index.html`), style (`styles.css`), behavior (`main.js`).
- No module or build folder indicates no transpilation/bundling step in current structure.

## Unknowns/Assumptions

- Assumes future scale remains small enough for root-level file organization.
- Assumes deployment pipeline (if any) is external to this repository because no CI/build files are present.
- Assumes `.planning/` is part of team workflow and should remain versioned.

---

*Structure analysis: 2026-04-29*
