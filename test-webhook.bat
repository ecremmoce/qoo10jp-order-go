@echo off
echo 웹훅 테스트 시작...

REM 환경변수 로드
if exist env (
    for /f "tokens=1,2 delims==" %%a in (env) do (
        if not "%%a"=="" if not "%%a:~0,1%"=="#" (
            set "%%a=%%b"
        )
    )
)

REM Go 테스트 스크립트 실행
go run scripts/test-webhook.go

echo.
echo 웹훅 테스트 완료
pause










