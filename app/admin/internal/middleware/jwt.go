package middleware

import (
	"gf-vue-admin/utility/jwt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// JWTAuth JWT认证中间件
func JWTAuth(r *ghttp.Request) {
	// 获取Authorization头
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"msg":  "缺少认证令牌",
		})
		return
	}

	// 检查Bearer前缀
	if !strings.HasPrefix(authHeader, "Bearer ") {
		r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"msg":  "认证令牌格式错误",
		})
		return
	}

	// 提取token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"msg":  "认证令牌不能为空",
		})
		return
	}

	// 验证token
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code": 401,
			"msg":  "认证令牌无效: " + err.Error(),
		})
		return
	}

	// 将用户信息存储到上下文中
	r.SetCtxVar("user_id", claims.UserID)
	r.SetCtxVar("username", claims.Username)
	r.SetCtxVar("email", claims.Email)

	// 继续处理请求
	r.Middleware.Next()
}

// GetUserFromContext 从上下文中获取用户信息
func GetUserFromContext(r *ghttp.Request) (userID uint64, username, email string, ok bool) {
	userIDVar := r.GetCtxVar("user_id")
	usernameVar := r.GetCtxVar("username")
	emailVar := r.GetCtxVar("email")

	if userIDVar == nil || usernameVar == nil || emailVar == nil {
		return 0, "", "", false
	}

	userID = userIDVar.Uint64()
	username = usernameVar.String()
	email = emailVar.String()

	return userID, username, email, true
}