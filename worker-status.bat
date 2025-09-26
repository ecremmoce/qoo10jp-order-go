@echo off
chcp 65001 > nul
echo Checking Worker Status...
echo Opening browser: http://localhost:8080/api/v1/scheduler/status
start http://localhost:8080/api/v1/scheduler/status
