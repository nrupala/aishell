package main

import "strings"

// promptTemplateBlock is the empty command template inside the system
// prompt itself. Current llama-cli builds echo the prompt into their
// output, so the echoed template is present in the captured stream and
// must never be mistaken for the model's answer.
const promptTemplateBlock = "[THE EXACT COMMAND LINE PAYLOAD TO EXECUTE]"

func extractCommandBlock(output string) string {
	// Collect every complete fenced block. The model's generated block
	// comes after any prompt echo, so the LAST real block wins; the
	// echoed template block is filtered out. (Before the 2026-10-30
	// end-to-end proof this returned the FIRST block — with a current
	// llama-cli that would have proposed the placeholder template as
	// the command.)
	lines := strings.Split(output, "\n")
	var blocks []string
	var capture []string
	inFence := false

	for _, ln := range lines {
		trimmed := strings.TrimSpace(ln)
		if !inFence && strings.HasPrefix(trimmed, "```") {
			inFence = true
			capture = nil
			continue
		}
		if inFence && trimmed == "```" {
			inFence = false
			block := strings.TrimSpace(strings.Join(capture, "\n"))
			if block != "" && block != promptTemplateBlock {
				blocks = append(blocks, block)
			}
			capture = nil
			continue
		}
		if inFence {
			capture = append(capture, ln)
		}
	}

	if len(blocks) > 0 {
		return blocks[len(blocks)-1]
	}

	fences := []string{"```cmd", "```powershell", "```bash", "```zsh", "```shell", "```"}
	for _, fence := range fences {
		if i := strings.Index(output, fence); i >= 0 {
			rest := output[i+len(fence):]
			if j := strings.Index(rest, "```"); j >= 0 {
				candidate := strings.TrimSpace(rest[:j])
				if candidate != "" && candidate != promptTemplateBlock {
					return candidate
				}
			}
		}
	}
	return ""
}
