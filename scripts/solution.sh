#!/bin/bash

d=$1
p=$2

mkdir "cmd/d${d}${p}"

cp resources/solution-template.go "cmd/d${d}${p}/main.go"

