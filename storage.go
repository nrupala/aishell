package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type WatchdogMetrics struct {
	CommandFrequencies map[string]int `json:"command_frequencies"`
}

type WindowsEventLog struct {
	EventID         int    `json:"EventID"`
	ProviderName    string `json:"ProviderName"`
	TimeCreated     string `json:"TimeCreated"`
	Task            string `json:"Task"`
	UserIntent      string `json:"UserIntent"`
	ScriptPayload   string `json:"ScriptPayload"`
	ExecutionResult string `json:"ExecutionResult"`
}

type UnixSyslogMessage struct {
	Timestamp int64  `json:"timestamp"`
	Host      string `json:"host"`
	Process   string `json:"process"`
	Message   string `json:"message"`
	Command   string `json:"command"`
	Status    string `json:"status"`
}

var (
	watchdogStore  WatchdogMetrics
	windowsHistory []WindowsEventLog
	unixHistory    []UnixSyslogMessage
)

func configDir() string {
	if runtime.GOOS == "windows" {
		base := os.Getenv("APPDATA")
		if base == "" {
			base = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
		return filepath.Join(base, "ai_shell")
	}
	base := os.Getenv("XDG_DATA_HOME")
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			home = "."
		}
		base = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(base, "ai_shell")
}

func metricsPath() string {
	return filepath.Join(configDir(), "watchdog_metrics.json")
}

func loadWatchdog() {
	watchdogStore.CommandFrequencies = make(map[string]int)
	data, err := os.ReadFile(metricsPath())
	if err == nil {
		_ = json.Unmarshal(data, &watchdogStore)
	}
	if watchdogStore.CommandFrequencies == nil {
		watchdogStore.CommandFrequencies = make(map[string]int)
	}
}

func saveWatchdog() {
	_ = os.MkdirAll(configDir(), 0o755)
	data, err := json.MarshalIndent(watchdogStore, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(metricsPath(), data, 0o644)
}

func trackWatchdog(command string) {
	watchdogStore.CommandFrequencies[command]++
	if watchdogStore.CommandFrequencies[command] == 3 {
		fmt.Printf("\n%s[Watchdog Alert]:%s You have executed this payload 3 times. Consider promoting it to a pre-baked alias.\n", ColorMagenta, ColorReset)
		fmt.Printf("  %s%s%s\n", ColorYellow, command, ColorReset)
	}
}

func recordPlatformLog(hostOS, intent, command, status string) {
	if hostOS == "windows" {
		event := WindowsEventLog{
			EventID:         4624,
			ProviderName:    "AIShellEngine",
			TimeCreated:     time.Now().Format("2006-01-02T15:04:05.0000000-07:00"),
			Task:            "AI_Intent_Translation_Gate",
			UserIntent:      intent,
			ScriptPayload:   command,
			ExecutionResult: status,
		}
		windowsHistory = append(windowsHistory, event)
		return
	}
	hostname, _ := os.Hostname()
	msg := UnixSyslogMessage{
		Timestamp: time.Now().Unix(),
		Host:      hostname,
		Process:   "ai_shell",
		Message:   fmt.Sprintf("Intent process complete. Input value: '%s'", intent),
		Command:   command,
		Status:    status,
	}
	unixHistory = append(unixHistory, msg)
}

func flushHistory(reader *bufio.Reader, hostOS string) {
	saveWatchdog()

	var count int
	var data []byte
	var name string
	historyDir := filepath.Join(configDir(), "history")

	if hostOS == "windows" {
		count = len(windowsHistory)
		name = fmt.Sprintf("AppEventLog_%s.json", time.Now().Format("20060102_150405"))
		if count > 0 {
			data, _ = json.MarshalIndent(windowsHistory, "", "  ")
		}
	} else {
		count = len(unixHistory)
		name = fmt.Sprintf("syslog_%d.json", time.Now().Unix())
		if count > 0 {
			data, _ = json.MarshalIndent(unixHistory, "", "  ")
		}
	}

	if count == 0 {
		fmt.Println("No transactional entries recorded. Session closed.")
		return
	}

	fmt.Print("Archive session history before closing? [y/N]: ")
	choice := readChoice(reader)
	if choice != "y" && choice != "yes" && choice != "" {
		fmt.Println("Session history purged.")
		return
	}

	_ = os.MkdirAll(historyDir, 0o755)
	path := filepath.Join(historyDir, name)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		fmt.Printf("%s[IO Flush Fault]:%s %v\n", ColorRed, ColorReset, err)
		return
	}
	fmt.Printf("%s[History Archived]:%s %s\n", ColorGreen, ColorReset, path)
}
