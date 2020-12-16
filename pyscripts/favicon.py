#!/bin/python3

import sys
import mmh3
import requests
import codecs


response = requests.get(sys.argv[1])
favicon = codecs.encode(response.content,"base64")
hash = mmh3.hash(favicon)
print(hash)
