import request from '@/utils/http'

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

// 菜单相关接口

// 获取菜单数据（用于路由）
export function fetchGetMenuList() {
  return request.get<Api.SystemManage.MenuTreeResponse>({
    url: '/menus/user'
  })
}

// 菜单管理API接口

// 获取菜单列表
export function fetchGetMenuManageList(params: Api.SystemManage.MenuSearchParams) {
  return request.get<Api.SystemManage.MenuList>({
    url: '/menus',
    params
  })
}

// 获取菜单树
export function fetchGetMenuTree() {
  return request.get<Api.SystemManage.MenuTreeResponse>({
    url: '/menus/tree'
  })
}

// 根据ID获取菜单详情
export function fetchGetMenuById(id: number) {
  return request.get<Api.SystemManage.MenuListItem>({
    url: `/menus/${id}`
  })
}

// 创建菜单
export function fetchCreateMenu(params: Api.SystemManage.MenuCreateParams) {
  return request.post<Api.SystemManage.MenuCreateResponse>({
    url: '/menus',
    params
  })
}

// 更新菜单
export function fetchUpdateMenu(params: Api.SystemManage.MenuUpdateParams) {
  const { id, ...updateData } = params
  return request.put<Api.SystemManage.CommonOperationResponse>({
    url: `/menus/${id}`,
    params: updateData
  })
}

// 删除菜单
export function fetchDeleteMenu(id: number) {
  return request.del<Api.SystemManage.CommonOperationResponse>({
    url: `/menus/${id}`
  })
}
