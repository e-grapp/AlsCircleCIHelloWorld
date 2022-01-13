.PHONY: default

BINARY=ksplit

GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GODOC=$(GOCMD) doc
GOGET=$(GOCMD) get

export TZ=Asia/Shanghai
export PACK=main
export FLAGS="-s -w -X '$(PACK).AppName=$(BINARY)' -X '$(PACK).BuildDate=`date '+%Y-%m-%dT%T%z'`' -X '$(PACK).BuildHost=`hostname`' -X '$(PACK).GoVersion=`go version`' -X '$(PACK).GitBranch=`git symbolic-ref -q --short HEAD`' -X '$(PACK).GitCommit=`git rev-parse --short HEAD`' -X '$(PACK).GitSummary=`git describe --tags --dirty --always`' -X '$(PACK).CIBuildNum=${CIRCLE_BUILD_NUM}'"

default:
	@echo "build target is required for $(BINARY)"
	@exit 1
doc:
	$(GODOC) -all ...
fmt:
	gofmt -l -w .
	goimports -l -w .
	goreturns -l -w .
build:
	$(GOBUILD) -v -ldflags $(FLAGS) -trimpath -o $(BINARY)
build_windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -v -ldflags $(FLAGS) -trimpath -o $(BINARY).exe
build_mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -v -ldflags $(FLAGS) -trimpath -o $(BINARY)
build_linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -ldflags $(FLAGS) -trimpath -o $(BINARY)
install: build
	mv $(BINARY) $$GOPATH/bin
run:
	$(GORUN) -v .
test:
	$(GOTEST) -race -cover -covermode=atomic -v -count 1 .
bench:
	$(GOTEST) -parallel=4 -run="none" -benchtime="5s" -benchmem -bench=.
