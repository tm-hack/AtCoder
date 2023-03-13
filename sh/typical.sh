#! /bin/bash

cd /home/taku/work/cpp/AtCoder || exit

if [ $# != 1 ]; then
    echo 引数は１つ必要です: 
    echo 第一引数:ファイル名
    exit 0
fi
FOLDER="./Typical90/" || exit
FILE=${1%.*}.cpp
mkdir -p "$FOLDER"
cd "$FOLDER" || exit
cp -n ../template.cpp ./"$FILE"
code "$FILE"