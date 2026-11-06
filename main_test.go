// SPDX-FileCopyrightText: 2026 2351485 ALBERTA INC.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestExtractCommandBlockFenceWithLanguage(t *testing.T) {
	out := thinkingToken + "\nUse powershell.\n" + responseToken + "\n" +
		"```powershell\nGet-ChildItem -Path . -Filter *.log -Recurse | Where-Object {$_.Length -gt 100MB}\n```\n"
	got := extractCommandBlock(out)
	want := "Get-ChildItem -Path . -Filter *.log -Recurse | Where-Object {$_.Length -gt 100MB}"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestExtractCommandBlockBashFence(t *testing.T) {
	out := "junk\n```bash\nfind . -type f -size +100M\n```\ntrailing"
	got := extractCommandBlock(out)
	if got != "find . -type f -size +100M" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractCommandBlockNoFence(t *testing.T) {
	if got := extractCommandBlock("just prose, no fence"); got != "" {
		t.Fatalf("got %q want empty", got)
	}
}

func TestStreamInferenceCapturesFullOutput(t *testing.T) {
	src := " pre " + thinkingToken + "\nreason\n" + responseToken + "\n```bash\necho hi\n```\n"
	r := strings.NewReader(src)
	cap := streamInference(r, thinkingToken, responseToken, nil)
	if !strings.Contains(cap, "reason") {
		t.Fatalf("missing reasoning in captured output: %q", cap)
	}
	if !strings.Contains(cap, "echo hi") {
		t.Fatalf("missing command in captured output: %q", cap)
	}
	if !strings.Contains(cap, " pre ") {
		t.Fatalf("missing leading content: %q", cap)
	}
}

func TestFlushableLenWithholdsPartialMarker(t *testing.T) {
	if n := flushableLen("abc think", thinkingToken, responseToken); n != 3 {
		t.Fatalf("expected flush up to partial marker, got %d", n)
	}
	if n := flushableLen("abcdef", thinkingToken, responseToken); n != 6 {
		t.Fatalf("expected full flush when no marker, got %d", n)
	}
	if n := flushableLen("abc thinking rest", thinkingToken, responseToken); n != 3 {
		t.Fatalf("expected flush to stop before full marker, got %d", n)
	}
}

func TestVerifySandboxSafetyBlacklist(t *testing.T) {
	if safe, black := verifySandboxSafety("rm -rf /"); safe || !black {
		t.Fatalf("expected destructive block, got safe=%v black=%v", safe, black)
	}
}

func TestVerifySandboxSafetyMutationReroute(t *testing.T) {
	if safe, black := verifySandboxSafety("echo hi >> out.txt"); safe || black {
		t.Fatalf("expected mutation reroute, got safe=%v black=%v", safe, black)
	}
}

func TestBuildSystemPromptContainsSchema(t *testing.T) {
	p := buildSystemPrompt("linux", "vim", "find logs")
	for _, frag := range []string{thinkingToken, responseToken, "```cmd", "User Intent: find logs"} {
		if !strings.Contains(p, frag) {
			t.Fatalf("system prompt missing %q\n%s", frag, p)
		}
	}
}

func TestFormatEditorFallback(t *testing.T) {
	got := formatEditorFallback("linux", "vim", "rm config.json")
	if got != "vim config.json" {
		t.Fatalf("got %q", got)
	}
	gotWin := formatEditorFallback("windows", "notepad.exe", "Remove-Item hosts")
	if !strings.Contains(gotWin, "notepad.exe") || !strings.Contains(gotWin, "hosts") {
		t.Fatalf("got %q", gotWin)
	}
}

func TestStreamInferenceTerminatesAcrossChunkEdges(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString(responseToken[:6])
	buf.WriteString(responseToken[6:])
	buf.WriteString("\n```bash\nls\n```\n")
	r := &buf
	cap := streamInference(r, thinkingToken, responseToken, nil)
	if !strings.Contains(cap, "ls") {
		t.Fatalf("marker split across writes not handled: %q", cap)
	}
}

func TestExtractCommandBlockSkipsEchoedTemplate(t *testing.T) {
	// Simulates a current llama-cli stream: the system prompt is echoed
	// (including its empty template block) before the model generates.
	out := "echoed prompt\n```cmd\n[THE EXACT COMMAND LINE PAYLOAD TO EXECUTE]\n```\n" +
		thinkingToken + "\nreason\n" + responseToken + "\n```cmd\nls -la\n```\n"
	got := extractCommandBlock(out)
	if got != "ls -la" {
		t.Fatalf("got %q want %q", got, "ls -la")
	}
}

func TestExtractCommandBlockTakesLastBlock(t *testing.T) {
	out := "```bash\necho first\n```\nnoise\n```bash\necho second\n```\n"
	if got := extractCommandBlock(out); got != "echo second" {
		t.Fatalf("got %q want %q", got, "echo second")
	}
}

func TestBuildExecCommandLinuxScopeHasNoServiceProps(t *testing.T) {
	cmd := buildExecCommand("linux", "ls")
	joined := strings.Join(cmd.Args, " ")
	if !strings.Contains(joined, "systemd-run --user --scope") {
		t.Fatalf("linux exec should use a systemd-run scope, got %q", joined)
	}
	for _, prop := range []string{"PrivateDevices", "ProtectHome", "ProtectSystem"} {
		if strings.Contains(joined, prop) {
			t.Fatalf("scope units reject %s (systemd 255 verified); must not be passed: %q", prop, joined)
		}
	}
	if cmd.Args[len(cmd.Args)-1] != "ls" {
		t.Fatalf("payload should be the last arg, got %q", joined)
	}
}

func TestLlamaArgsUseOnlyStableCoreFlags(t *testing.T) {
	args := llamaArgs("hello")
	joined := strings.Join(args, " ")
	for _, want := range []string{"-m", "-c", "-n", "--temp", "-p"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("llama args missing stable core flag %s: %q", want, joined)
		}
	}
	// No conversation or display flags at all: upstream removes and
	// renames that surface (both conversation toggles died within a
	// year, breaking v1.0.1 and v1.0.2 in production). One-shot
	// behaviour is enforced by runInference's process control, in
	// this codebase, where upstream cannot remove it.
	for _, bad := range []string{"-no-cnv", "--no-conversation", "-cnv", "--conversation",
		"--single-turn", " -st ", "--no-display-prompt", "--display-prompt"} {
		if strings.Contains(joined, bad) {
			t.Fatalf("llama args must not use conversation/display flag %q: %q", bad, joined)
		}
	}
	if args[len(args)-1] != "hello" {
		t.Fatalf("prompt should be the last arg, got %q", joined)
	}
}

