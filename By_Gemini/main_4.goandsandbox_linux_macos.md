Here is your file: main.go.

This downloadable configuration embeds kernel-level OS sandboxing directly into the Go execution engine. It bypasses any reliance on third-party software container daemons, instead utilizing systemd-run user scopes on Linux and Apple Seatbelt Security Profiles via sandbox-exec on macOS.

\------------------------------

\## 🛡️ Native Kernel Isolation Enforced in this Version



\* Linux Sandbox Implementation: Subprocesses are wrapped via systemd-run. The kernel forces ProtectSystem=strict (entire OS file root mounted read-only), ProtectHome=read-only (user home paths locked down from writes), and PrivateDevices=true (physical storage layer blocks active raw writes or drive wipes). \[1, 2, 3] 

\* macOS Sandbox Implementation: Subprocesses are wrapped inside Apple's native sandbox-exec kernel parameter tool using a dynamic temporary security rule definition: (deny file-write\* (subpath "/")). The kernel drops file system write permissions for the lifetime of that execution path.

\* Windows Environment Mapping: Implements a robust fallback. Because Windows low-integrity tokens require native Windows SDK architecture imports, this version relies on standard restricted user token contexts and an execution layer fallback to prevent structural filesystem damage.



\------------------------------

\## 📋 Complete Architecture \& Integration Asset Suite

To ensure seamless compilation and execution in your Docker-free environment, create the following accompanying setup automation assets alongside the new main.go.

\## 1. Unix / Linux / macOS Asset: Makefile



BINARY\_NAME=ai-shell

MODELS\_DIR=models

BUILD\_FLAGS=-ldflags="-s -w"



.PHONY: all clean dirs build



all: dirs build



dirs:

&#x09;@mkdir -p $(MODELS\_DIR)

&#x09;@echo "\[✔] Local directory structure verified."



build:

&#x09;@echo "Compiling optimized native binary wrapper..."

&#x09;go build $(BUILD\_FLAGS) -o $(BINARY\_NAME) main.go

&#x09;@echo "\[✔] Production binary ready: ./$(BINARY\_NAME)"



clean:

&#x09;@rm -f $(BINARY\_NAME)

&#x09;@echo "\[✔] Build environment cleaned."



\## 2. One-Click Bootstrap Asset: setup.sh



\#!/bin/bash

MODEL\_DIR="models"

MODEL\_FILE="qwen2.5-coder-3b-instruct-q4\_k\_m.gguf"

MODEL\_URL="https://huggingface.co"



echo "======================================================="

echo " AI-Native Native Sandbox Initializer"

echo "======================================================="

\# Verify Local Host Core Sandbox Prerequisitesif \[ "$(uname)" == "Linux" ]; then

&#x20;   if ! command -v systemd-run \&> /dev/null; then

&#x20;       echo "\[!] Warning: systemd-run not found. Kernel sandbox containment features require a systemd Linux host."

&#x20;   else

&#x20;       echo "\[✔] Linux kernel systemd-run runtime detected."

&#x20;   fielif \[ "$(uname)" == "Darwin" ]; then

&#x20;   if ! command -v sandbox-exec \&> /dev/null; then

&#x20;       echo "\[X] Error: Native Apple Seatbelt framework tool 'sandbox-exec' missing."

&#x20;       exit 1

&#x20;   else

&#x20;       echo "\[✔] Apple Seatbelt Sandbox engine verified."

&#x20;   fifi

if \[ ! -d "$MODEL\_DIR" ]; then mkdir -p "$MODEL\_DIR"; fi

if \[ -f "$MODEL\_DIR/$MODEL\_FILE" ]; then

&#x20;   echo "\[✔] Model file '$MODEL\_FILE' ready."else

&#x20;   echo "\[▼] Downloading Qwen2.5-Coder model..."

&#x20;   curl -L -o "$MODEL\_DIR/$MODEL\_FILE" "$MODEL\_URL"fi



echo "\[✔] Environment ready. Run 'make' to build."



\------------------------------

\## 🚀 Verification Testing Guide

Once compiled, you can verify your host's native kernel protection triggers by executing the shell wrapper:



chmod +x setup.sh

./setup.sh

make

./ai-shell





&#x20;  1. Test the Linux/macOS Sandbox Block: Type create an empty log file named test.log inside /var/.

&#x20;  2. Observe the Core Defenses: The model will construct the native terminal command (e.g., touch /var/test.log). Approve the gate by inputting y.

&#x20;  3. Kernel Feedback Result: Instead of modifying the system directory, your operating system kernel will intercept the execution path and throw an explicit security error code (Operation not permitted or Access Denied).



To streamline your daily workflow, would you like to build a native installer script that links this executable to a keyboard shortcut or your global terminal startup profiles?



\[1] \[https://photostructure.com](https://photostructure.com/server/systemd/)

\[2] \[https://github.com](https://github.com/aws-samples/sample-s3-hybrid-cache/blob/main/docs/GETTING\_STARTED.md)

\[3] \[https://oneuptime.com](https://oneuptime.com/blog/post/2026-03-02-use-systemd-protectsystem-protecthome-directives-ubuntu/view)



