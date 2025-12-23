@echo off
chcp 65001 > nul
echo Application Status Check...

REM Health Check first
powershell -Command "try { $response = Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/health' -Method GET -TimeoutSec 5; Write-Host 'OK Application Running'; } catch { Write-Host 'ERROR Application not running. Please start shopee-order-go.exe first.'; pause; exit; }"

echo Starting Worker...
powershell -Command "try { $response = Invoke-WebRequest -Uri 'http://localhost:8080/api/v1/scheduler/worker/start' -Method POST; $content = $response.Content | ConvertFrom-Json; Write-Host 'SUCCESS Worker Started: ' $content.message; } catch { Write-Host 'ERROR: ' $_.Exception.Message; }"
pause
