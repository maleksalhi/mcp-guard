# PROJECT_STATE — mcp-guard (v0.1.0)

## 1. Locked Scope (Zero Additions Allowed)
- **Language:** Go 1.22+ (single binary, zero CGO, `slog` -> `os.Stderr` only).
- **Core Proxy:** Transparent `stdio` JSON-RPC 2.0 wrapper (`tools/call`, `resources/read`), fail-closed, hard 10MB payload limit.
- **Policy Engine:** `policy.yaml` + `filepath.EvalSymlinks` path sanitization + rate limiter + session anti-exfiltration + response secret redaction.
- **Tier 2 Additions:** Tool Definition Pinning (SHA-256 on `tools/list`) + Tamper-Evident Async SQLite Audit Log (hash-chained via buffered channel) + `./mcp-guard logs` + `/metrics`.
- **HITL Gate:** Local HTTP (`localhost:8787`), `context.WithTimeout`, `notifications/cancelled`.
- **Proof:** `go test -race`, `FuzzParseJSONRPC`, `<1ms` overhead benchmark, `tests/attacks/`, `THREAT_MODEL.md`, `ROADMAP.md` (Tier 1 parked here), 20s demo GIF.

## 2. Progress Tracker
- [ ] **Day 1 (5%):** Repo init, `main.go` skeleton, `Makefile`, GitHub Actions CI (`go test -race`).
- [ ] **Days 2–3 (20%):** Transparent `stdio` proxy + 10MB limit + graceful shutdown.
- [ ] **Days 4–5 (40%):** JSON-RPC parser + `policy.yaml` + path/symlink sanitization + Tool Pinning.
- [ ] **Days 6–7 (55%):** Stateful rules (rate limit, anti-exfiltration) + response redaction.
- [ ] **Days 8–9 (70%):** HITL HTTP gate + timeouts + `notifications/cancelled`.
- [ ] **Days 10–11 (80%):** Async hash-chained SQLite audit log + CLI `logs` + `/metrics`.
- [ ] **Day 12 (90%):** Race tests, Fuzzing, `<1ms` benchmarks.
- [ ] **Day 13 (95%):** `tests/attacks/` + Prompt Injection Demo GIF.
- [ ] **Day 14 (100%):** `THREAT_MODEL.md`, `ROADMAP.md`, README, Tag `v0.1.0`.

## 3. Current Status & Architectural Decisions
- **Current Phase:** Day 1 (Setup & CI).
- **Key Decisions:** No `fmt.Println` to stdout; `modernc.org/sqlite` (pure Go); fail-closed on parse/policy error.
- **Next Action:** Push Day 1 to GitHub and verify green CI.
