package auth

import (
	"context"
	"fmt"
	"server/server/constant"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPassword struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPassword(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPassword {
	return &ResetPassword{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPassword) ResetPassword(req *types.ResetPasswordRequest) (resp *types.ResetPasswordResponse, err error) {

	var verificationUuidVal string
	if err = l.svcCtx.Cache.Get(fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, req.VerificationUuid), &verificationUuidVal); err != nil {
		return nil, RegisterError
	}
	if verificationUuidVal != req.VerificationCode {
		return nil, errors.New("验证码错误")
	}

	user, err := l.svcCtx.Model.SystemUser.FindOneByCondition(l.ctx, nil, condition.Condition{
		Field:    "email",
		Operator: condition.Equal,
		Value:    req.Email,
	})
	if err != nil {
		return nil, errors.New("用户名/密码错误")
	}

	user.Password = req.Password
	err = l.svcCtx.Model.SystemUser.Update(l.ctx, nil, user)

	return
}
