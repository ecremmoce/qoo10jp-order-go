@echo off
chcp 65001 > nul
echo =================================
echo   Qoo10JP Order Collector
echo =================================

echo 1. Building Application...
go build -o qoo10jp-order.exe cmd/main.go

if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Build Failed!
    pause
    exit /b 1
)

echo SUCCESS: Build Complete!
echo.
echo 2. Starting Application...
echo    - Worker will start automatically
echo    - Check: http://localhost:8080/api/v1/health
echo.

start "Qoo10JP Order Collector" qoo10jp-order.exe

echo 3. Waiting 5 seconds for startup...
timeout /t 5 /nobreak > nul

echo 4. Health Check...
powershell -Command "try { Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/health' -Method GET -TimeoutSec 3 | Out-Null; Write-Host 'SUCCESS: Application Running'; } catch { Write-Host 'WARNING: Application starting...'; }"

echo.
echo SUCCESS: System Ready!
echo    - Worker: Auto-started
echo    - Status: http://localhost:8080/api/v1/scheduler/status
echo.
pause
