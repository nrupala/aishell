\## 🛠️ Cross-Platform Build \& Compilation Checklist

Since Go supports out-of-the-box cross-compilation without heavy toolchains, you can build native execution targets for your target environments using these command combinations: \[2, 3] 



&#x20;  1. Build for Windows System Engine

&#x20;  

&#x20;  GOOS=windows GOARCH=amd64 go build -o ai-shell.exe main.go

&#x20;  

&#x20;  2. Build for Linux System Engine

&#x20;  

&#x20;  GOOS=linux GOARCH=amd64 go build -o ai-shell-linux main.go

&#x20;  

&#x20;  3. Build for macOS System Engine (Apple Silicon)

&#x20;  

&#x20;  GOOS=darwin GOARCH=arm64 go build -o ai-shell-mac main.go

&#x20;  

&#x20;  \[4, 5] 



To finalize the operational environment setup, ensure you download your chosen GGUF file (e.g., qwen2.5-coder-3b-instruct-q4\_k\_m.gguf) and place it exactly inside a folder named models located in the working directory of this executable.



How would you like to handle history logging or environment path setups next?



\---------------------------------------------------------------

\#References:

\[1] \[https://fusionauth.io](https://fusionauth.io/docs/get-started/quickstarts/api/quickstart-golang-api)

\[2] \[https://codesyconsulting.com](https://codesyconsulting.com/advantages-and-disadvantages-of-golang)

\[3] \[https://leapcell.io](https://leapcell.io/blog/understanding-the-go-build-command-in-go)

\[4] \[https://www.xa1.at](https://www.xa1.at/build-go/)

\[5] \[https://stackoverflow.com](https://stackoverflow.com/questions/42117934/go-1-7-cross-compilation-from-windows-to-linux-ubuntu)

\----------------------------

\# Compiling Make File: 

On Mac/Linux: Open your terminal and simply type make. 

To compile for all platforms simultaneously, run:

bash



make



\##On Windows: 

Double-click build.bat or execute it from PowerShell/CMD:

powershell

.\\build.bat







Script to check if model exists if llama-cli exists and if there is a dependency that is met. It downloads missing pieces. 



Here is the automated setup script (setup.sh for Unix/Linux/macOS and setup.bat for Windows). \[1] 

This script verifies if llama-cli is present in the system path and checks for the existence of the Qwen2.5-Coder model file in the models/ directory. It skips downloads if they are already present, keeping your environment clean.

\## 1. Unix / Linux / macOS: setup.sh

Create a file named setup.sh in your project root directory. \[2] 



\#!/bin/bash

\# Configuration

MODEL\_DIR="models"

MODEL\_FILE="qwen2.5-coder-3b-instruct-q4\_k\_m.gguf"

MODEL\_URL="https://huggingface.co"



echo "======================================================="

echo " AI-Native Shell Environment Setup Script"

echo "======================================================="

\# 1. Dependency Check: llama-cli

echo "Checking for llama-cli dependency..."if command -v llama-cli \&> /dev/null; then

&#x20;   echo "\[✔] llama-cli is already installed and available in system PATH."elif \[ -f "./llama-cli" ]; then

&#x20;   echo "\[✔] llama-cli found locally in project root directory."else

&#x20;   echo "\[!] llama-cli was not found in PATH or project root."

&#x20;   echo "    Please install llama.cpp or compile llama-cli and place it in your PATH."

&#x20;   echo "    Reference: https://github.com"

&#x20;   # We don't exit here so the user can still download the model ahead of timefi

\# 2. Dependency Check: Directory and Modelif \[ ! -d "$MODEL\_DIR" ]; then

&#x20;   echo "Creating target directory: ./$MODEL\_DIR"

&#x20;   mkdir -p "$MODEL\_DIR"fi



echo "Checking for reasoning model asset..."if \[ -f "$MODEL\_DIR/$MODEL\_FILE" ]; then

&#x20;   echo "\[✔] Model file '$MODEL\_FILE' already exists. Skipping download."else

&#x20;   echo "\[▼] Model file not found. Fetching from Hugging Face repository..."

&#x20;   if command -v curl \&> /dev/null; then

&#x20;       curl -L -o "$MODEL\_DIR/$MODEL\_FILE" "$MODEL\_URL"

&#x20;   elif command -v wget \&> /dev/null; then

&#x20;       wget -O "$MODEL\_DIR/$MODEL\_FILE" "$MODEL\_URL"

&#x20;   else

&#x20;       echo "\[X] Error: Neither curl nor wget was found. Please download the model manually into ./$MODEL\_DIR"

&#x20;       exit 1

&#x20;   fi

&#x20;   

&#x20;   if \[ $? -eq 0 ]; then

&#x20;       echo "\[✔] Model download successfully completed."

&#x20;   else

&#x20;       echo "\[X] Error: Model asset download failed."

&#x20;       exit 1

&#x20;   fifi



echo "======================================================="

echo " Environmental verification framework complete."

echo " Run 'make' to compile your AI-Native Shell wrapper."

echo "======================================================="



Make sure to make it executable before running:



chmod +x setup.sh

./setup.sh



\------------------------------

\## 2. Windows System Engine: setup.bat

Create a file named setup.bat in your project root directory. \[3] 



@echo off

setlocal enabledelayedexpansion



set MODEL\_DIR=models

set MODEL\_FILE=qwen2.5-coder-3b-instruct-q4\_k\_m.gguf

set MODEL\_URL=https://huggingface.co



echo =======================================================

echo  AI-Native Shell Environment Setup Script

echo =======================================================



:: 1. Dependency Check: llama-cli

echo Checking for llama-cli dependency...

where llama-cli >nul 2>nul

if %ERRORLEVEL% EQU 0 (

&#x20;   echo \[✔] llama-cli is already installed and available in system PATH.

) else (

&#x20;   if exist "llama-cli.exe" (

&#x20;       echo \[✔] llama-cli.exe found locally in project root directory.

&#x20;   ) else (

&#x20;       echo \[!] llama-cli was not found in PATH or project root.

&#x20;       echo     Please download the pre-compiled windows binary for llama.cpp 

&#x20;       echo     and drop llama-cli.exe into this folder or your system PATH.

&#x20;   )

)



:: 2. Dependency Check: Directory and Model

if not exist "%MODEL\_DIR%" (

&#x20;   echo Creating target directory: .\\%MODEL\_DIR%

&#x20;   mkdir "%MODEL\_DIR%"

)



echo Checking for reasoning model asset...

if exist "%MODEL\_DIR%\\%MODEL\_FILE%" (

&#x20;   echo \[✔] Model file '%MODEL\_FILE%' already exists. Skipping download.

) else (

&#x20;   echo \[▼] Model file not found. Fetching via PowerShell engine web client...

&#x20;   powershell -Command "Invoke-WebRequest -Uri '%MODEL\_URL%' -OutFile '%MODEL\_DIR%\\%MODEL\_FILE%'"

&#x20;   

&#x20;   if exist "%MODEL\_DIR%\\%MODEL\_FILE%" (

&#x20;       echo \[✔] Model download successfully completed.

&#x20;   ) else (

&#x20;       echo \[X] Error: Model asset download failed.

&#x20;       exit /b 1

&#x20;   )

)



echo =======================================================

echo  Environmental verification framework complete.

echo  Run build.bat to compile your AI-Native Shell wrapper.

echo =======================================================

pause



\------------------------------

\## 🚀 Verification Testing



&#x20;  1. Run the script for your OS platform (./setup.sh or .\\setup.bat).

&#x20;  2. Verify that it checks for llama-cli.

&#x20;  3. Check your models/ directory to see that qwen2.5-coder-3b-instruct-q4\_k\_m.gguf has downloaded successfully.

&#x20;  4. Run the script a second time to verify that it skips the download process altogether and confirms that all assets are ready. \[4, 5] 



If you are ready to configure custom initial profiles or pre-baked commands (aliases) for your shell engine, let me know!



\[1] \[https://github.com](https://github.com/meanmin/sense-claude)

\[2] \[https://cse312.com](https://cse312.com/hw/hw5.html)

\[3] \[https://ducky.ge](https://ducky.ge/docs.html)

\[4] \[https://help.sap.com](https://help.sap.com/doc/d847860d561a47568a936d5f3cbeb9da/4.1/en-US/sapcc\_uis/tool/setuptool.html)

\[5] \[https://software-dl.ti.com](https://software-dl.ti.com/jacinto7/esd/edgeai-sdk-j721e/latest/exports/docs/getting\_started.html)







\-----------------------------------------

Compilation and Setup RemindersPlace this downloaded main.go file inside your working project root directory.Run your workspace automation setup helper script to verify you have the correct dependencies:bash./setup.sh

\# Or on Windows platforms:

.\\setup.bat

Compile the system binaries into place natively via your automation engine files:bashmake

\# Or on Windows platforms:

.\\build.bat





\---------------------------------------







