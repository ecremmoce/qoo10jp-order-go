@echo off
chcp 65001 > nul
echo =================================
echo   Worker Status Check
echo =================================

echo 1. Checking Application Health...
powershell -Command "try { $response = Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/health' -Method GET -TimeoutSec 3; Write-Host 'SUCCESS: Application is running'; } catch { Write-Host 'ERROR: Application not running'; exit; }"

echo.
echo 2. Checking Worker Status...
powershell -Command "try { $response = Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/scheduler/status' -Method GET -TimeoutSec 5; $content = $response.Content | ConvertFrom-Json; Write-Host 'Worker Running:' $content.is_running; Write-Host 'Worker Count:' $content.worker_count; Write-Host 'Queue Length:' $content.queue_length; Write-Host 'Last Execution:' $content.last_execution_time; } catch { Write-Host 'ERROR: Failed to get worker status'; }"

echo.
echo 3. Opening detailed status in browser...
start http://localhost:8080/api/v1/scheduler/status

echo.
pause
