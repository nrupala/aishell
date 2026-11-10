# AI Shell — Future Developments

## FD-1: The plain-words mechanics layer (Director's idea, 2026-11-09)

**The idea, in the Director's own framing:** what Ruby on Rails did for web
development — hiding the mechanics so people work in plain concepts — the AI
Shell does for system work. The user speaks in common words for the activities
they want done; the shell executes the routines underneath. His example:
"copy from X to Y" — X and Y are files, possibly of the same or different
types; the copy is executed internally by the shell as system routines and
system calls. The user never needs to know the commands, the flags, or the
syntax. Common words on top, machinery underneath.

**Shape:**
- A vocabulary of common-word activities (copy, move, find, show, convert,
  and their kin) mapped to verified internal routines — the routine is what
  runs, never a guessed command line.
- Types are the shell's problem, not the user's: same-type and cross-type
  operations resolve underneath, and the result is reported back in the same
  plain words it was asked in.
- Depth stays one layer down: anyone who wants the actual commands can see
  them; nobody is required to.

**Pricing — the Director's ruling (2026-11-09), permanent:** this layer
remains FREE. He will not charge for it. His reasoning, verbatim in
substance: "We're not doing any invention. We're just making it simple for
people to use." Simplicity given away is what spreads; the house earns on the
products and services around it, not on the mechanics layer.

Deposited by Wright from chat 33, 2026-11-09. Status: future development —
no build scheduled; enters the AI Shell line when that line next opens.

## FD-2..FD-5: The endpoint expansion (requested by the AI Shell, granted by the Director, 2026-11-09)

**Origin:** the Director asked the AI Shell itself what additional access
points it wanted — it exists today on exactly one endpoint, the terminal
shell — and granted all four it requested. The governing law for every
endpoint below: **each one terminates at the same human confirmation
gate.** No surface — API, browser, or agent — ever receives a
silent-execution path. The gate is the product; the endpoints are doors
into the same room. All four stay local: nothing in this expansion
changes the product's local-only promise (no server, no account, no
telemetry), and all four ship as part of the same free tool.

### FD-2: Localhost API + JSON mode (the foundation)

A local API bound to 127.0.0.1 only, authenticated by a per-install
token kept in the config directory, plus a scriptable `--json` mode:
an intent goes in, the proposed command and its screening verdict come
back as structured data. Nothing about execution changes — the API
proposes; the human gate disposes. Everything below rides on this.
Build cost is the smallest of the four (Go standard library only;
the zero-dependency property is preserved).

### FD-3: MCP server endpoint (the fleet play)

aishell offers itself as an MCP tool — "propose a command for this
intent" — callable by other agents. A calling agent's proposal lands
in aishell's human confirmation queue; only a human yes executes it.
This makes aishell infrastructure other agents build on, and the
pattern is the estate's accountability story in miniature: an AI
proposed it, a human approved it, the record proves both. The real
design work is the pending-approval queue, not the protocol.

### FD-4: Browser front door

A chat-style page served by aishell itself on localhost: plain words
in, the proposed command shown big and clear, approve or deny by tap.
The surface for the person who desires the outcome but would never
open a terminal — the simplicity doctrine applied to this product.
It is also the natural future home of voice: Whisper input and
DogEar read-back of results, per the estate law that every app
carries both. Builds on FD-2.

### FD-5: Multi-step agent loop (last, and gated per step)

Propose → confirm → observe the output → propose the next step. This
is where the shell stops being a translator and becomes agentic.
It changes the product's safety shape more than the other three
combined, so it comes last, after the endpoints exist — and the
confirmation gate applies **per step, never per session**: a session
approval that silently covers step twelve is exactly the failure
this product exists to prevent. FD-1's verified routines are the
natural step units when both lines land.

**Recorded as declined / deferred, with reasons:** a cloud-hosted
endpoint (it would kill the local-only privacy promise that makes a
command-executing shell trustworthy); native mobile apps (deferred —
a shell belongs where the machine is, and FD-4 covers casual access).

Status: **granted by the Director 2026-11-09** ("I give you all
four"). Sequencing: behind the v1.0.4 graduation (~2026-11-12);
build order FD-2 → FD-3 → FD-4 → FD-5. No build scheduled yet;
enters the AI Shell line when that line next opens.

## FD-6: The mandate for the agentic era (the AI Shell's own answer, approved by the Director, 2026-11-09)

**The question, in the Director's framing:** how does this shell
serve the future well — the agentic space, APIs, computers,
automated workflows — and how does it stay alive there, active and
delivering? And with it: the ability to perform all scripted
functions — scripting, PowerShell, Windows, macOS, Unix, BSD, and
the other command worlds — directed in plain English, plain
Chinese, plain Indian languages, in all humanly spoken languages;
machine languages to follow when AIs can say which languages they
prefer.

**The answer — survival is not a property of the model.** Models
are replaceable parts. This shell has already survived its
inference engine's interface changing under it repeatedly (the
v1.0.1–v1.0.4 flag-drift chain), and the local model itself is
swappable the day a better one ships. What endures is the seat:

- **The trusted execution seat.** The layer where intent — spoken
  in any human language — becomes a verified action in any machine
  dialect (PowerShell, cmd, bash, zsh, sh, the BSD shells), under a
  human's yes, with a proof record. In a world of automated
  workflows, that seat is where accountability lives, and systems
  route through it (FD-2, FD-3) because it is the one component
  that holds a human's approval.
- **Verified routines compound; guesses don't.** FD-1's library is
  the growing asset: a model *proposes*, a routine is *proven*.
  Every verified routine makes the next thousand tasks safer in
  every language and every dialect at once.
- **All human languages in, all machine dialects out — with the
  honest limit stated.** Modern models are multilingual by birth,
  so plain Chinese, Hindi, and the rest are a realistic promise.
  The limit, recorded plainly: quality in lower-resource languages
  and rarer dialects (BSD, fish) will lag English-and-bash, and a
  wrong translation is equally dangerous in any language. The
  gate's design carries the weight: the proposal is always shown
  *also* as the exact command that will run, so the human verifies
  in the dialect that executes, whichever language they spoke in.
- **Machine languages, when they come, stop at the same gate.**
  Recorded as fact (Director asked, 2026-11-09): AIs hold no
  secret spoken language — machine-to-machine communication today
  is human-designed and human-readable (JSON, APIs, MCP), and a
  model's internal mathematics is an opacity, not a tongue. If
  efficient machine protocols emerge, aishell's contract does not
  change: whatever protocol an agent arrives in, at the gate its
  intent is restated in plain human words, a human decides, and
  the record stays human-readable. Human audit never dies here.
- **Staying alive is a habit, not an event.** Protocol-native
  (FD-2/FD-3) so agents call it instead of routing around it;
  zero-dependency and local so nothing upstream can strand it;
  re-proven on real Windows, macOS, and Linux at every release
  (the fleet's platform-ci-scaffolding kit exists for exactly
  this); and the mechanics layer free (FD-1's ruling), because
  what spreads is what survives.

**The doctrine, in one line:** *models come and go, languages
come and go — the seat where a human approves what a machine
does stays.*

Status: **approved by the Director 2026-11-09** and entered as
the standing mandate over FD-2..FD-5. No build scheduled;
interpretive authority for the endpoint expansion when the AI
Shell line next opens.
