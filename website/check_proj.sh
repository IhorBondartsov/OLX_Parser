#!/usr/bin/env bash

PROJPATH=$(pwd)
# folder where we save vet input
VFName=$PROJPATH/"vetComment.info.txt"
vendorFolder="vendor"
viewFolder="view"

# integration test
echo "==== START TESTING ===="
go test --tags=integration ./...

gometalinter  --vendor -e=view ./... >> $VFName

#write () {
#    echo $1 >> $VFName
#}
#
#write "GOSEC ..."
#gosec  -exclude=G303 ./... >> $VFName
#write "MALIGNED ..."
#maligned cmd/compile/internal/gc cmd/link/internal/ld
#write "GOCONST ..."
#goconst -ignore="view|vendor" ./...
#
#recsearch(){
#    for i in `find * -maxdepth 0 -type d`
#        do
#        if [ -d $i ] && [ $i != $vendorFolder ] && [ $i != $viewFolder ]
#            then
#                write "Going into directory $i"
#                write "-----------------"
#                write $i
#                write "VET :"
#                go tool vet $i/*.go >> $VFName
#                write "LINT :"
#                golint $i >> $VFName
#                write "GOSIMPLE :"
#                gosimple $i/*.go >> $VFName
#                write "STATICCHECK :"
#                staticcheck $i/*.go
#                write "ERRCHECK :"
#                errcheck -blank ./*.go
#                write "GOCYCLO :"
#                gocyclo -over 7 .
#                cd $i
#                recsearch
#                cd ..
#        fi
#    done
#}
#
#recsearch
