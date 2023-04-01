// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	adminv1 "github.com/ikaiguang/go-srv-services/api/admin/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/entity"
)

var _ = timeutil.YmdHms

// AssembleListAdmin assemble listAdmin
func AssembleListAdmin(dataModels []*entities.Admin) []*adminv1.Admin {
	var newDataModels = make([]*adminv1.Admin, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = AssembleAdmin(dataModels[index])
	}
	return newDataModels
}

// AssembleAdmin assemble Admin
func AssembleAdmin(dataModel *entities.Admin) *adminv1.Admin {
	newDataModel := &adminv1.Admin{
		Id:            dataModel.Id,                                  // ID
		AdminUuid:     dataModel.AdminUuid,                           // uuid
		CreatedTime:   dataModel.CreatedTime.Format(timeutil.YmdHms), // 创建时间
		UpdatedTime:   dataModel.UpdatedTime.Format(timeutil.YmdHms), // 最后修改时间
		IsDeleted:     dataModel.IsDeleted,                           // 是否已删除
		AdminEmail:    dataModel.AdminEmail,                          // 邮箱
		AdminNickname: dataModel.AdminNickname,                       // 昵称
		AdminAvatar:   dataModel.AdminAvatar,                         // 头像
		AdminGender:   dataModel.AdminGender,                         // 性别
		AdminStatus:   dataModel.AdminStatus,                         // 状态
		RegisterFrom:  dataModel.RegisterFrom,                        // 注册来源
	}
	// DeletedTime 删除时间
	if dataModel.DeletedTime != nil {
		newDataModel.DeletedTime = dataModel.DeletedTime.Format(timeutil.YmdHms)
	}
	// ActiveBeginTime 激活开始时间
	if dataModel.ActiveBeginTime != nil {
		newDataModel.ActiveBeginTime = dataModel.ActiveBeginTime.Format(timeutil.YmdHms)
	}
	// ActiveEndTime 激活结束时间
	if dataModel.ActiveEndTime != nil {
		newDataModel.ActiveEndTime = dataModel.ActiveEndTime.Format(timeutil.YmdHms)
	}
	// DisableTime 禁用时间
	if dataModel.DisableTime != nil {
		newDataModel.DisableTime = dataModel.DisableTime.Format(timeutil.YmdHms)
	}
	// BlacklistTime 黑名单时间
	if dataModel.BlacklistTime != nil {
		newDataModel.BlacklistTime = dataModel.BlacklistTime.Format(timeutil.YmdHms)
	}

	return newDataModel
}