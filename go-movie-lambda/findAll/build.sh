#!/bin/sh

echo "Build the binary"
GOARCH=amd64 GOOS=linux go build -o app findAll.go
echo "Create a zip archive for deployment"
zip a deployment.zip app
echo "Cleaning up"
rm app
