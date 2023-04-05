package assemblers

import (
	userv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
	srvs "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/service"
)

// AuthInfo ...
func (s *Assembler) AuthInfo(userModel *entities.User) *userv1.Info {
	return &userv1.Info{
		Id:           userModel.Id,
		UserUuid:     userModel.UserUuid,
		UserNickname: userModel.UserNickname,
		UserAvatar:   userModel.UserAvatar,
		UserGender:   srvs.ToUserGenderEnum(userModel.UserGender),
		UserStatus:   srvs.ToUserStatusEnum(userModel.UserStatus),
	}
}

// LoginResp ...
func (s *Assembler) LoginResp(authInfo *userv1.Info, signedString string) *userv1.LoginResp {
	return &userv1.LoginResp{
		UserInfo: authInfo,
		Token:    signedString,
	}
}
