#filename=Makefile
#updatedAt=20170505
#contact=hopley@ipcloud.net
#note=starter Makefile for a github project that can be used in CI for drone golang container
SHELL=/bin/sh

#
bsdo:
	cd /go/src/github.com/f6systems/bsdo/cmd/bsdo
	go build -v

