\# AI-Native Cross-Platform Shell Interface



A lightweight, zero-dependency command-line interface wrapper written in Go. The application intercepts natural language user inputs, streams local inference via a background `llama.cpp` process, extracts clean system commands optimized for the host Operating System, and executes them following manual user confirmation via a security sandbox gate.



\---



\## 🏗️ Quick Start Architecture Setup



\### 1. Project Folder Structure

Ensure your workspace is structured as follows before running the compilation scripts:

```text

.

├── main.go               # Core Go orchestrator

├── Makefile              # Unix/Linux/macOS build pipeline

├── build.bat             # Windows native build pipeline

└── models/               # Target local GGUF model repository

&#x20;   └── qwen2.5-coder-3b-instruct-q4\_k\_m.gguf (To be downloaded)

```



\### 2. Download the Reasoning Model

To enable lightweight execution across standard system hardware and mobile devices, we use the optimized `Qwen2.5-Coder` model quantization series.



1\. Navigate into your local models directory:

&#x20;  ```bash

&#x20;  cd models

&#x20;  ```

2\. Download the recommended model binary footprint using `curl` or your browser:

&#x20;  ```bash

&#x20;  curl -L -O https://huggingface.co

&#x20;  ```



\---



\## 🛠️ Compilation \& Execution



Build binaries automatically using the compiled script engine matching your development computer platform:



\*   \*\*On macOS / Linux Environments\*\*:

&#x20;   ```bash

&#x20;   make

&#x20;   ./ai-shell

&#x20;   ```

\*   \*\*On Windows OS Environments\*\*:

&#x20;   ```powershell

&#x20;   .\\build.bat

&#x20;   .\\ai-shell.exe

&#x20;   ```



\---



\## ⚙️ Configuring the `llama-cli` Dependency Path



The Go orchestrator communicates directly with `llama.cpp` by calling the `llama-cli` binary inside its execution path loop. If `llama-cli` is not globally mapped, your shell will throw an execution path runtime error.



Choose the mapping method that best matches your deployment context below:



\### Option A: System Path Mapping (Recommended)

Place your pre-compiled `llama-cli` binary directly into one of your operating system's global system path variables so it can be invoked safely from anywhere on the computer:



\*   \*\*Linux / macOS\*\*: Move the binary file to `/usr/local/bin/` or `\~/.local/bin/`.

\*   \*\*Windows\*\*: Move the binary file to `C:\\Windows\\System32\\` or add its folder destination to your \*\*Environment Variables\*\* system paths setup dashboard.



\### Option B: Local Relative Path Fallback

If you prefer to lock your runtime dependencies strictly within this directory boundary, place the compiled `llama-cli` binary file directly inside your root workspace folder. 



Then, update the string array declaration inside your `main.go` source file to execute using an explicit local path structure:



```go

// Inside main.go -> runInference() function:

cmd := exec.Command("./llama-cli", // Prepended with './' for direct directory target lookup

&#x20;   "-m", "models/qwen2.5-coder-3b-instruct-q4\_k\_m.gguf",

&#x20;   // ... remaining flags

)

```



\---



\## 🔒 Security \& Sandbox Framework Rules

\*   \*\*Zero-Direct-Action\*\*: The AI shell is explicitly forbidden from running destructive modifications natively.

\*   \*\*Native App Relaying\*\*: Intent files that require modification or creation actions are structured by the model to safely execute inside a system terminal fallback, launching native applications like \*\*Notepad\*\*, \*\*Vim\*\*, or \*\*TextEdit\*\*.



