#!/bin/bash

echo "Pulling from remote..."
cd $1
git pull

echo "Building and restarting Golang service..."
cd server/go
/home/ahmad/.go/bin/go build main.go
systemctl restart helloworld-go
cd ../../

echo "Restarting Node.js service"
systemctl restart helloworld-node

echo "Done!"
