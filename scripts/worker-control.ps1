# Worker Control Script
param(
    [Parameter(Mandatory=$true)]
    [ValidateSet("start", "stop", "status", "restart")]
    [string]$Action
)

$BaseUrl = "http://localhost:8080/api/v1/scheduler"

function Invoke-WorkerAPI {
    param($Endpoint, $Method = "GET")
    
    try {
        $response = Invoke-WebRequest -Uri "$BaseUrl/$Endpoint" -Method $Method -ErrorAction Stop
        $content = $response.Content | ConvertFrom-Json
        return $content
    }
    catch {
        Write-Error "API 호출 실패: $($_.Exception.Message)"
        return $null
    }
}

switch ($Action) {
    "start" {
        Write-Host "Worker 시작 중..." -ForegroundColor Yellow
        $result = Invoke-WorkerAPI "worker/start" "POST"
        if ($result) {
            Write-Host "✅ Worker 시작됨: $($result.message)" -ForegroundColor Green
        }
    }
    
    "stop" {
        Write-Host "Worker 중지 중..." -ForegroundColor Yellow
        $result = Invoke-WorkerAPI "worker/stop" "POST"
        if ($result) {
            Write-Host "⏹️ Worker 중지됨: $($result.message)" -ForegroundColor Red
        }
    }
    
    "status" {
        Write-Host "Worker 상태 확인 중..." -ForegroundColor Yellow
        $result = Invoke-WorkerAPI "status"
        if ($result) {
            Write-Host "📊 Worker 상태:" -ForegroundColor Cyan
            Write-Host "  - 실행 중: $($result.is_running)" -ForegroundColor White
            Write-Host "  - Worker 수: $($result.worker_count)" -ForegroundColor White
            Write-Host "  - Queue 길이: $($result.queue_length)" -ForegroundColor White
            Write-Host "  - 마지막 실행: $($result.last_execution_time)" -ForegroundColor White
        }
    }
    
    "restart" {
        Write-Host "Worker 재시작 중..." -ForegroundColor Yellow
        Invoke-WorkerAPI "worker/stop" "POST" | Out-Null
        Start-Sleep -Seconds 2
        $result = Invoke-WorkerAPI "worker/start" "POST"
        if ($result) {
            Write-Host "🔄 Worker 재시작됨: $($result.message)" -ForegroundColor Green
        }
    }
}
