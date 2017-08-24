@echo off

SET GOPATH="%~dp0..\.."
go build -o tool-debug.exe
