package consts

// 用户状态枚举
const (
	UserStatusActive   = "active"   // 启用
	UserStatusDisabled = "disabled" // 禁用
)

// 用户状态切片，用于验证
var UserStatusList = []string{
	UserStatusActive,
	UserStatusDisabled,
}

// 用户默认状态
const UserDefaultStatus = UserStatusActive
