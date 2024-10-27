@echo off
mkdir dist
mkdir dist\windows
mkdir dist\linux
mkdir dist\macos

echo Building for Windows...
set GOOS=windows
set GOARCH=amd64
go build -o dist\windows\torri.exe main.go
powershell Compress-Archive -Path dist\windows\torri.exe -DestinationPath dist\torri-windows.zip

echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o dist\linux\torri main.go
powershell Compress-Archive -Path dist\linux\torri -DestinationPath dist\torri-linux.zip

echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o dist\macos\torri main.go
powershell Compress-Archive -Path dist\macos\torri -DestinationPath dist\torri-macos.zip

echo Cleaning up intermediate files...
rmdir /s /q dist\windows
rmdir /s /q dist\linux
rmdir /s /q dist\macos

echo Build and packaging complete. Files are in the 'dist' folder.
