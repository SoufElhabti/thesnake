import requests
import sys

url = sys.argv[1]
values = [] 

for i in range(100):
    r= requests.get(url)
    values.append(int(r.elapsed.total_seconds()))

average = sum(values) / float(len(values))
print('jitter is '+str(average))

