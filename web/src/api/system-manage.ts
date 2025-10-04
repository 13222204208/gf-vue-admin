import request from '@/utils/http'
import { AppRouteRecord } from '@/types/router'
import { asyncRoutes } from '@/router/routes/asyncRoutes'
import { menuDataToRouter } from '@/router/utils/menuToRouter'

// 获取用户列表
export function fetchGetUserList(params: Api.SystemManage.UserSearchParams) {
  return request.get<Api.SystemManage.UserList>({
    url: '/users',
    params
  })
}

// 获取角色列表
export function fetchGetRoleList(params: Api.SystemManage.RoleSearchParams) {
  return request.get<Api.SystemManage.RoleList>({
    url: '/roles',
    params
  })
}

// 创建角色
export function fetchCreateRole(params: Api.SystemManage.RoleCreateParams) {
  return request.post<Api.SystemManage.RoleCreateResponse>({
    url: '/roles',
    params
  })
}

// 根据ID获取角色详情
export function fetchGetRoleById(id: number) {
  return request.get<Api.SystemManage.RoleDetailResponse>({
    url: `/roles/${id}`
  })
}

// 更新角色
export function fetchUpdateRole(params: Api.SystemManage.RoleUpdateParams) {
  const { id, ...updateData } = params
  return request.put<Api.SystemManage.CommonOperationResponse>({
    url: `/roles/${id}`,
    params: updateData
  })
}

// 删除角色
export function fetchDeleteRole(id: number) {
  return request.del<Api.SystemManage.CommonOperationResponse>({
    url: `/roles/${id}`
  })
}

// 批量删除角色
export function fetchDeleteRolesBatch(ids: number[]) {
  return request.del<Api.SystemManage.CommonOperationResponse>({
    url: '/roles/batch',
    params: { ids }
  })
}

interface MenuResponse {
  menuList: AppRouteRecord[]
}

// 获取菜单数据（模拟）
// 当前使用本地模拟路由数据，实际项目中请求接口返回 asyncRoutes.ts 文件的数据
export async function fetchGetMenuList(delay = 300): Promise<MenuResponse> {
  try {
    // 模拟接口返回的菜单数据
    const menuData = asyncRoutes
    // 处理菜单数据
    const menuList = menuData.map((route) => menuDataToRouter(route))
    // 模拟接口延迟
    await new Promise((resolve) => setTimeout(resolve, delay))

    return { menuList }
  } catch (error) {
    throw error instanceof Error ? error : new Error('获取菜单失败')
  }
}
