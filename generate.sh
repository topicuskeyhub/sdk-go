#!/bin/bash

mkdir -p tmp
cd tmp
wget https://github.com/microsoft/kiota/releases/download/v1.28.0/linux-x64.zip
unzip linux-x64.zip
cd ..
find . -type f -name "*.go" -print0 | xargs -0 grep -Z -L "not-generated" | xargs -0 rm -f
find . -type d -empty -delete
env KIOTA_Generation:MaxDegreeOfParallelism=1 tmp/kiota update -o .
rm -rf tmp
