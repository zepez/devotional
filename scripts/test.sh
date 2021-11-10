#!/usr/bin/env bash

cd ../scraper
echo "==================scraper==================" > ../scripts/results.txt  
ls
go test -v ./... >> ../scripts/results.txt  

cd ../backend
echo "==================backend==================" >> ../scripts/results.txt  
ls >> ../scripts/results.txt  

cd ../frontend
echo "==================frontend==================" >> ../scripts/results.txt  
ls >> ../scripts/results.txt  

cd ../scripts

$SHELL