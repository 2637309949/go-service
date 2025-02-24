@echo off
setlocal enabledelayedexpansion

set GOPATH=%GOPATH%
set PATH=%PATH%;%GOPATH%\bin

set INSTALL_PATH=C:\\go-micro
set SERVICE_NAME=upload.exe
set CMD= 

:: 设置绝对路径
set PROTO_DIR=%CD%\..\..\proto\web\upload
set INCLUDE_DIR=%CD%\..\..\comm\include

for %%F in (%PROTO_DIR%\*.proto) do (
    echo %%F
    protoc -I%INCLUDE_DIR% --plugin=protoc-gen-web=%GOPATH%\bin\protoc-gen-micro.exe --proto_path=%CD%\..\.. --web_out=plugins=web:%PROTO_DIR% %%F
)
:: 移除 ,omitempty
for %%F in (%PROTO_DIR%\*.go) do (
    powershell -Command "(Get-Content -Path '%%F') -replace ',omitempty', '' | Set-Content -Path '%%F'"
)

:: 构建执行文件
echo build %SERVICE_NAME%
go build

:: 执行命令
if "%1"=="install" (
    set CMD=install
)
if "%1"=="up" (
    set CMD=install
)
if "%1"=="release" (
    set CMD=release
)
if "%CMD%"=="release" (
    echo mkdir %INSTALL_PATH%
    mkdir "%INSTALL_PATH%"
    echo copy %SERVICE_NAME% %INSTALL_PATH%
    copy /y "%SERVICE_NAME%" "%INSTALL_PATH%"
)
if "%CMD%"=="install" (
    echo mkdir %INSTALL_PATH%
    mkdir "%INSTALL_PATH%"
    echo copy %SERVICE_NAME% %INSTALL_PATH%
    copy /y "%SERVICE_NAME%" "%INSTALL_PATH%"
    set FOUND_OLD_PROCESS=false
    tasklist /fi "imagename eq %SERVICE_NAME%" /nh | find /i "%SERVICE_NAME%" >nul
    if %errorlevel% equ 0 (
        for /f "tokens=1,2,* delims= " %%a in ('tasklist /fi "imagename eq %SERVICE_NAME%" /nh') do (
            set OLD_PID=%%a
            echo kill !OLD_PID!
            taskkill /f /pid !OLD_PID!
            echo start %SERVICE_NAME%
            copy /y "%SERVICE_NAME%" "%INSTALL_PATH%"
            start "" powershell -Command "Start-Process '%INSTALL_PATH%\%SERVICE_NAME%' -WindowStyle Hidden"
            set FOUND_OLD_PROCESS=true
        )
    )
    if !FOUND_OLD_PROCESS! == false (
        echo start %SERVICE_NAME%
        start "" powershell -Command "Start-Process '%INSTALL_PATH%\%SERVICE_NAME%' -WindowStyle Hidden"
    )
)
endlocal
