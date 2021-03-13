#!bin/bash


echo "Building Vue frontend..."
cd frontend
npm install
npm run build
cd ..

echo "Running..."
go build ./main
./main/main