@echo off
chcp 65001 > nul
echo Force Stopping Qoo10JP Application...

REM Kill process by name
taskkill /f /im qoo10jp-order.exe 2>nul
if %ERRORLEVEL% EQU 0 (
    echo SUCCESS: Application terminated
) else (
    echo INFO: No running application found
)

REM Kill by port 8080
for /f "tokens=5" %%a in ('netstat -aon ^| findstr :8080 ^| findstr LISTENING') do (
    echo Killing process using port 8080: %%a
    taskkill /f /pid %%a 2>nul
)

echo Application stopped.
pause
