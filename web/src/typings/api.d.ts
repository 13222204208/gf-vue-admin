/**
 * namespace: Api
 *
 * 所有接口相关类型定义
 * 在.vue文件使用会报错，需要在 eslint.config.mjs 中配置 globals: { Api: 'readonly' }
 */

declare namespace Api {
  /** 通用类型 */
  namespace Common {
    /** 分页参数 */
    interface PaginationParams {
      /** 当前页码 */
      current: number
      /** 每页条数 */
      size: number
      /** 总条数 */
      total: number
    }

    /** 通用搜索参数 */
    type CommonSearchParams = Pick<PaginationParams, 'current' | 'size'>

    /** 分页响应基础结构 */
    interface PaginatedResponse<T = any> {
      records: T[]
      current: number
      size: number
      total: number
    }

    /** 启用状态 */
    type EnableStatus = 'active' | 'disabled'
  }

  /** 认证类型 */
  namespace Auth {
    /** 登录参数 */
    interface LoginParams {
      username: string
      password: string
    }

    /** 登录响应 */
    interface LoginResponse {
      token: string
      refreshToken: string
    }

    /** 用户信息 */
    interface UserInfo {
      buttons: string[]
      roles: string[]
      userId: number
      userName: string
      email: string
      avatar?: string
    }
  }

  /** 系统管理类型 */
  namespace SystemManage {
    /** 用户列表 */
    type UserList = Api.Common.PaginatedResponse<UserListItem>

    /** 用户列表项 */
    interface UserListItem {
      id: number
      avatar: string
      status: string
      username: string
      nickname: string
      phone: string
      email: string
      userRoles: string[]
      createdAt: string
      updatedAt: string
    }

    /** 用户搜索参数 */
    type UserSearchParams = Partial<
      Pick<UserListItem, 'id' | 'userName' | 'userGender' | 'userPhone' | 'userEmail' | 'status'> &
        Api.Common.CommonSearchParams
    >

    /** 角色列表 */
    type RoleList = Api.Common.PaginatedResponse<RoleListItem>

    /** 角色列表项 */
    interface RoleListItem {
      id: number
      name: string
      displayName: string
      description: string
      status: Api.Common.EnableStatus
      createdAt: string
      updatedAt: string
    }

    /** 角色搜索参数 */
    type RoleSearchParams = Partial<
      Pick<RoleListItem, 'name' | 'displayName' | 'status'> &
        Api.Common.CommonSearchParams
    >

    /** 角色创建参数 */
    interface RoleCreateParams {
      name: string
      displayName: string
      description?: string
      status: Api.Common.EnableStatus
    }

    /** 角色更新参数 */
    interface RoleUpdateParams {
      id: number
      name?: string
      displayName?: string
      description?: string
      status?: Api.Common.EnableStatus
    }

    /** 角色创建响应 */
    interface RoleCreateResponse {
      id: number
    }

    /** 角色详情响应 */
    interface RoleDetailResponse extends RoleListItem {}

    /** 角色删除参数 */
    interface RoleDeleteParams {
      id: number
    }

    /** 角色批量删除参数 */
    interface RoleBatchDeleteParams {
      ids: number[]
    }

    /** 通用操作响应 */
    interface CommonOperationResponse {
      success: boolean
      message?: string
    }
  }
}
