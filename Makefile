all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64 windows-amd64 windows-arm64
linux-amd64:
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && go build  -ldflags "-w -s " -o ./bin/git-plus-linux-amd64 main.go
linux-arm64:
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=arm64 && go build  -ldflags "-w -s " -o ./bin/git-plus-linux-arm64 main.go
darwin-amd64:
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && go build  -ldflags "-w -s " -o ./bin/git-plus-darwin-amd64 main.go
darwin-arm64:
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=arm64 && go build  -ldflags "-w -s " -o ./bin/git-plus-darwin-arm64 main.go
windows-amd64:
	export CGO_ENABLED=0 && export GOOS=windows && export GOARCH=amd64 && go build  -ldflags "-w -s " -o ./bin/git-plus-windows-amd64 main.go
windows-arm64:
	export CGO_ENABLED=0 && export GOOS=windows && export GOARCH=arm64 && go build  -ldflags "-w -s " -o ./bin/git-plus-windows-arm64 main.go