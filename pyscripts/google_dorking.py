import urllib3
import sys



API_KEY = 'AIzaSyA99eo_AJYkHj_DGG9_zf1wqQQg1aUiZwI' 
query = sys.argv[1]


res = urllib3.urlopen("https://www.googleapis.com/customsearch/v1?"+"q="+urlencode(query)).read()
res = res.decode('utf-8')
res = json.loads(res)
print(res)

