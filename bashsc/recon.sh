#!/bin/bash


amass enum -d $1 -config ~/amassconfig/config.ini -o domains 

httpx -l domains  -title -web-server -status-code --follow-redirects -o infos

cat domains | waybackurls > wayback

# this one liner is damn good 
#shodan domain att.com| awk '{print $3}'|  httpx -silent | ~/go/bin/anew | xargs -I@ ~/go/bin/jaeles scan -c 100 -s /jaeles-signatures/ -u @
cat domains| awk '{print $3}'|  httpx -silent | ~/go/bin/anew | xargs -I@ ~/go/bin/jaeles scan -c 100 -s /jaeles-signatures/ -u @

