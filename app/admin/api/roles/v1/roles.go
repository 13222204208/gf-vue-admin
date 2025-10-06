package v1

import (
	"gf-vue-admin/app/admin/internal/model/common"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateReq 创建角色请求
type CreateReq struct {
	g.Meta      `path:"/roles" tags:"Roles" method:"post" summary:"创建角色"`
	Name        string `json:"name" v:"required|length:2,50#角色名称不能为空|角色名称长度为2-50位" dc:"角色名称（英文唯一标识）"`
	DisplayName string `json:"displayName" v:"required|length:2,50#显示名称不能为空|显示名称长度为2-50位" dc:"角色显示名称"`
	Description string `json:"description" v:"max-length:200#描述长度不能超过200字符" dc:"角色描述"`
	Status      string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态: active=启用, disabled=禁用"`
}

type CreateRes struct {
	g.Meta `mime:"application/json"`
	Id     uint64 `json:"id" dc:"角色ID"`
}

// UpdateReq 更新角色请求
type UpdateReq struct {
	g.Meta      `path:"/roles/{id}" tags:"Roles" method:"put" summary:"更新角色"`
	Id          uint64 `json:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
	Name        string `json:"name" v:"length:2,50#角色名称长度为2-50位" dc:"角色名称（英文唯一标识）"`
	DisplayName string `json:"displayName" v:"length:2,50#显示名称长度为2-50位" dc:"角色显示名称"`
	Description string `json:"description" v:"max-length:200#描述长度不能超过200字符" dc:"角色描述"`
	Status      string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态: active=启用, disabled=禁用"`
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
}

// DeleteReq 删除角色请求
type DeleteReq struct {
	g.Meta `path:"/roles/{id}" tags:"Roles" method:"delete" summary:"删除角色"`
	Id     uint64 `json:"id" v:"required|min:1#角色ID不能为空|角色ID必须大于0" dc:"角色ID"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}

type GetByIdRes struct {
	g.Meta      `mime:"application/json"`
	Id          uint64      `json:"id" dc:"主键ID"`
	Name        string      `json:"name" dc:"角色名称（英文唯一标识）"`
	DisplayName string      `json:"displayName" dc:"角色显示名称"`
	Description string      `json:"description" dc:"角色描述"`
	Status      string      `json:"status" dc:"状态: active=启用, disabled=禁用"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// GetListReq 获取角色列表请求
type GetListReq struct {
	g.Meta `path:"/roles" tags:"Roles" method:"get" summary:"获取角色列表"`
	common.CurrentReq
	Name        string `json:"name" dc:"角色名称（模糊搜索）"`
	DisplayName string `json:"displayName" dc:"显示名称（模糊搜索）"`
	Status      string `json:"status" v:"in:active,disabled#状态值不正确" dc:"状态筛选"`
}

type GetListRes struct {
	g.Meta `mime:"application/json"`
	common.CurrentRes
	List []GetByIdRes `json:"list" dc:"角色列表"`
}

// GetActiveListReq 获取所有启用角色，只返回角色ID和角色名称
type GetActiveListReq struct {
	g.Meta `path:"/roles/active" tags:"Roles" method:"get" summary:"获取所有启用角色"`
}

// GetActiveListRes 获取所有启用角色，只返回角色ID和角色名称
type GetActiveListRes struct {
	g.Meta `mime:"application/json"`
	List   []struct {
		Id          uint64 `json:"id" dc:"主键ID"`
		DisplayName string `json:"displayName" dc:"角色显示名称"`
	} `json:"list" dc:"角色列表"`
}
