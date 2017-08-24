@echo off

SET GOPATH="%~dp0..\.."
go build -ldflags "-w -s" -o tool-release.exe
