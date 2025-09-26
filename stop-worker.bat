@echo off
chcp 65001 > nul
echo Stopping Worker Service...

REM Try API first
powershell -Command "try { $response = Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/scheduler/worker/stop' -Method POST -TimeoutSec 5; $content = $response.Content | ConvertFrom-Json; Write-Host 'SUCCESS: Worker Stopped via API -' $content.message; } catch { Write-Host 'WARNING: API call failed, trying force stop...'; }"

echo.
echo If API failed, use force stop:
echo   1. Press Ctrl+C in the application console window
echo   2. Or run: kill-app.bat
echo   3. Or close the console window directly
echo.
pause
