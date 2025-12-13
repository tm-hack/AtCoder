#!/bin/bash

sendfile="main.go"

dir=$(cat ./contest)
echo -e CONTEST: $dir
echo -n QUESTION?
read q
mkdir -p _result/_$dir/$q
mv -i $sendfile _result/_$dir/$q
git add _result/_$dir/$q/$sendfile
git commit -m "$dir $q"
cp -i _template/$sendfile ./$sendfile
