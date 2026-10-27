# AI-Shell — Build Chat History

Durable, append-only record of work performed on the AI-Native Cross-Platform Shell
wrapper (`D:\aishell`). New entries are added at the end; past entries are never
rewritten.

---

## Entry 001 — Initial consolidation & green build

- **Date:** 2026-08-15
- **Role:** Agent (OpenCode) under ASF-methodology policy (verified-before-done).

### Context
The project directory held multiple independent `main.go` variants produced as
iterations by a prior agent in `By_Gemini/` (`main.go`, `main_2.go`, `main_3.go`,
`main(1).go`) plus a root `main.go`. Most of them were structurally broken
(markdown/code interleaved, broken escape sequences, mangled ` thinking` /
` response` delimiter tokens, non-streaming inference, missing sandbox wrappers,
wrong log paths). `go vet`/`go test` failed across `./...` because the `By_Gemini`
`.go` files were treated as packages.

### Decisions
1. **Single consolidated program.** Fold all feature sets into one `package main`
   split across five files for structural hygiene, rather than many single-file
   main programs:
   - `main.go` — entry point, interactive loop, OS detection, aliases, help, gate
   - `inference.go` — system-prompt builder + streaming `llama-cli` runner
   - `parser.go` — marker-aware markdown-fence payload extraction
   - `execution.go` — sandboxed subprocess spawn + security gate
   - `storage.go` — watchdog + platform-specific session logging
2. **Archive the superseded `By_Gemini` variants.** Renamed `.go` -> `.txt` so Go
   stops compiling them but the content remains readable as reference. (All converge
   into the new `main.go`.)
3. **Real verified model URL.** The blueprints referenced a placeholder
   `https://huggingface.co`. Verified against the Hugging Face API that the correct
   asset is `bartowski/Qwen2.5-Coder-3B-Instruct-GGUF` ->
   `Qwen2.5-Coder-3B-Instruct-Q4_K_M.gguf`.
4. **Build scripts build the package** (`.`) so multi-file compilation works, across
   Windows/Linux/macOS. Added a `test` target to the Makefile.
5. **Add tests** (`main_test.go`) covering parser, streaming, and security logic.

### Implementation highlights
- **Streaming robustness:** `runInference` uses `cmd.StdoutPipe()` (not buffered
  `cmd.Run()`). `streamInference` reads rune-by-rune and uses a `flushableLen`
  holdback window so the ` thinking` / ` response` delimiter tokens are captured
  even when split across chunk edges; reasoning chain prints in real time (dim
  styling). Fence extraction is a pure-`strings` line parser (no regex), tolerant of
  `cmd` / `powershell` / `bash` / `zsh` / `shell` / bare fences.
- **Pure OS isolation (no Docker):**
  - Linux: `systemd-run --user --scope -p PrivateDevices=true -p
    ProtectHome=read-only -p ProtectSystem=strict /bin/bash -c <payload>`
  - macOS: `sandbox-exec -p "(version 1) (allow default) (deny file-write* (subpath "/"))"
    /bin/zsh -c <payload>`
  - Windows: `powershell.exe -NoProfile -Command <payload>`
  - `os/exec` delegates `Stdin`/`Stdout`/`Stderr` to the subprocess so `vim`/`less`
    stay interactive.
- **Descriptor/logging sanity:**
  - Windows: local ISO-8601 timestamps (`2006-01-02T15:04:05.0000000-07:00`),
    written to `%APPDATA%\ai_shell`, history file `AppEventLog_<datetime>.json`.
  - POSIX: true integer Unix epoch `int64`, written to `~/.local/share/ai_shell`
    (honors `XDG_DATA_HOME`), history file `syslog_<epoch>.json`.
- **Security gate:** destructive-pattern blacklist (hard block) + write-mutation
  operators (reroute to native viewer/editor). Prompt-injection unsafe outputs yield
  an empty payload -> gate refuses.
- **Watchdog:** counts executed command patterns across runs via JSON metrics,
  recommends alias promotion at count == 3.

### Verification (evidence)
- `gofmt -l` on root sources: clean.
- `go vet ./...`: clean.
- `go test ./...`: PASS (parser, streaming, security, prompt-schema, editor-fallback,
  marker-split-across-chunks). One test expectation corrected after confirming the
  flash/token-holdback length (3, not 4) is the correct outcome.
- Cross-builds all succeed (`CGO_ENABLED=0`):
  - `windows/amd64` -> `ai-shell.exe` ~2.48 MB
  - `linux/amd64`   -> `ai-shell-linux` ~2.37 MB
  - `darwin/arm64`  -> `ai-shell-mac` ~2.23 MB

### Outcome
Project compiles, tests pass, and all three target architectures build. Ready for
runtime bring-up (`setup.sh`/`setup.bat` then `make`/`build.bat` then `ai-shell`).

---

## Entry 002 — Setup scripts & Makefile promotion

- **Date:** 2026-08-15

### Context
Follow-up to Entry 001: add the dependency-checker setup assets referenced by the
blueprints and finalize the build tooling.

### Decisions / actions
1. Created `setup.sh` (Unix/Linux/macOS) and `setup.bat` (Windows) that:
   - verify `llama-cli` presence (PATH then project root),
   - create `models/`,
   - download `Qwen2.5-Coder-3B-Instruct-Q4_K_M.gguf` (curl/wget on Unix;
     PowerShell `Invoke-WebRequest` on Windows) only if absent.
2. Promoted `Makefile.txt` -> `Makefile` with package targets and a `test` target.

### Verification
`go test ./...` still PASS. Makefile contains `go build .`, `go vet ./...`,
`go test ./...` targets.

### Outcome
Asset set complete: `main.go`, `inference.go`, `parser.go`, `execution.go`,
`storage.go`, `main_test.go`, `go.mod`, `Makefile`, `build.bat`, `setup.sh`,
`setup.bat`.

---

## Status

- [x] 001 — consolidation, streaming, sandboxing, logging, tests, cross-builds
- [x] 002 — setup scripts + Makefile promotion
- [ ] 003 — runtime bring-up on target host (model present, `llama-cli` in PATH)