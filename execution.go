package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var DestructiveBlacklist = []string{
	"rm -rf", "rm -fr", "rmdir /s", "del /f", "format ", "mkfs",
	"dd if=", ">/dev/sda", "drop database", "Format-Volume",
	"Remove-Item -Recurse -Force", "shred ",
	"rm --recursive",
}

var MutationOperators = []string{
	">", ">>", "Out-File", "Set-Content", "Add-Content", "tee ",
}

func verifySandboxSafety(command string) (bool, bool) {
	lower := strings.ToLower(command)
	for _, restricted := range DestructiveBlacklist {
		if strings.Contains(lower, strings.ToLower(restricted)) {
			return false, true
		}
	}
	for _, op := range MutationOperators {
		if strings.Contains(command, op) {
			return false, false
		}
	}
	return true, false
}

func formatEditorFallback(hostOS, editor, original string) string {
	parts := strings.Fields(original)
	target := "sandbox_output.txt"
	if len(parts) > 0 {
		target = strings.Trim(parts[len(parts)-1], `"' `)
	}
	if hostOS == "windows" {
		return fmt.Sprintf("Start-Process %s '%s'", editor, target)
	}
	return fmt.Sprintf("%s %s", editor, target)
}

func buildExecCommand(hostOS, payload string) *exec.Cmd {
	switch hostOS {
	case "windows":
		return exec.Command("powershell.exe", "-NoProfile", "-Command", payload)
	case "linux":
		// systemd-run --scope, with NO service sandbox properties.
		// The original code passed -p PrivateDevices / ProtectHome /
		// ProtectSystem here, but scope units reject every one of them
		// ("Unknown assignment") on systemd 255 / Ubuntu 24.04 — so
		// every user-approved command failed without running (found in
		// the 2026-10-30 Aetheris end-to-end proof, session s1). The
		// service-unit form (--wait --pipe) accepts the properties but
		// PrivateDevices makes it die with status 218 in a user
		// manager, and ProtectHome did not actually block a home write
		// in testing — claiming either would be false. A bare scope
		// still groups the command under the user manager; on Linux
		// the screening + confirmation gate is the control, and the
		// docs now say exactly that.
		return exec.Command("systemd-run",
			"--user", "--scope",
			"/bin/bash", "-c", payload)
	case "darwin":
		return exec.Command("sandbox-exec",
			"-p", `(version 1) (allow default) (deny file-write* (subpath "/"))`,
			"/bin/zsh", "-c", payload)
	default:
		return exec.Command("/bin/sh", "-c", payload)
	}
}

func executeCommand(hostOS, payload string) {
	cmd := buildExecCommand(hostOS, payload)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("\n%s[OS Sandbox Block / Runtime Error]:%s %v\n", ColorRed, ColorReset, err)
	}
}
