# Changelog — aishell

## 1.0.1 — end-to-end proof fixes (2026-10-30, Aetheris proof)

The first full end-to-end run (real Qwen2.5-Coder-3B model + llama-cli staged on the Aetheris box, six scripted sessions) found three defects that unit tests and static review had missed. All three are fixed here; the proof was then re-run green.

- **Inference hung forever with current llama.cpp.** `llama-cli` now defaults to conversation mode — it generated, then waited for a next chat turn, and aishell blocked on the capture pipe indefinitely (one session produced 396 MB of spinner output). `runInference` now passes `-no-cnv --single-turn --no-display-prompt`, so llama-cli generates once and exits (observed: ~9 s, clean exit).
- **Prompt echo could have poisoned parsing.** Current llama-cli echoes the prompt — including the system prompt's empty command template — into its output. `extractCommandBlock` now collects all fenced blocks, filters the template placeholder, and takes the last real block (two new tests cover this).
- **Linux execution never ran.** The Linux sandbox invocation passed `PrivateDevices` / `ProtectHome` / `ProtectSystem` to a `systemd-run --scope` unit, which rejects all of them ("Unknown assignment", systemd 255) — every approved command failed. The service-unit alternative was tested and rejected on evidence (status 218 with PrivateDevices in a user manager; ProtectHome did not actually block a home write). Linux now runs the confirmed command in a bare `systemd-run --scope`, and the docs no longer claim a Linux OS sandbox: on Linux the screening + confirmation gate is the control.
- Session evidence (transcripts): benign approve → executes and prints output; benign decline → nothing runs; destructive intent → model refuses, no dangerous command produced; write intent → rerouted to editor, no file created; builtins + alias interception work; session history archives only on explicit yes.

## Unreleased — free-release prep (release-prep branch)

Prepared 2026-10-26 by Wright (2351485 ALBERTA INC.). Not published.

- Storefront README rewritten: clean markdown, verified build/test instructions, honest security model with stated limits.
- Added PRIVACY.md (source-verified), NOTICE, LICENSE (Apache 2.0, proposed — owner's licence word pending), .github/FUNDING.yml, SECURITY.md, CI workflow (vet + test + all-target builds).
- Fix: `setup.sh` saved the model under the Hugging Face filename, which does not match the path `main.go` loads (`models/qwen2.5-coder-3b-instruct-q4_k_m.gguf`) on case-sensitive filesystems — first run on Linux/macOS would have failed to find the model. Setup now saves under the loaded name.
- Fix: session history was archived on an empty answer at the `[y/N]` archive prompt (default meant archive). History now archives only on an explicit `y`/`yes`; verified by running the built binary both ways.
- Verified at prep: `go vet` clean, `go test` 10/10 pass, binaries built from this source for Windows (amd64), Linux (amd64), macOS (arm64).

## 1.0.0 — 2026-08-15 (original build, custody import)

- Original complete build (By_Gemini build record): Go sources, tests, Makefile/build scripts, prebuilt binaries for Windows/Linux/macOS, design docs. Imported as-is into this repository's `main` branch as the custody record — see the import commits.
