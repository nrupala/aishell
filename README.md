# aishell — an AI-native shell that never runs anything behind your back

Copyright © 2026 2351485 ALBERTA INC.

Type what you want in plain language. aishell turns it into a clean command for your operating system, shows it to you, and runs it **only if you say yes**.

Inference is fully local: aishell drives a `llama-cli` (llama.cpp) process with a small local model on your own machine. No account, no cloud, no telemetry. The confirmation gate is the product: destructive patterns are blocked outright, file-writing commands are rerouted to your editor, and nothing executes silently — ever.

## How it works

1. You type an intent — `find every log file over 100 MB`.
2. A local Qwen2.5-Coder 3B model (via `llama-cli`) translates it into one command block for your shell (PowerShell, bash, or zsh).
3. aishell extracts the command and screens it:
   - **Destructive blacklist** (`rm -rf`, `mkfs`, `dd if=`, `Format-Volume`, …) → blocked, full stop.
   - **Write/mutation operators** (`>`, `>>`, `Out-File`, `Set-Content`, `tee`, …) → rerouted: the file opens in your native editor instead.
4. The surviving command is printed, and you decide: `Execute command on host system? [y/N]`
5. On macOS it then runs inside `sandbox-exec` with file writes denied. On Linux it runs in a `systemd-run` scope (process grouping under your user manager — no extra OS sandbox in this version; an earlier design claimed systemd sandbox properties that scope units in fact reject, so every approved command failed — that claim was removed, see CHANGELOG). On Windows it runs in PowerShell after your confirmation. On Linux and Windows, the gate is the screening plus your confirmation.
6. On exit, aishell offers to archive a session log **locally** — only if you explicitly answer yes. Command-frequency counts (the "watchdog", which suggests aliases for commands you run often) are kept locally too. Nothing ever leaves your machine; see [PRIVACY.md](PRIVACY.md).

Built-in extras: `help`, `aliases` (pre-baked shortcuts like `sysinfo`, `netstat`, `diskuse`), `exit`.

## Requirements

- Go 1.21 or newer (to build)
- `llama-cli` from [llama.cpp](https://github.com/ggml-org/llama.cpp), in your PATH (or in the project folder — see `main.go`)
- The model file, in `./models/` with exactly this name:
  `qwen2.5-coder-3b-instruct-q4_k_m.gguf`
  (Qwen2.5-Coder-3B-Instruct, Q4_K_M — the setup scripts fetch it from Hugging Face under the right name.)

Zero Go dependencies. One module, six source files, a test suite.

## Quick start

```bash
# Linux / macOS
./setup.sh     # checks llama-cli, downloads the model
make           # builds Windows + Linux + macOS binaries
./ai-shell-linux        # (or ./ai-shell-mac)
go test ./...           # 10 tests: parsing, sandbox screening, streaming, prompts
```

```powershell
# Windows
.\setup.bat
.\build.bat
.\ai-shell.exe
```

Prebuilt binaries for Windows, Linux (amd64), and macOS (Apple Silicon) are attached to each release — built from the tagged source, with SHA-256 sums in the release notes.

## Security model — read this before you trust it

aishell is a *gate*, not a guarantee. The blacklist and mutation screening are plain string checks in `execution.go` — readable, auditable, and bypassable by a sufficiently creative generated command. That is exactly why the final step is a human one: **you read the command and you press y.** Do not remove the confirmation prompt. If a command looks wrong, the correct move is `N` — aishell logs the decline and moves on.

Known limits, stated plainly:

- Windows execution has no OS sandbox — confirmation is the only gate there.
- The screening lists are finite. A novel destructive phrasing can pass the screen; it cannot pass you unless you approve it.

## Privacy

Runs fully local. No telemetry, no analytics, no account. Session history and watchdog counts are stored only on your machine, and session history is archived only on your explicit yes at exit. Full statement: [PRIVACY.md](PRIVACY.md).

## Licence

Apache License 2.0 — see [LICENSE](LICENSE) and [NOTICE](NOTICE).

aishell is free. If it saved you time, you can support its maintainer via the links in the Funding panel — voluntary, never required.

---
A 2351485 ALBERTA INC. tool, built under AIMLDS.
