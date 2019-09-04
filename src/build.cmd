@echo off

SET GOPATH="%~dp0..\.."

IF "%GOOS%"=="" SET GOOS=windows
IF "%GOARCH%"=="" SET GOARCH=amd64
SET VER=-debug
SET OPT=
IF "%BUILD%"=="RELEASE" SET VER=-release
IF "%BUILD%"=="RELEASE" SET OPT=-ldflags "-w -s"
IF "%BUILD%"=="RELEASE" SET UPX_EXE=upx.exe

SET EXT=.exe
IF not "%GOOS%"=="windows" SET EXT=

SET ARM=
IF not "%GOARM%"=="" SET ARM=-arm%GOARM%

SET OUT_EXE=tool-%GOOS%-%GOARCH%%VER%%ARM%%EXT%

echo Build "%OUT_EXE%" ...

go build %OPT% -o %OUT_EXE%

if exist %UPX_EXE% %UPX_EXE% -9 %OUT_EXE%
