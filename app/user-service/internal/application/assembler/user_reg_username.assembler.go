// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	userv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

var _ = timeutil.YmdHms

// UserRegUsernameList assemble listUserRegUsername
func (s *Assembler) UserRegUsernameList(dataModels []*entities.UserRegUsername) []*userv1.UserRegUsername {
	var newDataModels = make([]*userv1.UserRegUsername, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = s.UserRegUsername(dataModels[index])
	}
	return newDataModels
}

// UserRegUsername assemble UserRegUsername
func (s *Assembler) UserRegUsername(dataModel *entities.UserRegUsername) *userv1.UserRegUsername {
	newDataModel := &userv1.UserRegUsername{
		Id:          dataModel.Id,                                  // ID
		UserName:    dataModel.UserName,                            // 用户名
		CreatedTime: dataModel.CreatedTime.Format(timeutil.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timeutil.YmdHms), // 最后修改时间
		UserUuid:    dataModel.UserUuid,                            // uuid
	}

	return newDataModel
}
