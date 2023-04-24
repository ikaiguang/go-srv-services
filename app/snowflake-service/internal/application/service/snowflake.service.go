package services

import (
	"context"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	snowflakeerrorv1 "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/errors"
	snowflakev1 "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/resources"
	snowflakeservicev1 "github.com/ikaiguang/go-srv-services/api/snowflake-service/v1/services"
	assemblers "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/application/assembler"
	srvs "github.com/ikaiguang/go-srv-services/app/snowflake-service/internal/domain/service"
	"strings"
)

// snowflakeNodeID ...
type snowflakeNodeID struct {
	snowflakeservicev1.UnimplementedSrvSnowflakeNodeIDServer

	assembler    *assemblers.Assembler
	snowflakeSrv *srvs.SnowflakeSrv
}

// NewSnowflakeNodeID ...
func NewSnowflakeNodeID(
	assembler *assemblers.Assembler,
	snowflakeSrv *srvs.SnowflakeSrv,
) snowflakeservicev1.SrvSnowflakeNodeIDServer {
	return &snowflakeNodeID{
		assembler:    assembler,
		snowflakeSrv: snowflakeSrv,
	}
}

// GetNodeId 获取节点ID
func (s *snowflakeNodeID) GetNodeId(ctx context.Context, in *snowflakev1.GetNodeIdReq) (*snowflakev1.SnowflakeNodeID, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := snowflakeerrorv1.ERROR_INSTANCE_ID_IS_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}

	dataModel, err := s.snowflakeSrv.GetNodeId(ctx, in)
	if err != nil {
		return nil, err
	}
	return s.assembler.SnowflakeNodeID(dataModel), nil
}

// ExtendNodeId 续期
func (s *snowflakeNodeID) ExtendNodeId(ctx context.Context, in *snowflakev1.ExtendNodeIdReq) (*snowflakev1.ExtendNodeIdResp, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := snowflakeerrorv1.ERROR_INSTANCE_ID_IS_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}
	success, err := s.snowflakeSrv.ExtendNodeId(ctx, in)
	if err != nil {
		return nil, err
	}
	return &snowflakev1.ExtendNodeIdResp{Success: success}, nil
}
