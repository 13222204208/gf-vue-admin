package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "gf-vue-admin/app/admin/api/auth/v1"
	"gf-vue-admin/utility/jwt"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	// 使用刷新令牌生成新的访问令牌
	newAccessToken, err := jwt.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		return nil, gerror.New("刷新令牌失败: " + err.Error())
	}

	// 解析刷新令牌获取用户信息
	claims, err := jwt.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, gerror.New("解析刷新令牌失败: " + err.Error())
	}

	// 生成新的刷新令牌
	newRefreshToken, err := jwt.GenerateRefreshToken(claims.UserID, claims.Username, claims.Email)
	if err != nil {
		return nil, gerror.New("生成新刷新令牌失败: " + err.Error())
	}

	return &v1.RefreshTokenRes{
		Token:        newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}