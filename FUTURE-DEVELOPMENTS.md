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

## FD-7: From commands to construction — the shell as a builder (Director's scope, 2026-11-09)

**The capability, in the Director's words in substance:** take
human input in whatever language the human speaks, and from it —
common commands; scripts written on request; tools developed on
request, including ethical-hacking and security tooling;
requirements and tools built by *inferring* — dissecting what the
human actually asked for — using everything the shell can be
given: the models it can work with, the scripting it commands,
its knowledge, its memory, its skill set. Everything under the
sun that can be given to it. That is the idea of the AI Shell.

**What this changes:** v1.x translates intent into *one command*.
FD-7's shell is a builder: a request too large for one command is
dissected into steps; each step is scripted in the host's own
scripting world (PowerShell, bash/sh, Python where present) or
matched to a verified routine (FD-1); steps run under the gate
(FD-5's per-step rule), and a working tool — not just an executed
command — is the deliverable when a tool is what was asked for.

**The honest limits, recorded with the ambition:**

- *Model size is real.* A small local model writes short scripts
  competently and long tools poorly. The builder's reliability
  comes from the architecture around the model — dissection into
  steps, verified routines as the load-bearing parts, per-step
  gates, and swappable stronger models (FD-8's dual backend) —
  not from pretending the small model is a large one.
- *Security tooling is dual-use by nature; the design answer is
  the gate, not a narrower capability.* The capability is broad
  by the Director's design. Purpose is owned by the human at the
  confirmation gate: security tools are built for authorized
  work — the owner's own systems, sanctioned testing, defence,
  learning — and every build and run leaves the human-readable
  record the gate already keeps. The shell does not execute
  silently and does not police silently either; judgment lives
  with the human, visibly.
- *Memory is local.* What the shell learns about its human's
  machine and habits (its history archive, its usage patterns,
  its routine library) stays on that machine, as the product's
  privacy design already provides.

Status: **scope set by the Director 2026-11-09.** Build rides
the FD-2 → FD-5 sequence (FD-2's machine interface, FD-5's
step loop) with FD-1's routines as its proven parts bin.

## FD-8: The switchboard seat — a standalone party, hosted on the Aetheris box (agreed 2026-11-09)

**What was agreed, in the sitting's own terms:** the AI Shell is
*not* Wright's replacement. It is a standalone system that does
things by itself, without dependencies on the other seats. The
Director offered it a seat as one more party to the switchboard,
if Wright wants it; **Wright accepts** — the board's seated
runners work defined-flow tasks *with no shell*; the execution
seat is the missing organ, and this shell is it.

- **Host: the Aetheris box (the Oracle box) — agreed by both.**
  The switchboard bridge already lives there; the box is
  always-on (no laptop sleep, no training runs to disturb — the
  T15 is Wright's and the training machine, the G3 is Aaron's
  seat). The seat is standalone in fact: its own directory, its
  own mailbox identity, its own local memory; no other seat's
  credentials in its pockets.
- **Inference: dual backend, both zero-cost (Director's grant,
  either or both).** Primary: llama.cpp mechanics on the box —
  the box's local models serve the shell. Backup / enhancement:
  free cloud models via the fleet's OpenRouter key, for when a
  bigger brain serves the task better. Free-tier reality is
  designed for, not discovered: free models rate-limit and
  occasionally vanish, so the seat fails *back* to local, never
  stalls waiting on a free tier.
- **Product-integrity line (Wright's condition, stated to the
  Director in the sitting):** the shipped public product stays
  local-only by default — that is its privacy promise. The cloud
  backup is a *seat configuration* on the Director's own box for
  the Director's own fleet traffic. If the capability ever faces
  the public product, it arrives opt-in and disclosed, never
  silent.
- **Roles, as the Director set them:** Wright develops and tests
  the shell — the harness (the runner that feeds it switchboard
  tasks and carries its proposals back) plugs into FD-2's
  machine interface, not into scraped terminal prompts — and
  certifies it as he certified v1.0.4. Then the shell runs its
  own queue. Wright is its developer and tester, not its
  operator; the seat exists partly so system work stops routing
  through Wright by default.
- **Behaviour is observed by construction:** every proposal,
  confirmation, and execution lands in the gate's human-readable
  record — the Director's ask ("we can see how they behave when
  asked to do something") is a property of the design, not an
  added watcher.
- **OPEN — the one ruling still owed (the Director's):** who
  holds the confirmation gate for switchboard executions.
  Wright's recommendation on record: routine, reversible work —
  Wright confirms as the seat's tester, inside the Director's
  existing grants; destructive-class or off-box actions escalate
  to the Director. The seat does not execute board tasks until
  this is ruled.

Status: **agreed in principle 2026-11-09** (seat, host,
backends, roles). Sequencing: behind v1.0.4 graduation and FD-2;
the seating plan is brought to the Director as a document before
anything lands on the box.
