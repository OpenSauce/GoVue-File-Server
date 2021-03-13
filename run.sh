#!bin/bash

git pull

pkill -9 main

echo "Building Vue frontend..."
cd frontend
npm install
npm run build
cd ..

echo "Running..."
cd main
go build
./main & 