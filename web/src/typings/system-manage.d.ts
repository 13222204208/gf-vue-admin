/**
 * 系统管理相关API类型定义
 */
declare namespace Api {
  /** 系统管理类型 */
  namespace SystemManage {
    /** 用户相关类型 */
    type UserList = Api.Common.PaginatedResponse<UserListItem>

    interface UserListItem {
      id: number
      avatar: string
      status: string
      username: string
      nickname: string
      phone: string
      email: string
      roleId: number
      roleName: string
      createdAt: string
    }

    type UserSearchParams = Partial<
      Pick<UserListItem, 'username' | 'phone' | 'status'> & Api.Common.CommonSearchParams
    >

    /** 用户创建参数 */
    interface UserCreateParams {
      username: string
      password: string
      nickname?: string
      phone?: string
      email?: string
      role_id: number
      status: Api.Common.EnableStatus
    }

    /** 用户更新参数 */
    interface UserUpdateParams {
      id: number
      username?: string
      nickname?: string
      phone?: string
      email?: string
      role_id?: number
      status?: Api.Common.EnableStatus
    }

    /** 用户创建响应 */
    interface UserCreateResponse {
      id: number
    }

    /** 用户删除参数 */
    interface UserDeleteParams {
      id: number
    }

    /** 角色相关类型 */
    type RoleList = Api.Common.PaginatedResponse<RoleListItem>

    interface RoleListItem {
      id: number
      name: string
      displayName: string
      description: string
      status: Api.Common.EnableStatus
      createdAt: string
    }

    interface ActiveRoleList {
      list: ActiveRoleListItem[]
    }

    interface ActiveRoleListItem {
      id: number
      displayName: string
    }

    type RoleSearchParams = Partial<
      Pick<RoleListItem, 'name' | 'displayName' | 'status'> & Api.Common.CommonSearchParams
    >

    interface RoleCreateParams {
      name: string
      displayName: string
      description?: string
      status: Api.Common.EnableStatus
    }

    interface RoleUpdateParams {
      id: number
      name?: string
      displayName?: string
      description?: string
      status?: Api.Common.EnableStatus
    }

    interface RoleCreateResponse {
      id: number
    }

    interface RoleDeleteParams {
      id: number
    }

    interface CommonOperationResponse {
      success: boolean
      message?: string
    }
  }
}
