#!/usr/bin/env bash

cd ../scraper
echo "==================scraper==================" > ../scripts/test_results.txt  
ls
go test -v ./... >> ../scripts/test_results.txt  

cd ../backend
echo "==================backend==================" >> ../scripts/test_results.txt  
ls >> ../scripts/test_results.txt  

cd ../frontend
echo "==================frontend==================" >> ../scripts/test_results.txt  
ls >> ../scripts/test_results.txt  

cd ../scripts

$SHELL