#!/usr/bin/bash
gofmt -s -w .
cat .env | awk '{split($0,a,"="); print a[1] "="}' > .env-template
