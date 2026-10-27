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

func executeCommand(hostOS, payload string) {
	var cmd *exec.Cmd

	switch hostOS {
	case "windows":
		cmd = exec.Command("powershell.exe", "-NoProfile", "-Command", payload)
	case "linux":
		cmd = exec.Command("systemd-run",
			"--user", "--scope",
			"-p", "PrivateDevices=true",
			"-p", "ProtectHome=read-only",
			"-p", "ProtectSystem=strict",
			"/bin/bash", "-c", payload)
	case "darwin":
		cmd = exec.Command("sandbox-exec",
			"-p", `(version 1) (allow default) (deny file-write* (subpath "/"))`,
			"/bin/zsh", "-c", payload)
	default:
		cmd = exec.Command("/bin/sh", "-c", payload)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("\n%s[OS Sandbox Block / Runtime Error]:%s %v\n", ColorRed, ColorReset, err)
	}
}
