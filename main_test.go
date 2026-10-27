package main

import (
	"bytes"
	"strings"
	"testing"
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
	cap := streamInference(r, thinkingToken, responseToken)
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
	cap := streamInference(r, thinkingToken, responseToken)
	if !strings.Contains(cap, "ls") {
		t.Fatalf("marker split across writes not handled: %q", cap)
	}
}
