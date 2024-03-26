#!/usr/bin/env bash
# 参考:
# https://learnku.com/articles/36922
# https://blog.csdn.net/jarvan5/article/details/118026721
# https://cloud.tencent.com/developer/article/1817925
# usage:
# run this script at the directory where the .proto file locate
# codegen.sh protofilename.proto

if [ $# != 1 ] ; then
  echo "USAGE: $0 protofilename.proto"
  exit 1;
fi

proto=$1
protoc --go_out=plugins=grpc,paths=source_relative:. $proto
