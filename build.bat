@echo off
SET BINARY_NAME=ai-shell
SET BUILD_FLAGS=-ldflags="-s -w"

echo =======================================================
echo  AI-Native Shell Automated Build Engine
echo =======================================================

:: Create target model infrastructure
if not exist "models" (
    mkdir models
    echo [ OK ] Created models folder.
)

echo.
echo Building for Windows (AMD64)...
set GOOS=windows
set GOARCH=amd64
go build %BUILD_FLAGS% -o %BINARY_NAME%.exe .
if %ERRORLEVEL% EQU 0 (echo [ OK ] Windows build complete: %BINARY_NAME%.exe) else (echo [ FAIL ] Windows build failed)

echo.
echo Building for Linux (AMD64)...
set GOOS=linux
set GOARCH=amd64
go build %BUILD_FLAGS% -o %BINARY_NAME%-linux .
if %ERRORLEVEL% EQU 0 (echo [ OK ] Linux build complete: %BINARY_NAME%-linux) else (echo [ FAIL ] Linux build failed)

echo.
echo Building for macOS (Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build %BUILD_FLAGS% -o %BINARY_NAME%-mac .
if %ERRORLEVEL% EQU 0 (echo [ OK ] macOS build complete: %BINARY_NAME%-mac) else (echo [ FAIL ] macOS build failed)

echo.
echo =======================================================
echo Compilation suite processed.
pause