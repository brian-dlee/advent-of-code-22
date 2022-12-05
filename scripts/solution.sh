#!/bin/bash

d=$1
p=$2

if [[ ! -e "cmd/d${d}${p}" ]]; then
  mkdir "cmd/d${d}${p}"
fi

cp resources/solution-template.go "cmd/d${d}${p}/main.go"
