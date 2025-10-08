package menus

import (
	"context"

	v1 "gf-vue-admin/app/admin/api/menus/v1"
	"gf-vue-admin/app/admin/internal/dao"
	"gf-vue-admin/app/admin/internal/model/entity"
)

func (c *ControllerV1) GetTree(ctx context.Context, req *v1.GetTreeReq) (res *v1.GetTreeRes, err error) {
	// 构建查询条件
	model := dao.SysMenu.Ctx(ctx)

	// 查询所有菜单数据，按排序字段和ID排序
	var menus []entity.SysMenu
	err = model.OrderAsc(dao.SysMenu.Columns().OrderNum).
		OrderAsc(dao.SysMenu.Columns().Id).
		Scan(&menus)
	if err != nil {
		return nil, err
	}

	// 构建菜单树
	menuTree := buildMenuTree(menus, 0)

	return &v1.GetTreeRes{
		List: menuTree,
	}, nil
}

// buildMenuTree 递归构建菜单树
func buildMenuTree(menus []entity.SysMenu, parentId uint64) []v1.MenuTreeItem {
	var tree []v1.MenuTreeItem

	for _, menu := range menus {
		if menu.ParentId == parentId {
			// 转换为API响应格式
			treeItem := v1.MenuTreeItem{
				GetByIdRes: v1.GetByIdRes{
					Id:            menu.Id,
					ParentId:      menu.ParentId,
					Title:         menu.Title,
					Icon:          menu.Icon,
					Path:          menu.Path,
					Component:     menu.Component,
					Redirect:      menu.Redirect,
					IsLink:        menu.IsLink,
					LinkUrl:       menu.LinkUrl,
					IsIframe:      menu.IsIframe,
					IsCache:       menu.IsCache,
					IsHide:        menu.IsHide,
					IsHideTab:     menu.IsHideTab,
					IsFullPage:    menu.IsFullPage,
					IsFixedTab:    menu.IsFixedTab,
					IsFirstLevel:  menu.IsFirstLevel,
					IsAuthButton:  menu.IsAuthButton,
					AuthMark:      menu.AuthMark,
					AuthList:      menu.AuthList,
					Roles:         menu.Roles,
					ShowBadge:     menu.ShowBadge,
					ShowTextBadge: menu.ShowTextBadge,
					ActivePath:    menu.ActivePath,
					OrderNum:      menu.OrderNum,
					Status:        menu.Status,
					Remark:        menu.Remark,
				},
			}

			// 递归查找子菜单
			treeItem.Children = buildMenuTree(menus, menu.Id)

			tree = append(tree, treeItem)
		}
	}

	return tree
}
