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

// 创建用户
export function fetchCreateUser(params: Api.SystemManage.UserCreateParams) {
  return request.post<Api.SystemManage.UserCreateResponse>({
    url: '/users',
    params
  })
}

// 更新用户
export function fetchUpdateUser(params: Api.SystemManage.UserUpdateParams) {
  const { id, ...updateData } = params
  return request.put<Api.SystemManage.CommonOperationResponse>({
    url: `/users/${id}`,
    params: updateData
  })
}

// 删除用户
export function fetchDeleteUser(id: number) {
  return request.del<Api.SystemManage.CommonOperationResponse>({
    url: `/users/${id}`
  })
}

// 获取角色列表
export function fetchGetRoleList(params: Api.SystemManage.RoleSearchParams) {
  return request.get<Api.SystemManage.RoleList>({
    url: '/roles',
    params
  })
}
//获取已启用所有角色
export function fetchGetEnabledRoleList() {
  return request.get<Api.SystemManage.ActiveRoleList>({
    url: '/roles/active'
  })
}

// 创建角色
export function fetchCreateRole(params: Api.SystemManage.RoleCreateParams) {
  return request.post<Api.SystemManage.RoleCreateResponse>({
    url: '/roles',
    params
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
