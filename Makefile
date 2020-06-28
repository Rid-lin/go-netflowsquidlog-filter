.PHONY: build

build: buildwothoutdebug pack

buildfordebug:
	go build -o build/go-netflow2squid -v ./

buildwothoutdebug:
	go build --ldflags "-w -s" -o build/go-netflow2squid -v ./

run: build
	build/go-netflow2squid
	
.DUFAULT_GOAL := build

pack:
	upx --ultra-brute .\go-netflow2squid