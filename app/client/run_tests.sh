#!/bin/bash

for ((i=1; i<=30; i++))
do
    go run main.go

    mv logs.txt ./file_$i.txt
done