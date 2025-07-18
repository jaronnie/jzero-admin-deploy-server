package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/server/auth"
	types "github.com/jzero-io/jzero-admin/server/server/types/auth"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/server/constant"
)

type CodeLogin struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewCodeLogin(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *CodeLogin {
	return &CodeLogin{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *CodeLogin) CodeLogin(req *types.CodeLoginRequest) (resp *types.LoginResponse, err error) {
	config, err := l.svcCtx.Config.GetConfig()
	if err != nil {
		return nil, err
	}

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

	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_user_role.UserId, user.Id).
		Build()...)
	if err != nil {
		return nil, err
	}
	var roleIds []int64
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	marshal, err := json.Marshal(auth.Auth{
		Id:		int(user.Id),
		Username:	user.Username,
		RoleIds:	roleIds,
	})
	if err != nil {
		return nil, err
	}

	var claims map[string]any
	err = json.Unmarshal(marshal, &claims)
	if err != nil {
		return nil, err
	}

	// token 过期时间
	expirationTime := time.Now().Add(time.Duration(config.Jwt.AccessExpire) * time.Second).Unix()
	claims["exp"] = expirationTime

	token, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(config.Jwt.RefreshExpire) * time.Second).Unix()
	refreshToken, err := CreateToken(l.svcCtx.MustGetConfig().Jwt.AccessSecret, claims)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:		token,
		RefreshToken:	refreshToken,
	}, nil
}