// writeFakeLlama installs a fake llama-cli script in a temp dir and
// puts it first in PATH. The script ignores its arguments, emits
// body, then sleeps — standing in for a conversation-mode build
// that generates once and then waits forever for the next turn.
func writeFakeLlama(t *testing.T, body string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("fake llama-cli is a POSIX shell script")
	}
	dir := t.TempDir()
	// The sleep's descriptors are redirected: an orphaned sleep that
	// inherits the test binary's stderr would hold go test's output
	// pipe open and stall the suite for the full 60s.
	script := "#!/bin/sh\nprintf '%s' '" + body + "'\nsleep 60 >/dev/null 2>&1\n"
	if err := os.WriteFile(filepath.Join(dir, "llama-cli"), []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func TestRunInferenceTerminatesAfterCompleteBlock(t *testing.T) {
	writeFakeLlama(t, " thinking\nreason\n response\n```cmd\nls\n```\n")
	old := inferenceTimeout
	inferenceTimeout = 30 * time.Second
	defer func() { inferenceTimeout = old }()

	start := time.Now()
	out, err := runInference("list files")
	if err != nil {
		t.Fatalf("runInference: %v", err)
	}
	if elapsed := time.Since(start); elapsed > 10*time.Second {
		t.Fatalf("inference did not terminate on completion (took %s) — the fake sleeps 60s", elapsed)
	}
	if got := extractCommandBlock(out); got != "ls" {
		t.Fatalf("got block %q from output %q", got, out)
	}
}

func TestRunInferenceTimesOut(t *testing.T) {
	writeFakeLlama(t, "")
	old := inferenceTimeout
	inferenceTimeout = time.Second
	defer func() { inferenceTimeout = old }()

	_, err := runInference("list files")
	if err == nil || !strings.Contains(err.Error(), "timed out") {
		t.Fatalf("expected timeout error, got %v", err)
	}
}
