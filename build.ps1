$Env:CGO_ENABLED="1"
$Env:GOROOT_FINAL="/usr"

$Env:GOOS="windows"
$Env:GOARCH="amd64"
go build -a -trimpath -buildmode=c-shared -asmflags "-s -w" -ldflags "-s -w" -o aiodns.bin
if (-Not $?) { exit $lastExitCode }

upx --ultra-brute aiodns.bin
exit $lastExitCode
