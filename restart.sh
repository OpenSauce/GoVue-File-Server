#!bin/bash

pkill -9 main

echo "Running..."
cd main
go build
./main & 