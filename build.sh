#!/bin/bash

GOOS=linux GOARCH=amd64  go build api/api.go
GOOS=linux GOARCH=amd64  go build loginsvc/loginsvc.go
GOOS=linux GOARCH=amd64  go build usersvc/usersvc.go