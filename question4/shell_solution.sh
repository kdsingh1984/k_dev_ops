#!/bin/bash

## variable assignment
SRC_FILE="./hostnames.txt"
DST_FILE="./exmaple_hosts.txt"

## check if source file exists
if [ ! -f $SRC_FILE ]; then
    echo "$SRC_FILE does not exists"
    exit 1
fi

## actual logic
cat $SRC_FILE | grep "xyz.com" | sed "s/xyz/exampledomain/g" > $DST_FILE

## check if above execution was successful
if [ $? -ne 0 ]; then
    echo "Not able to generate file"
fi
