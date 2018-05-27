#!/bin/bash

# CONTROLLER
env GOOS=windows GOARCH=amd64 go build -o releases/dnsrrResolver-win-amd64.exe 
env GOOS=windows GOARCH=386 go build -o releases/dnsrrResolver-win-i386.exe 

env GOOS=linux GOARCH=386 go build -o releases/dnsrrResolver-linux-i386 
env GOOS=linux GOARCH=amd64 go build -o releases/dnsrrResolver-linux-amd64 

env GOOS=darwin GOARCH=amd64 go build -o releases/dnsrrResolver-osx-amd64 