Set-Location ..

$Env:CGO_ENABLED="1"
$Env:GOROOT_FINAL="/usr"

$Env:GOOS="windows"
$Env:GOARCH="amd64"
go build -a -trimpath -buildmode=c-shared -asmflags "-s -w" -ldflags "-s -w" -gcflags "-l=4" -o aiodns.bin
if (! $?) { exit 1 }

upx --ultra-brute aiodns.bin

Set-Location scripts
exit 0
