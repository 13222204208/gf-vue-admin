package common

import "github.com/gogf/gf/v2/frame/g"

// CurrentReq 分页请求结构
type CurrentReq struct {
	Current int `json:"current" v:"min:1#页码必须大于0" dc:"页码" d:"1"`
	Size    int `json:"size" v:"between:1,100#每页数量必须在1-100之间" dc:"每页数量" d:"10"`
}

// CurrentRes 分页响应结构
type CurrentRes struct {
	Total   int `json:"total" dc:"总数"`
	Current int `json:"current" dc:"当前页码"`
	Size    int `json:"size" dc:"每页数量"`
}

// ListReq 通用列表请求结构（包含分页）
type ListReq struct {
	g.Meta `method:"get"`
	CurrentReq
}

// ListRes 通用列表响应结构（包含分页）
type ListRes struct {
	g.Meta `mime:"application/json"`
	CurrentRes
}

// GetCurrentInfo 从分页请求中获取分页信息
func (p *CurrentReq) GetCurrentInfo() (current, size int) {
	current = p.Current
	if current <= 0 {
		current = 1
	}

	size = p.Size
	if size <= 0 {
		size = 10
	} else if size > 100 {
		size = 100
	}

	return current, size
}

// BuildCurrentRes 构建分页响应
func BuildCurrentRes(total, current, size int) CurrentRes {
	return CurrentRes{
		Total:   total,
		Current: current,
		Size:    size,
	}
}
