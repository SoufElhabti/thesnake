#!/bin/bash

filename=$1
n=1
while read line; do 
python3 linkfinder.py -i $line -o $n+'.html'
n=$((n+1))
done < $filename
