package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"server/server/auth"
	types "server/server/types/auth"
	"server/server/svc"
	"server/server/constant"
	"server/pkg/jwt"
)

type CodeLogin struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewCodeLogin(ctx context.Context, svcCtx *svc.ServiceContext) *CodeLogin {
	return &CodeLogin{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *CodeLogin) CodeLogin(req *types.CodeLoginRequest) (resp *types.LoginResponse, err error) {
	// check verificationUuid
	var verificationUuidVal string
	if err = l.svcCtx.Cache.Get(fmt.Sprintf("%s:%s", constant.CacheVerificationCodePrefix, req.VerificationUuid), &verificationUuidVal); err != nil {
		return nil, RegisterError
	}
	if verificationUuidVal != req.VerificationCode {
		return nil, errors.New("验证码错误")
	}

	user, err := l.svcCtx.Model.ManageUser.FindOneByCondition(l.ctx, nil, condition.Condition{
		Field:		"email",
		Operator:	condition.Equal,
		Value:		req.Email,
	})
	if err != nil {
		return nil, errors.New("用户名/密码错误")
	}

	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.NewChain().
		Equal("user_id", user.Id).
		Build()...)
	if err != nil {
		return nil, err
	}
	var roleIds []int64
	for _, userRole := range userRoles {
		roleIds = append(roleIds, userRole.RoleId)
	}

	j := jwt.NewJwt(l.svcCtx.Config.Jwt.AccessSecret)
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
	expirationTime := time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.AccessExpire) * time.Second).Unix()
	claims["exp"] = expirationTime

	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	claims["exp"] = time.Now().Add(time.Duration(l.svcCtx.Config.Jwt.RefreshExpire) * time.Second).Unix()
	refreshToken, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Token:		token,
		RefreshToken:	refreshToken,
	}, nil
}
