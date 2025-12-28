# ===================================================
# Qoo10JP Order Go - One-Click Deploy Script (PowerShell)
# ===================================================

param(
    [string]$Server = "192.168.0.203",
    [string]$User = "james",
    [string]$Action = "deploy"  # deploy, start, stop, restart, status, logs
)

$AppName = "qoo10jp-order-go"
$RemotePath = "/home/$User/qoo10jp-order-go"
$Port = "8082"  # Shopee: 8080, Lazada: 8081, Qoo10JP: 8082
$SSH_OPTS = "-o ServerAliveInterval=30 -o ServerAliveCountMax=3 -o ConnectTimeout=10"

function Build-Linux {
    Write-Host "[BUILD] Building Linux binary for Qoo10JP..." -ForegroundColor Cyan
    $env:GOOS = "linux"
    $env:GOARCH = "amd64"
    $env:CGO_ENABLED = "0"
    
    # 프로젝트 루트로 이동
    $scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
    $projectRoot = Split-Path -Parent $scriptPath
    Push-Location $projectRoot
    
    go build -a -installsuffix cgo -o qoo10jp-order-go-linux ./cmd/main.go
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "[OK] Build successful!" -ForegroundColor Green
    } else {
        Write-Host "[ERROR] Build failed!" -ForegroundColor Red
        Pop-Location
        exit 1
    }
    Pop-Location
}

function Deploy-App {
    Write-Host "[DEPLOY] Uploading files to $Server..." -ForegroundColor Cyan
    
    $scriptPath = Split-Path -Parent $MyInvocation.MyCommand.Path
    $projectRoot = Split-Path -Parent $scriptPath
    
    # Upload files
    scp $SSH_OPTS "$projectRoot\qoo10jp-order-go-linux" "${User}@${Server}:/tmp/qoo10jp-order-go"
    scp $SSH_OPTS "$projectRoot\env" "${User}@${Server}:/tmp/qoo10jp.env"
    scp $SSH_OPTS -r "$projectRoot\web" "${User}@${Server}:/tmp/qoo10jp-web"
    
    Write-Host "[DEPLOY] Installing on server..." -ForegroundColor Cyan
    
    # Stop existing process and install new files
    $installScript = @"
pkill -f qoo10jp-order-go || true
mkdir -p $RemotePath
mv /tmp/qoo10jp-order-go $RemotePath/
mv /tmp/qoo10jp.env $RemotePath/.env
rm -rf $RemotePath/web
mv /tmp/qoo10jp-web $RemotePath/web
chmod +x $RemotePath/qoo10jp-order-go

# 포트를 8081로 설정
sed -i 's/PORT=8080/PORT=$Port/' $RemotePath/.env || true

cat > $RemotePath/start.sh << 'STARTSCRIPT'
#!/bin/bash
cd /home/james/qoo10jp-order-go
export `$(cat .env | grep -v '^#' | xargs)
./qoo10jp-order-go
STARTSCRIPT
chmod +x $RemotePath/start.sh
"@
    
    ssh $SSH_OPTS "${User}@${Server}" $installScript
    
    Write-Host "[OK] Deploy complete!" -ForegroundColor Green
}

function Start-App {
    Write-Host "[START] Starting Qoo10JP application on port $Port..." -ForegroundColor Cyan
    # Use screen instead of nohup to avoid SSH session hanging (Cursor Connection Error)
    ssh $SSH_OPTS "${User}@${Server}" "screen -ls | grep qoo10jp && screen -S qoo10jp -X quit; cd $RemotePath && screen -dmS qoo10jp ./start.sh; sleep 2; screen -ls | grep qoo10jp"
    Start-Sleep -Seconds 2
    Get-Status
}

function Stop-App {
    Write-Host "[STOP] Stopping Qoo10JP application..." -ForegroundColor Cyan
    # Stop screen session and process
    ssh $SSH_OPTS "${User}@${Server}" "screen -S qoo10jp -X quit 2>/dev/null; pkill -f qoo10jp-order-go || true"
    Write-Host "[OK] Stopped!" -ForegroundColor Green
}

function Restart-App {
    Stop-App
    Start-Sleep -Seconds 2
    Start-App
}

function Get-Status {
    Write-Host "[STATUS] Checking Qoo10JP status..." -ForegroundColor Cyan
    $process = ssh $SSH_OPTS "${User}@${Server}" "ps aux | grep qoo10jp-order-go | grep -v grep"
    
    if ($process) {
        Write-Host "[OK] Running" -ForegroundColor Green
        Write-Host $process
        
        # API health check
        try {
            $health = Invoke-WebRequest -Uri "http://${Server}:${Port}/api/v1/health" -UseBasicParsing -TimeoutSec 5
            Write-Host "[OK] API OK: $($health.Content)" -ForegroundColor Green
        } catch {
            Write-Host "[WARN] API not responding on port $Port" -ForegroundColor Yellow
        }
    } else {
        Write-Host "[ERROR] Not running" -ForegroundColor Red
    }
}

function Get-Logs {
    param([int]$Lines = 100)
    Write-Host "[LOGS] Recent Qoo10JP logs ($Lines lines)..." -ForegroundColor Cyan
    ssh $SSH_OPTS "${User}@${Server}" "tail -$Lines $RemotePath/qoo10jp-order-go.log"
}

function Show-BothStatus {
    Write-Host "`n========== Shopee Order Go (Port 8080) ==========" -ForegroundColor Magenta
    $shopeeProcess = ssh $SSH_OPTS "${User}@${Server}" "ps aux | grep shopee-order-go | grep -v grep"
    if ($shopeeProcess) {
        Write-Host "[OK] Running" -ForegroundColor Green
    } else {
        Write-Host "[--] Not running" -ForegroundColor Gray
    }
    
    Write-Host "`n========== Qoo10JP Order Go (Port $Port) ==========" -ForegroundColor Cyan
    Get-Status
}

# Main
switch ($Action) {
    "deploy" {
        Build-Linux
        Deploy-App
        Start-App
    }
    "build" { Build-Linux }
    "start" { Start-App }
    "stop" { Stop-App }
    "restart" { Restart-App }
    "status" { Get-Status }
    "all-status" { Show-BothStatus }
    "logs" { Get-Logs }
    default {
        Write-Host "Qoo10JP Order Go - Deploy Script" -ForegroundColor Cyan
        Write-Host "================================" -ForegroundColor Cyan
        Write-Host ""
        Write-Host "Usage: .\deploy.ps1 -Action [command]" -ForegroundColor Yellow
        Write-Host ""
        Write-Host "Commands:" -ForegroundColor White
        Write-Host "  deploy     - Build, upload, and start" -ForegroundColor Gray
        Write-Host "  build      - Build Linux binary only" -ForegroundColor Gray
        Write-Host "  start      - Start application" -ForegroundColor Gray
        Write-Host "  stop       - Stop application" -ForegroundColor Gray
        Write-Host "  restart    - Restart application" -ForegroundColor Gray
        Write-Host "  status     - Check status" -ForegroundColor Gray
        Write-Host "  all-status - Check both Shopee and Qoo10JP" -ForegroundColor Gray
        Write-Host "  logs       - View recent logs" -ForegroundColor Gray
    }
}
