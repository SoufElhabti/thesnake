#!/bin/bash


amass enum -d $1 -config ~/amassconfig/config.ini -o domains 

httpx -l domains  -title -web-server -status-code --follow-redirects -o infos

cat domains | waybackurls > wayback


