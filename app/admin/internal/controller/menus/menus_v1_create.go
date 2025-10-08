package menus

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "gf-vue-admin/app/admin/api/menus/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 检查菜单路径是否已存在
	count, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Path, req.Path).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("菜单路径已存在")
	}

	// 如果指定了父级菜单，检查父级菜单是否存在
	if req.ParentId > 0 {
		parentCount, err := dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, req.ParentId).Count()
		if err != nil {
			return nil, err
		}
		if parentCount == 0 {
			return nil, gerror.New("父级菜单不存在")
		}
	}

	// 设置默认状态
	status := req.Status
	if status == "" {
		status = "active"
	}

	// 构建插入数据，只插入非空字段
	insertData := do.SysMenu{
		Title:  req.Title, // 必填字段
		Path:   req.Path,  // 必填字段
		Status: status,    // 默认或指定状态
	}

	// 可选字段，只有不为空时才设置
	if req.ParentId > 0 {
		insertData.ParentId = req.ParentId
	}
	if req.Icon != "" {
		insertData.Icon = req.Icon
	}
	if req.Component != "" {
		insertData.Component = req.Component
	}
	if req.Redirect != "" {
		insertData.Redirect = req.Redirect
	}
	if req.LinkUrl != "" {
		insertData.LinkUrl = req.LinkUrl
	}
	if req.AuthMark != "" {
		insertData.AuthMark = req.AuthMark
	}
	if req.AuthList != "" {
		insertData.AuthList = req.AuthList
	}
	if req.Roles != "" {
		insertData.Roles = req.Roles
	}
	if req.ShowTextBadge != "" {
		insertData.ShowTextBadge = req.ShowTextBadge
	}
	if req.ActivePath != "" {
		insertData.ActivePath = req.ActivePath
	}
	if req.Remark != "" {
		insertData.Remark = req.Remark
	}

	// 布尔类型字段（0/1），设置实际值
	insertData.IsLink = req.IsLink
	insertData.IsIframe = req.IsIframe
	insertData.IsCache = req.IsCache
	insertData.IsHide = req.IsHide
	insertData.IsHideTab = req.IsHideTab
	insertData.IsFullPage = req.IsFullPage
	insertData.IsFixedTab = req.IsFixedTab
	insertData.IsFirstLevel = req.IsFirstLevel
	insertData.IsAuthButton = req.IsAuthButton
	insertData.ShowBadge = req.ShowBadge

	// 数字类型字段，设置实际值（包括0）
	insertData.OrderNum = req.OrderNum

	// 插入数据
	insertId, err := dao.SysMenu.Ctx(ctx).Data(insertData).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{
		Id: gconv.Uint64(insertId),
	}, nil
}
