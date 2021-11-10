#!/usr/bin/env bash

cd ../scraper
echo "==================scraper==================" > ../logs/test.log  
go test -v ./... >> ../logs/test.log  

cd ../backend
echo "==================backend==================" >> ../logs/test.log  
go test -v ./... >> ../logs/test.log  

cd ../frontend
echo "==================frontend==================" >> ../logs/test.log  
ls >> ../logs/test.log  

cd ../scripts

$SHELL