// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menus

import (
	"context"

	v1 "gf-vue-admin/app/admin/api/menus/v1"
)

type IMenusV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetById(ctx context.Context, req *v1.GetByIdReq) (res *v1.GetByIdRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetTree(ctx context.Context, req *v1.GetTreeReq) (res *v1.GetTreeRes, err error)
	GetUserMenus(ctx context.Context, req *v1.GetUserMenusReq) (res *v1.GetUserMenusRes, err error)
}
