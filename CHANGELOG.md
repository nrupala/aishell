# Changelog — aishell

## 1.0.4 — one-shot behaviour lives in aishell now (2026-11-05, supersedes 1.0.3)

1.0.3's `--single-turn` worked on every build checked — but it is the same class of dependency that broke 1.0.1 and 1.0.2: a conversation flag upstream can rename or remove at will. Per the owner's direction, the one-shot behaviour now lives in aishell's own source, which upstream cannot remove.

- `llamaArgs` passes only the stable core (`-m`, `-c`, `-n`, `--temp`, `--log-disable`, `-p`) — unchanged in meaning since llama.cpp's earliest builds. No conversation or display flags at all.
- `runInference` manages the llama-cli process directly: it terminates the process the moment a complete response block has been captured (whatever the build's conversation behaviour — chat loop, extra generation, spinner), and an overall 10-minute timeout bounds a wedged run, the failure mode behind the 2026-10-30 396 MB spinner hang.
- The argument test now forbids the whole conversation/display flag class, and two new tests drive a fake conversation-mode llama-cli that would wait forever, proving aishell terminates it on completion and on timeout.

## 1.0.3 — the one-shot flag that survives (2026-11-05, certification proofs, supersedes 1.0.2)

The 1.0.2 fix was itself overtaken within hours: the newest llama.cpp builds (Homebrew on macOS, the MSVC dev build on Windows) removed the conversation toggle entirely, so `--no-conversation` is now rejected the same way `-no-cnv` was. Both certification proofs hit `error: invalid argument: --no-conversation` on every inference.

- Inference now passes **only** `--single-turn --no-display-prompt` as one-shot controls. `--single-turn` (generate one turn, then exit) exists with the same semantics in every build checked: the Aetheris proof build (bb4caa7, verified by direct invocation), ggml-org b8417, and the newest Homebrew/MSVC builds (verified from their own `--help` and by the re-run proofs).
- The argument-pinning test now also forbids both conversation toggles, so a future "helpful" flag addition fails the suite instead of failing users.

## 1.0.2 — cross-platform llama-cli flag fix (2026-11-05, certification proofs)

Found by the macOS CI proof (a real Apple Silicon runner, Homebrew llama.cpp): newer llama.cpp builds removed the short `-no-cnv` alias, so every inference failed with `error: invalid argument: -no-cnv` on a current `llama-cli`. The Linux proof build (bb4caa7) still accepted the short form, which is why this only surfaced on macOS.

- Inference now uses the long flag forms only (`--no-conversation --single-turn --no-display-prompt`), which both the older proof build and the newest builds accept. A unit test pins the argument list so the short aliases cannot creep back.
- The banner, OS detection, and session flow on macOS were unaffected — only the inference invocation was rejected; failure was graceful (`[Inference Error]`, clean exit).
- README: the Quick start now leads with the prebuilt-binary route (an independent cold-user walk found the build-only ordering left binary users without an end-to-end path), and the stale test count is removed.

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
