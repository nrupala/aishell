package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorBold    = "\033[1m"
	ColorDim     = "\033[2m"

	modelPath     = "models/qwen2.5-coder-3b-instruct-q4_k_m.gguf"
	contextWindow = 4096
	maxTokens     = 512
	temperature   = 0.1

	thinkingToken = " thinking"
	responseToken = " response"
)

var PreBakedAliases = map[string]string{
	"sysinfo": "Show system resources, CPU, memory, and OS architecture metrics.",
	"netstat": "List active local network connections, open listening ports, and routing statistics.",
	"diskuse": "Show an aggregate summary of disk storage usage and available partitions.",
}

func main() {
	hostOS := runtime.GOOS
	shellEngine, defaultEditor := getOSContext(hostOS)

	loadWatchdog()
	printBanner(hostOS, shellEngine)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n%sAI-Shell>%s ", ColorGreen, ColorReset)
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "quit" {
			saveWatchdog()
			flushHistory(reader, hostOS)
			fmt.Println("Exiting AI-Native Shell. Sandbox closed safely.")
			return
		}

		switch strings.ToLower(input) {
		case "help", "-h", "--help":
			printHelp()
			continue
		case "aliases", "list-aliases":
			printAliases()
			continue
		}

		if expanded, ok := PreBakedAliases[strings.ToLower(input)]; ok {
			fmt.Printf("%s[Alias Intercepted]:%s '%s' -> '%s'\n", ColorBlue, ColorReset, input, expanded)
			input = expanded
		}

		systemPrompt := buildSystemPrompt(hostOS, defaultEditor, input)

		fmt.Printf("\n%s[Thinking Process]%s\n", ColorYellow, ColorReset)
		raw, err := runInference(systemPrompt)
		if err != nil {
			fmt.Printf("%s[Inference Error]:%s %v\n", ColorRed, ColorReset, err)
			recordPlatformLog(hostOS, input, "", "INFERENCE_ERROR")
			continue
		}

		command := extractCommandBlock(raw)
		if command == "" {
			fmt.Printf("%s[Security/Parsing Guard]: No valid or safe command block was extracted.%s\n", ColorRed, ColorReset)
			recordPlatformLog(hostOS, input, "", "BLOCKED_PARSE")
			continue
		}

		safe, blacklisted := verifySandboxSafety(command)
		if !safe {
			if blacklisted {
				fmt.Printf("%s[CRITICAL SECURITY BLOCK]: Destructive pattern detected. Aborted.%s\n", ColorRed, ColorReset)
				recordPlatformLog(hostOS, input, command, "BLOCKED_SECURITY")
				continue
			}
			fmt.Printf("%s[Sandbox Intercept]: Write mutation detected. Rerouting to native viewer/editor.%s\n", ColorYellow, ColorReset)
			command = formatEditorFallback(hostOS, defaultEditor, command)
		}

		fmt.Printf("\n%sProposed Command:%s\n", ColorCyan, ColorReset)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Printf("%s%s%s\n", ColorYellow, command, ColorReset)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Print("Execute command on host system? [y/N]: ")

		choice := readChoice(reader)
		if choice == "y" || choice == "yes" {
			fmt.Printf("%s[Executing Subprocess]%s\n", ColorGreen, ColorReset)
			executeCommand(hostOS, command)
			trackWatchdog(command)
			recordPlatformLog(hostOS, input, command, "EXECUTED")
		} else {
			fmt.Printf("%s[Execution Aborted by User]%s\n", ColorRed, ColorReset)
			recordPlatformLog(hostOS, input, command, "USER_DECLINED")
		}
	}
}

func getOSContext(goos string) (string, string) {
	switch goos {
	case "windows":
		return "powershell", "notepad.exe"
	case "linux":
		return "/bin/bash -c", "vim"
	case "darwin":
		return "/bin/zsh -c", "open -e"
	default:
		return "/bin/sh -c", "vi"
	}
}

func readChoice(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(line))
}

func printBanner(hostOS, shellEngine string) {
	fmt.Printf("%s[AI-Native Shell Initialized with Watchdog Systems]%s\n", ColorGreen, ColorReset)
	fmt.Printf("Detected OS: %s%s%s | Target Engine: %s%s%s\n", ColorCyan, hostOS, ColorReset, ColorCyan, shellEngine, ColorReset)
	fmt.Println("-----------------------------------------------------------------")
}

func printHelp() {
	fmt.Printf("\n%sAI SHELL MANUAL%s\n", ColorBold, ColorReset)
	fmt.Printf("  %-15s - Displays the structural welcome menu\n", "help")
	fmt.Printf("  %-15s - Lists all currently configured shortcut aliases\n", "aliases")
	fmt.Printf("  %-15s - Safely terminates the interactive shell session\n", "exit / quit")
	fmt.Println("  Plain strings are processed by the local Qwen2.5-Coder model into shell code.")
	fmt.Println("  Destructive commands are caught and rejected before reaching the gate.")
	fmt.Println("  Direct file mutation commands launch native text view/editors only.")
}

func printAliases() {
	fmt.Printf("\n%s[Configured Pre-Baked Commands & Aliases]%s\n", ColorBold, ColorReset)
	fmt.Println("-----------------------------------------------------------------")
	for alias, expansion := range PreBakedAliases {
		fmt.Printf("  %s%-10s%s -> %s\n", ColorYellow, alias, ColorReset, expansion)
	}
	fmt.Println("-----------------------------------------------------------------")
}
