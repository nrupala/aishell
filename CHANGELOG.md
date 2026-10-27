# Changelog — aishell

## Unreleased — free-release prep (release-prep branch)

Prepared 2026-10-26 by Wright (2351485 ALBERTA INC.). Not published.

- Storefront README rewritten: clean markdown, verified build/test instructions, honest security model with stated limits.
- Added PRIVACY.md (source-verified), NOTICE, LICENSE (Apache 2.0, proposed — owner's licence word pending), .github/FUNDING.yml, SECURITY.md, CI workflow (vet + test + all-target builds).
- Fix: `setup.sh` saved the model under the Hugging Face filename, which does not match the path `main.go` loads (`models/qwen2.5-coder-3b-instruct-q4_k_m.gguf`) on case-sensitive filesystems — first run on Linux/macOS would have failed to find the model. Setup now saves under the loaded name.
- Fix: session history was archived on an empty answer at the `[y/N]` archive prompt (default meant archive). History now archives only on an explicit `y`/`yes`; verified by running the built binary both ways.
- Verified at prep: `go vet` clean, `go test` 10/10 pass, binaries built from this source for Windows (amd64), Linux (amd64), macOS (arm64).

## 1.0.0 — 2026-08-15 (original build, custody import)

- Original complete build (By_Gemini build record): Go sources, tests, Makefile/build scripts, prebuilt binaries for Windows/Linux/macOS, design docs. Imported as-is into this repository's `main` branch as the custody record — see the import commits.
