package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/server/constant"
	types "github.com/jzero-io/jzero-admin/server/server/types/auth"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user"
)

type ResetPassword struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewResetPassword(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *ResetPassword {
	return &ResetPassword{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *ResetPassword) ResetPassword(req *types.ResetPasswordRequest) (resp *types.ResetPasswordResponse, err error) {
	// check verificationUuid
	var verificationUuidVal string
	if err = l.svcCtx.Cache.Get(fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, req.VerificationUuid), &verificationUuidVal); err != nil {
		return nil, RegisterError
	}
	if verificationUuidVal != req.VerificationCode {
		return nil, errors.New("验证码错误")
	}

	user, err := l.svcCtx.Model.ManageUser.FindOneByCondition(l.ctx, nil, condition.Condition{
		Field:		manage_user.Email,
		Operator:	condition.Equal,
		Value:		req.Email,
	})
	if err != nil {
		return nil, errors.New("用户名/密码错误")
	}

	user.Password = req.Password
	err = l.svcCtx.Model.ManageUser.Update(l.ctx, nil, user)

	return
}
