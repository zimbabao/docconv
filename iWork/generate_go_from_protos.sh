#!/bin/bash

cd `dirname $0`

proto_prefixes=(KN TN TP TSA TSCE TSD TSK TSP TSS TST TSWP)
for prefix in "${proto_prefixes[@]}"; do 
  echo $prefix;
  protoc --proto_path=pb-schema --go_out=paths=source_relative:${prefix} pb-schema/${prefix}*
done

echo "TSCH"
protoc --proto_path=pb-schema/ --go_out=paths=source_relative:TSCH  pb-schema/TSCHArchives.proto pb-schema/TSCH3DArchives.proto pb-schema/TSCHArchives.Common.proto pb-schema/TSCHCommandArchives.proto

echo "TSCH/PreUFF"
protoc --proto_path=pb-schema/ --go_out=paths=source_relative:TSCH/PreUFF pb-schema/TSCHPreUFFArchives.proto

echo "TSCH/Generated"
protoc --proto_path=pb-schema/ --go_out=paths=source_relative:TSCH/Generated pb-schema/TSCHArchives.GEN.proto
