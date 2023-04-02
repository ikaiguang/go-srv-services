#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/common/admin.common.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/enums/admin.enum.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/errors/admin.error.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/resources/admin.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/resources/admin_auth.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/resources/admin_reg_email.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/resources/admin_reg_mobile.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/resources/admin_reg_username.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/admin-service/v1/services/admin_auth.service.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/common/user.common.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/enums/user.enum.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/errors/user.error.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/org.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/user.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/user_auth.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/user_reg_email.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/user_reg_mobile.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/resources/user_reg_username.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/services/org.service.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/user-service/v1/services/user_auth.service.v1.proto
