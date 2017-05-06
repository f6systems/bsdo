#filename=Makefile
#updatedAt=20170505
#contact=hopley@ipcloud.net
#note=starter Makefile for a github project that can be used in CI for drone golang container
SHELL=/bin/bash

#
bsdo:
	cd ${GOPATH}/src/github.com/f6systems/bsdo/cmd/bsdo && go get && go build -v

