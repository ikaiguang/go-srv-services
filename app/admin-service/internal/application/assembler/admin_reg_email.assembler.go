// Package assemblers
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package assemblers

import (
	timeutil "github.com/ikaiguang/go-srv-kit/kit/time"

	adminv1 "github.com/ikaiguang/go-srv-services/api/admin/v1/resources"
	entities "github.com/ikaiguang/go-srv-services/app/admin-service/internal/domain/entity"
)

var _ = timeutil.YmdHms

// AssembleListAdminRegEmail assemble listAdminRegEmail
func AssembleListAdminRegEmail(dataModels []*entities.AdminRegEmail) []*adminv1.AdminRegEmail {
	var newDataModels = make([]*adminv1.AdminRegEmail, len(dataModels))
	for index := range dataModels {
		newDataModels[index] = AssembleAdminRegEmail(dataModels[index])
	}
	return newDataModels
}

// AssembleAdminRegEmail assemble AdminRegEmail
func AssembleAdminRegEmail(dataModel *entities.AdminRegEmail) *adminv1.AdminRegEmail {
	newDataModel := &adminv1.AdminRegEmail{
		Id:          dataModel.Id,                                  // ID
		AdminEmail:  dataModel.AdminEmail,                          // 管理员手机号码
		CreatedTime: dataModel.CreatedTime.Format(timeutil.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timeutil.YmdHms), // 最后修改时间
		AdminUuid:   dataModel.AdminUuid,                           // uuid
	}

	return newDataModel
}
