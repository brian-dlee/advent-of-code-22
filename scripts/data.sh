#!/bin/bash

d=$1
p=$2
t=$3

case "$t" in
s1|s2|s3|input) ;;
*) echo "invalid input type: $t. provide s1, s2, s3, or input." >&2; exit 1;;
esac

mkdir "data/d${d}"

pbpaste >"data/d${d}/${p}.${t}.txt"

