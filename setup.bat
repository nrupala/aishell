@echo off
setlocal enabledelayedexpansion

set MODEL_DIR=models
set MODEL_FILE=Qwen2.5-Coder-3B-Instruct-Q4_K_M.gguf
set MODEL_URL=https://huggingface.co/bartowski/Qwen2.5-Coder-3B-Instruct-GGUF/resolve/main/%MODEL_FILE%

echo =======================================================
echo  AI-Native Shell Environment Setup Script
echo =======================================================

echo == 1. Dependency Check: llama-cli ==
where llama-cli >nul 2>nul
if %ERRORLEVEL% EQU 0 (
    echo [ OK ] llama-cli is available in PATH.
) else (
    if exist "llama-cli.exe" (
        echo [ OK ] llama-cli.exe found in the project root directory.
    ) else (
        echo [ ! ] llama-cli not found in PATH or project root.
        echo       Install llama.cpp and place llama-cli.exe in PATH, or drop it here.
    )
)

echo == 2. Directory and Model Check ==
if not exist "%MODEL_DIR%" (
    mkdir "%MODEL_DIR%"
    echo [ OK ] Created target directory: .\%MODEL_DIR%
)

if exist "%MODEL_DIR%\%MODEL_FILE%" (
    echo [ OK ] Model asset is present. Skipping download.
) else (
    echo [...] Model asset missing. Fetching via PowerShell web client...
    powershell -Command "Invoke-WebRequest -Uri '%MODEL_URL%' -OutFile '%MODEL_DIR%\%MODEL_FILE%'"
    if exist "%MODEL_DIR%\%MODEL_FILE%" (
        echo [ OK ] Model download completed.
    ) else (
        echo [X] Error: model download failed. Check the URL and network.
        exit /b 1
    )
)

echo =======================================================
echo  Environment verification framework complete.
echo  Run build.bat to compile the AI-Native Shell wrapper.
echo =======================================================
pause