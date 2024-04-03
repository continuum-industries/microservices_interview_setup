#!/bin/bash

# Get current directory.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

echo $DIR

PROTO_DIR=${DIR}/../../shared/protobuffers
DEST_DIR=${DIR}/../src/protobuffers
mkdir $DEST_DIR
files=$(find "${PROTO_DIR}" -name '*.proto')
for file in $files; do
  echo $file

  FOLDERNAME=`(grep "^package " $file | awk '{print $2}' | sed 's/;*$//g')` ##Extracts the package name from the protobuffer file
  echo $FOLDERNAME

  PROTO_GEN_FOLDERNAME="$DEST_DIR/$FOLDERNAME"
  mkdir $PROTO_GEN_FOLDERNAME

  protoc -I ${PROTO_DIR} --go_out=paths=source_relative:${PROTO_GEN_FOLDERNAME} --go-grpc_out=paths=source_relative:${PROTO_GEN_FOLDERNAME} ${file}
done
