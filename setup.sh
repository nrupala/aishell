#!/bin/sh
# AI-Native Shell environment setup for Unix / Linux / macOS
# Verifies llama-cli, creates the models directory, and downloads the
# Qwen2.5-Coder-3B-Instruct Q4_K_M GGUF asset if it is not already present.

MODEL_DIR="models"
MODEL_FILE="Qwen2.5-Coder-3B-Instruct-Q4_K_M.gguf"
# The binary loads this exact (lowercase) path — see modelPath in main.go.
# On case-sensitive filesystems the Hugging Face filename does NOT match it,
# so the download is saved under the name the binary actually opens.
LOCAL_MODEL_FILE="qwen2.5-coder-3b-instruct-q4_k_m.gguf"
MODEL_URL="https://huggingface.co/bartowski/Qwen2.5-Coder-3B-Instruct-GGUF/resolve/main/${MODEL_FILE}"

echo "======================================================="
echo " AI-Native Shell Environment Setup Script"
echo "======================================================="

echo "== 1. Dependency Check: llama-cli =="
if command -v llama-cli >/dev/null 2>&1; then
    echo "[ OK ] llama-cli is available in PATH."
elif [ -f "./llama-cli" ]; then
    echo "[ OK ] llama-cli found in the project root directory."
else
    echo "[ ! ] llama-cli not found in PATH or project root."
    echo "      Install llama.cpp and place llama-cli in PATH, or drop it here."
fi

echo "== 2. Directory and Model Check =="
if [ ! -d "$MODEL_DIR" ]; then
    mkdir -p "$MODEL_DIR"
    echo "[ OK ] Created target directory: ./$MODEL_DIR"
fi

if [ -f "$MODEL_DIR/$LOCAL_MODEL_FILE" ]; then
    echo "[ OK ] Model asset is present. Skipping download."
else
    echo "[...] Model asset missing. Fetching from Hugging Face..."
    if command -v curl >/dev/null 2>&1; then
        curl -L -o "$MODEL_DIR/$LOCAL_MODEL_FILE" "$MODEL_URL"
    elif command -v wget >/dev/null 2>&1; then
        wget -O "$MODEL_DIR/$LOCAL_MODEL_FILE" "$MODEL_URL"
    else
        echo "[X] Error: neither curl nor wget was found. Download manually into ./$MODEL_DIR"
        exit 1
    fi

    if [ -f "$MODEL_DIR/$LOCAL_MODEL_FILE" ]; then
        echo "[ OK ] Model download completed."
    else
        echo "[X] Error: model download failed. Check the URL and network."
        exit 1
    fi
fi

echo "======================================================="
echo " Environment verification framework complete."
echo " Run 'make' (Unix) or build.bat (Windows) to compile."
echo "======================================================="