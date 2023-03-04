#! /bin/bash

cd /home/taku/work/cpp/AtCoder || exit

if [ $# != 2 ]; then
    echo 引数は２つ必要です: 
    echo 第一引数:ABCの回
    echo 第二引数:ファイル名
    exit 0
fi
FOLDER="/home/taku/work/cpp/AtCoder/ABC/$1" || exit
FILE=${2%.*}.cpp
mkdir -p "$FOLDER"
cd "$FOLDER" || exit
cp -n ../../template.cpp ./"$FILE"
code "$FILE"