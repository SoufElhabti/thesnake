#!/bin/python3

import mmh3
import requests
import codecs
 
response = requests.get('http://192.168.80.158:8983/solr/img/favicon.ico')
favicon = codecs.encode(response.content,"base64")
hash = mmh3.hash(favicon)
print(hash)
