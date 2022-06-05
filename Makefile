GOROOT= /usr/local/go
GOPATH = $(HOME)/go
GOBIN = $(GOPATH)/bin

help:
	# Development commands:
	# make build_simple - build executable for your arch and OS	
	#
	# Testing commands:
	# make test - run project tests
	# make cover - run project tests save result 
	#
	# Deployment commands:
	# make install_golang - install golang#

build_simple:
	go build

cover:	
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

test:
	go test -cover -v ./...

mysql:
	sudo docker-compose -f docker/mysql/docker-compose.yaml up --build --abort-on-container-exit

install_golang:	
	wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
	mkdir -p ~/go/{bin,pkg,src}
	tar -C /usr/local/ -xzf go1.16.7.linux-amd64.tar.gz
	rm go1.16.7.linux-amd64.tar.gz
