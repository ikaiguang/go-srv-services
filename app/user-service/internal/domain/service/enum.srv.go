package srvs

import (
	userenumv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/enums"
)

// ToUserGenderEnum ...
func ToUserGenderEnum(gender string) userenumv1.UserGenderEnum_UserGender {
	return userenumv1.UserGenderEnum_UserGender(userenumv1.UserGenderEnum_UserGender_value[gender])
}

// ToUserStatusEnum ...
func ToUserStatusEnum(status string) userenumv1.UserStatusEnum_UserStatus {
	return userenumv1.UserStatusEnum_UserStatus(userenumv1.UserStatusEnum_UserStatus_value[status])
}
