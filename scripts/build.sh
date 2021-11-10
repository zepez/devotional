#!/usr/bin/env bash

cd ../scraper
echo "==================scraper==================" > ../logs/build.log  
go build >> ../logs/build.log

cd ../backend
echo "==================backend==================" >> ../logs/build.log  
go build >> ../logs/build.log

cd ../frontend
echo "==================frontend==================" >> ../logs/build.log  
ls >> ../logs/build.log  

cd ../scripts

$SHELL