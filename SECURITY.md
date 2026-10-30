# Security policy — aishell

Copyright © 2026 2351485 ALBERTA INC.

## The model, in one paragraph

aishell generates shell commands with a local model and executes them only after the user confirms at a `[y/N]` prompt. Destructive patterns are blacklisted, write/mutation commands are rerouted to a text editor, and confirmed commands run inside an OS sandbox on macOS (`sandbox-exec`, file writes denied); on Linux they run in a `systemd-run` scope and on Windows directly in PowerShell — on those two, the gate is the confirmation itself. The screening is string-based and deliberately simple — the human confirmation is the final gate, and the README says so plainly. A bypass of the screening that still requires user confirmation is a hardening opportunity; a way to execute *without* confirmation is a critical vulnerability.

## Reporting a vulnerability

Please report privately, not as a public issue:

- Preferred: GitHub private vulnerability reporting on this repository (Security → Report a vulnerability), once enabled on the public repo.
- Until then / alternatively: contact the maintainer through the AIMLDS site contact page at https://aimlds.org.

Tell us the version, your OS, the exact input and generated command, and what executed. We will acknowledge, reproduce against the current source, and fix in the open once a patched release exists — the finding, never your identity, gets the credit unless you ask otherwise.

## Scope notes

- The local model (llama.cpp / Qwen2.5-Coder GGUF) is a separate component under its own licence and security process; aishell pins no model version beyond the setup scripts' download.
- Session history and watchdog files are local user data (see PRIVACY.md); their contents are as sensitive as your shell history — treat them that way.
