#!/usr/bin/env bash

PROJPATH=$(pwd)
# folder where we save vet input
VFName=$PROJPATH/"vetComment.info.txt"
vendorFolder="vendor"
viewFolder="view"

# integration test
echo "==== START TESTING ===="
go test --tags=integration ./...

write () {
    echo $1 >> $VFName
}

recsearch(){
    for i in `find * -maxdepth 0 -type d`
        do
        if [ -d $i ] && [ $i != $vendorFolder ] && [ $i != $viewFolder ]
            then
                write "Going into directory $i"
                write "-----------------"
                write $i
                write "VET :"
                go tool vet $i/*.go >> $VFName
                write "LINT :"
                golint $i >> $VFName
                write "GOSIMPLE :"
                gosimple $i/... >> $VFName
                write "STATICCHECK :"
                staticcheck $i
                cd $i
                recsearch
                cd ..
        fi
    done
}



recsearch