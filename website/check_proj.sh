#!/usr/bin/env bash

PROJPATH=$GOPATH/github.com/IhorBondartsov/OLX_Parser/website
# folder where we save vet input
VFName="vetComment.info.txt"

# integration test
go test --tags=integration ./...

# find all directs in parents fold
dircts=`find * -maxdepth 0 -type d`
for addr in $dircts
    do
        touch $VFName
        echo "-------------------------------"$addr"------------------------------------------" >> $VFName
        go tool vet $addr >> $VFName
    done
