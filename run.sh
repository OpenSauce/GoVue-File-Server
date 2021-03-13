#!bin/bash

echo "Building Go server..."
go build main/

echo "Building Vue frontend..."
cd frontend
npm install
npm run build
cd ..

echo "Running..."
go run main/server.go
