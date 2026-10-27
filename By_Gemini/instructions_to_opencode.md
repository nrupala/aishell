Hey OpenCode, you are compiling a highly optimized, cross-platform, AI-native shell wrapper written in Go. The goal is to build this with absolute structural hygiene. 



Please adhere closely to these three critical implementation details during code generation:



1\. STREAMING ROBUSTNESS: Ensure the stdout reader loop from the background `llama-cli` process is highly resilient. It must capture the <think> and </think> delimiters dynamically, printing the model's reasoning chain in real-time, and cleanly isolate the markdown fences (```cmd, ```bash, etc.) to extract the payload without syntax edge cases.



2\. PURE OS ISOLATION: Do not use Docker or external container software. Rely completely on the host system execution wrappers mapped out in the code (systemd-run on Linux with its explicit ProtectSystem/ProtectHome flags, and sandbox-exec with the Seatbelt profile on macOS). Ensure Go's os/exec package safely delegates Stdin, Stdout, and Stderr directly to the sandbox subprocess so interactive tools like 'less' or 'vim' remain responsive inside the shell loop.



3\. GO DESCRIPTOR SANITY: Maintain the platform-specific logging logic. On Windows, use local ISO-8601 timestamps and write to %APPDATA%, while on POSIX platforms, write true integer Unix Epoch timestamps to standard XDG data directory paths (\~/.local/share/).



Review the provided blueprint files carefully, generate clean, standard library-focused Go code, and ensure it passes the verification test plan on all target architectures. Let's build it!



