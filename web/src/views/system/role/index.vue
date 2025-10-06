<template>
  <div class="role-page">
    <RoleSearch
      v-show="showSearchBar"
      v-model="searchForm"
      @search="handleSearch"
      @reset="resetSearchParams"
    />
    <ArtTableHeader
      v-model:show-search-bar="showSearchBar"
      v-model:column-checks="columnChecks"
      title="角色管理"
      :loading="loading"
      @refresh="refreshData"
    >
      <template #left>
        <ElSpace wrap>
          <ElButton @click="showDialog('add')" v-ripple>新增角色</ElButton>
        </ElSpace>
      </template>
    </ArtTableHeader>
    <ArtTable
      :columns="columns"
      :data="data"
      :loading="loading"
      :pagination="pagination"
      @pagination:size-change="handleSizeChange"
      @pagination:current-change="handleCurrentChange"
    />
    <RoleEditDialog
      v-model="dialogVisible"
      :dialog-type="dialogType"
      :role-data="currentRoleData"
      @success="refreshData"
    />
    <RolePermissionDialog
      v-model="permissionDialog"
      :role-data="currentRoleData"
      @success="refreshData"
    />
  </div>
</template>

<script setup lang="ts">
  import { ButtonMoreItem } from '@/components/core/forms/art-button-more/index.vue'
  import { Setting, Edit, Delete } from '@element-plus/icons-vue'
  import { useTable } from '@/composables/useTable'
  import { fetchGetRoleList, fetchDeleteRole } from '@/api/system-manage'
  import ArtButtonMore from '@/components/core/forms/art-button-more/index.vue'
  import RoleSearch from './modules/role-search.vue'
  import RoleEditDialog from './modules/role-edit-dialog.vue'
  import RolePermissionDialog from './modules/role-permission-dialog.vue'

  defineOptions({ name: 'Role' })

  type RoleListItem = Api.SystemManage.RoleListItem

  // 搜索表单
  const searchForm = ref({
    displayName: undefined,
    name: undefined,
    status: undefined
  })

  const showSearchBar = ref(false)

  const dialogVisible = ref(false)
  const permissionDialog = ref(false)
  const currentRoleData = ref<RoleListItem | undefined>(undefined)

  const {
    columns,
    columnChecks,
    data,
    loading,
    pagination,
    getData,
    searchParams,
    resetSearchParams,
    handleSizeChange,
    handleCurrentChange,
    refreshData
  } = useTable({
    // 核心配置
    core: {
      apiFn: fetchGetRoleList,
      apiParams: {
        current: 1,
        size: 20,
        ...searchForm.value // 添加这一行
      },
      columnsFactory: () => [
        {
          prop: 'id',
          label: '角色ID',
          width: 100
        },
        {
          prop: 'displayName',
          label: '角色显示名称',
          minWidth: 120
        },
        {
          prop: 'name',
          label: '角色名称',
          minWidth: 120
        },
        {
          prop: 'description',
          label: '角色描述',
          minWidth: 150,
          showOverflowTooltip: true
        },
        {
          prop: 'status',
          label: '角色状态',
          width: 100,
          formatter: (row) => {
            const statusConfig =
              row.status === 'active'
                ? { type: 'success', text: '启用' }
                : { type: 'warning', text: '禁用' }
            return h(
              ElTag,
              { type: statusConfig.type as 'success' | 'warning' },
              () => statusConfig.text
            )
          }
        },
        {
          prop: 'createdAt',
          label: '创建日期',
          width: 180,
          sortable: true
        },
        {
          prop: 'operation',
          label: '操作',
          width: 80,
          fixed: 'right',
          formatter: (row) =>
            h('div', [
              h(ArtButtonMore, {
                list: [
                  {
                    key: 'permission',
                    label: '菜单权限',
                    icon: Setting
                  },
                  {
                    key: 'edit',
                    label: '编辑角色',
                    icon: Edit
                  },
                  {
                    key: 'delete',
                    label: '删除角色',
                    icon: Delete,
                    color: '#f56c6c'
                  }
                ],
                onClick: (item: ButtonMoreItem) => buttonMoreClick(item, row)
              })
            ])
        }
      ]
    }
  })

  const dialogType = ref<'add' | 'edit'>('add')

  const showDialog = (type: 'add' | 'edit', row?: RoleListItem) => {
    dialogVisible.value = true
    dialogType.value = type
    currentRoleData.value = row
  }

  /**
   * 搜索处理
   * @param params 搜索参数
   */
  const handleSearch = (params: Record<string, any>) => {
    // 搜索参数赋值
    Object.assign(searchParams, params)
    getData()
  }

  const buttonMoreClick = (item: ButtonMoreItem, row: RoleListItem) => {
    switch (item.key) {
      case 'permission':
        showPermissionDialog(row)
        break
      case 'edit':
        showDialog('edit', row)
        break
      case 'delete':
        deleteRole(row)
        break
    }
  }

  const showPermissionDialog = (row?: RoleListItem) => {
    permissionDialog.value = true
    currentRoleData.value = row
  }

  const deleteRole = async (row: RoleListItem) => {
    try {
      await ElMessageBox.confirm(
        `确定删除角色"${row.displayName}"吗？此操作不可恢复！`,
        '删除确认',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      )

      await fetchDeleteRole(row.id)
      ElMessage.success('删除成功')
      refreshData()
    } catch (error: any) {
      if (error !== 'cancel') {
        console.error('删除失败:', error)
        ElMessage.error(error?.message || '删除失败，请重试')
      } else {
        ElMessage.info('已取消删除')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .role-page {
    padding-bottom: 15px;
  }
</style>
