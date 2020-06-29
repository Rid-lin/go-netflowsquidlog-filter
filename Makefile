.PHONY: build

build: buildwothoutdebug_linux pack

all: buildwothoutdebug buildwothoutdebug_linux pack

buildfordebug:
	go build -o build/go-netflowsquidlog-filter -v ./

buildwothoutdebug:
	go build --ldflags "-w -s" -o build/go-netflowsquidlog-filter.exe -v ./

buildwothoutdebug_linux:
	set GOOS=linux&& go build --ldflags "-w -s" -o build/go-netflowsquidlog-filter -v ./

run: build
	build/go-netflowsquidlog-filter
	
.DUFAULT_GOAL := build

pack:
	upx --ultra-brute build\go-netflowsquidlog-filter*