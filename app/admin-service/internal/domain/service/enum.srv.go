package srvs

import (
	adminenumv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/enums"
)

// ToAdminGenderEnum ...
func ToAdminGenderEnum(gender string) adminenumv1.AdminGenderEnum_AdminGender {
	return adminenumv1.AdminGenderEnum_AdminGender(adminenumv1.AdminGenderEnum_AdminGender_value[gender])
}

// ToAdminStatusEnum ...
func ToAdminStatusEnum(status string) adminenumv1.AdminStatusEnum_AdminStatus {
	return adminenumv1.AdminStatusEnum_AdminStatus(adminenumv1.AdminStatusEnum_AdminStatus_value[status])
}
