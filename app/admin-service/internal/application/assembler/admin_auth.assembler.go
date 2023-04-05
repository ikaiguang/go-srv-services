package assemblers

import (
	adminv1 "github.com/ikaiguang/go-srv-services/api/admin-service/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/entity"
	srvs "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/service"
)

// AuthInfo ...
func (s *Assembler) AuthInfo(adminModel *entities.Admin) *adminv1.Info {
	return &adminv1.Info{
		Id:            adminModel.Id,
		AdminUuid:     adminModel.AdminUuid,
		AdminNickname: adminModel.AdminNickname,
		AdminAvatar:   adminModel.AdminAvatar,
		AdminGender:   srvs.ToAdminGenderEnum(adminModel.AdminGender),
		AdminStatus:   srvs.ToAdminStatusEnum(adminModel.AdminStatus),
	}
}

// LoginResp ...
func (s *Assembler) LoginResp(authInfo *adminv1.Info, signedString string) *adminv1.LoginResp {
	return &adminv1.LoginResp{
		AdminInfo: authInfo,
		Token:     signedString,
	}
}
