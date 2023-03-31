#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/config/v1/config.v1.proto
