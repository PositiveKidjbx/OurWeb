# External Integrations

**Analysis Date:** 2026-04-29

## Summary
The current codebase has minimal external integration: browser-loaded Google Fonts and no active API, database, authentication, monitoring, webhook, or backend service integration.

## Key Findings
- `index.html` preconnects and loads Google Fonts from `fonts.googleapis.com` and `fonts.gstatic.com`.
- All business content is static in `index.html`; contact details are plain text in footer with no submit workflow.
- `docs/PRD.md` explicitly marks backend form submission as out of scope in this phase.
- No SDK imports, API clients, or external service credentials are present in `main.js`/`styles.css`/`index.html`.

## Unknowns/Assumptions
- Future integration points (form submit endpoint, CMS, analytics) are implied by product roadmap language but not implemented.
- Deployment platform is unspecified; this document assumes static hosting due to current artifact shape.
- Secret management process is unknown because no environment config files are present.

## APIs & External Services

**Typography/CDN:**
- Google Fonts - Web font delivery for UI typography.
  - SDK/Client: Browser `<link>` stylesheet include in `index.html`.
  - Auth: None.

**Reference-only (non-runtime):**
- `https://www.okia.com/okia-vietnam-factory` appears in `docs/PRD.md` as a design reference URL, not a runtime API dependency.

## Data Storage

**Databases:**
- Not detected.
  - Connection: Not applicable.
  - Client: Not applicable.

**File Storage:**
- Local repository filesystem only (`assets/images/placeholders/*`).

**Caching:**
- None application-specific detected.

## Authentication & Identity

**Auth Provider:**
- None.
  - Implementation: Not applicable.

## Monitoring & Observability

**Error Tracking:**
- None detected.

**Logs:**
- Browser console/runtime defaults only; no structured logging framework detected.

## CI/CD & Deployment

**Hosting:**
- Not explicitly configured; static site hosting is the compatible model for current files.

**CI Pipeline:**
- Not detected (no workflow/config files observed in current repository root tree).

## Environment Configuration

**Required env vars:**
- None detected for current implementation.

**Secrets location:**
- Not applicable for current implementation.

## Webhooks & Callbacks

**Incoming:**
- None.

**Outgoing:**
- None.

---

*Integration audit: 2026-04-29*
