#!/usr/bin/env bash

PROJPATH=$(pwd)
# folder where we save vet input
VFName=$PROJPATH/"vetComment.info.txt"
vendorFolder="vendor"
viewFolder="view"

# integration test
echo "==== START TESTING ===="
go test --tags=integration ./...

recsearch(){
    for i in `find * -maxdepth 0 -type d`
        do
        if [ -d $i ] && [ $i != $vendorFolder ] && [ $i != $viewFolder ]
            then
                echo "Going into directory $i"
                echo "-----------------" >> $VFName
                echo `pwd` >> $VFName
                echo "-----------------" >> $VFName
                go tool vet $i/*.go >> $VFName
                cd $i
                recsearch
                cd ..
        fi
    done
}

recsearch