// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	orgv1 "github.com/ikaiguang/go-srv-services/api/org/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/user-service/internal/domain/entity"
)

var _ = timeutil.YmdHms

// AssembleListOrg assemble listOrg
func AssembleListOrg(dataModels []*entities.Org) []*orgv1.Org {
	var newDataModels = make([]*orgv1.Org, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = AssembleOrg(dataModels[index])
	}
	return newDataModels
}

// AssembleOrg assemble Org
func AssembleOrg(dataModel *entities.Org) *orgv1.Org {
	newDataModel := &orgv1.Org{
		Id:          dataModel.Id,                                  // ID
		OrgUuid:     dataModel.OrgUuid,                             // uuid
		CreatedTime: dataModel.CreatedTime.Format(timeutil.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timeutil.YmdHms), // 最后修改时间
		IsDeleted:   dataModel.IsDeleted,                           // 是否已删除
		OrgEmail:    dataModel.OrgEmail,                            // 邮箱
		OrgName:     dataModel.OrgName,                             // 名称
		OrgLogo:     dataModel.OrgLogo,                             // logo
		OrgStatus:   dataModel.OrgStatus,                           // 状态
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