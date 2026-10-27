package main

import "strings"

func extractCommandBlock(output string) string {
	lines := strings.Split(output, "\n")
	var capture []string
	inFence := false

	for _, ln := range lines {
		trimmed := strings.TrimSpace(ln)
		if !inFence && strings.HasPrefix(trimmed, "```") {
			inFence = true
			continue
		}
		if inFence && trimmed == "```" {
			inFence = false
			break
		}
		if inFence {
			capture = append(capture, ln)
		}
	}

	if len(capture) > 0 {
		return strings.TrimSpace(strings.Join(capture, "\n"))
	}

	fences := []string{"```cmd", "```powershell", "```bash", "```zsh", "```shell", "```"}
	for _, fence := range fences {
		if i := strings.Index(output, fence); i >= 0 {
			rest := output[i+len(fence):]
			if j := strings.Index(rest, "```"); j >= 0 {
				return strings.TrimSpace(rest[:j])
			}
		}
	}
	return ""
}
