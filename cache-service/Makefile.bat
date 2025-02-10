@echo off
setlocal enabledelayedexpansion

set GOPATH=%GOPATH%
set PATH=%PATH%;%GOPATH%\bin

:: 设置绝对路径
set PROTO_DIR=%CD%\..\proto\cache
set INCLUDE_DIR=%CD%\..\comm\include

for %%F in (%PROTO_DIR%\*.proto) do (
    echo %%F
    protoc -I%INCLUDE_DIR% --plugin=protoc-gen-go=%GOPATH%\bin\protoc-gen-go.exe --plugin=protoc-gen-micro=%GOPATH%\bin\protoc-gen-micro.exe --proto_path=%CD%\.. --micro_out=%PROTO_DIR% --go_out=%PROTO_DIR% --validate_out=lang=go:%PROTO_DIR% %%F
)
:: 移除 ,omitempty
for %%F in (%PROTO_DIR%\*.go) do (
    powershell -Command "(Get-Content -Path '%%F') -replace ',omitempty', '' | Set-Content -Path '%%F'"
)

endlocal
