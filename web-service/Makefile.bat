@echo off
set currentDir=%cd%

for /d %%d in (%currentDir%\*) do (
    if exist "%%d\Makefile.bat" (
        pushd %%d
        call Makefile.bat
        popd
    )
)

cd %currentDir%

