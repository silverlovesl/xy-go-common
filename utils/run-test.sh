#! /bin/bash

go test ./ -coverprofile=./coverage/cover.out
go tool cover -html=./coverage/cover.out -o ./coverage/cover.html