#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src app/admin-service/internal/conf/config.v1.proto
