// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	userv1 "github.com/ikaiguang/go-srv-services/api/user-service/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

var _ = timeutil.YmdHms

// AssembleListUserRegMobile assemble listUserRegMobile
func AssembleListUserRegMobile(dataModels []*entities.UserRegMobile) []*userv1.UserRegMobile {
	var newDataModels = make([]*userv1.UserRegMobile, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = AssembleUserRegMobile(dataModels[index])
	}
	return newDataModels
}

// AssembleUserRegMobile assemble UserRegMobile
func AssembleUserRegMobile(dataModel *entities.UserRegMobile) *userv1.UserRegMobile {
	newDataModel := &userv1.UserRegMobile{
		Id:          dataModel.Id,                                  // ID
		UserMobile:  dataModel.UserMobile,                          // 用户手机号码
		CreatedTime: dataModel.CreatedTime.Format(timeutil.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timeutil.YmdHms), // 最后修改时间
		UserUuid:    dataModel.UserUuid,                            // uuid
	}

	return newDataModel
}
