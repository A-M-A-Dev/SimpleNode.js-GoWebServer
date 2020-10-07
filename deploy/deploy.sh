#!/bin/bash

$path = $1

echo "Pulling from remote..."
cd $path
git pull

echo "Building and restarting Golang service..."
cd server/go
go build main.go
systemctl restart helloworld-go
cd ../../

echo "Restarting Node.js service"
systemctl restart helloworld-node

echo "Done!"
