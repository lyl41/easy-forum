#!/bin/bash

fmt:
	goimports -l -w .

dev: fmt
	go build -o output/bin/forum-MacOs .

linux: fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin/forum-server .

deploy: fmt linux
	tar -czvf output/bin/forum-server.tar.gz output/bin/forum-server
	scp output/bin/forum-server.tar.gz root@108.160.139.109:/root/forum
	rm -f output/bin/forum-server.tar.gz
