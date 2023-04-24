#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src app/snowflake-service/internal/conf/config.v1.proto
