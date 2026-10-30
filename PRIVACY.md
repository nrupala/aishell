# Privacy — aishell

Copyright © 2026 2351485 ALBERTA INC.

aishell runs fully local. This statement was verified against the released source (no network code exists in the Go sources — check for yourself: there is not a single HTTP call in the module).

## What happens on your machine, and nowhere else

- Your natural-language input is sent only to a `llama-cli` process running on your own machine, with a model file on your own disk.
- aishell has no account system, no telemetry, no analytics, and no update channel. It cannot phone home; there is no home to phone.
- Commands execute only on your machine, only after your manual confirmation at the prompt.

## What aishell stores locally

- **Watchdog counts** — how often each executed command has run, in `watchdog_metrics.json` under your per-user data folder (`%APPDATA%\ai_shell` on Windows, `$XDG_DATA_HOME/ai_shell` or `~/.local/share/ai_shell` on Linux/macOS). This powers the "you run this often — make it an alias" suggestion. It never leaves the machine.
- **Session history (opt-in only)** — when you exit, aishell asks `Archive session history before closing? [y/N]`. Only an explicit `y`/`yes` writes the session's intents, proposed commands, and outcomes to a JSON file in the `history` subfolder of the same local data folder. Any other answer — including just pressing Enter — purges the in-memory history and writes nothing.

Delete either at any time by deleting those files; aishell recreates empty state on next run.

## The one network step, and it is yours

The optional setup scripts (`setup.sh` / `setup.bat`) download the GGUF model file from Hugging Face. That is a download you initiate, from Hugging Face's servers, under their terms. The aishell binary itself makes no network connections; if you already have the model file, you never need the download.

## If you change the setup

aishell talks to inference only through the local `llama-cli` subprocess. If you modify it to use a remote endpoint, traffic to that endpoint is governed by that operator — that configuration is not the released one described here.
