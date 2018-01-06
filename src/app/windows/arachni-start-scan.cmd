@echo off
..\tool-windows-amd64-release.exe tool-windows-amd64-debug.exe arachni scan-start -s "http://192.168.99.100:7331" -u arachni -p Pass123 %*

