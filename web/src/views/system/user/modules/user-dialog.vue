<template>
  <ElDialog
    v-model="dialogVisible"
    :title="dialogType === 'add' ? '添加用户' : '编辑用户'"
    width="30%"
    align-center
  >
    <ElForm ref="formRef" :model="formData" :rules="rules" label-width="80px">
      <ElFormItem label="用户名" prop="username">
        <ElInput v-model="formData.username" placeholder="请输入用户名" />
      </ElFormItem>
      <!-- 昵称字段 -->
      <ElFormItem label="昵称" prop="nickname">
        <ElInput v-model="formData.nickname" placeholder="请输入昵称" />
      </ElFormItem>
      <!-- 密码字段：添加时显示且必填，编辑时隐藏 -->
      <ElFormItem v-if="dialogType === 'add'" label="密码" prop="password">
        <ElInput
          v-model="formData.password"
          type="password"
          placeholder="请输入密码"
          show-password
        />
      </ElFormItem>
      <ElFormItem label="手机号" prop="phone">
        <ElInput v-model="formData.phone" placeholder="请输入手机号" />
      </ElFormItem>
      <ElFormItem label="邮箱" prop="email">
        <ElInput v-model="formData.email" placeholder="请输入邮箱" />
      </ElFormItem>
      <ElFormItem label="角色" prop="role_id">
        <ElSelect v-model="formData.role_id">
          <ElOption
            v-for="role in roleList"
            :key="role.id"
            :value="role.id"
            :label="role.displayName"
          />
        </ElSelect>
      </ElFormItem>
      <!--状态-->
      <ElFormItem label="状态" prop="status">
        <ElRadioGroup v-model="formData.status">
          <ElRadio label="active">正常</ElRadio>
          <ElRadio label="disabled">禁用</ElRadio>
        </ElRadioGroup>
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="dialog-footer">
        <ElButton @click="dialogVisible = false">取消</ElButton>
        <ElButton type="primary" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import { fetchGetEnabledRoleList, fetchCreateUser, fetchUpdateUser } from '@/api/system-manage'
  import type { FormInstance, FormRules } from 'element-plus'

  interface Props {
    visible: boolean
    type: string
    userData?: any
  }

  interface Emits {
    (e: 'update:visible', value: boolean): void
    (e: 'submit'): void
  }

  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  // 角色列表数据
  const roleList = ref<Api.SystemManage.ActiveRoleListItem[]>([])

  // 对话框显示控制
  const dialogVisible = computed({
    get: () => props.visible,
    set: (value) => emit('update:visible', value)
  })

  const dialogType = computed(() => props.type)

  // 表单实例
  const formRef = ref<FormInstance>()

  // 表单数据
  const formData = reactive({
    username: '',
    nickname: '',
    phone: '',
    password: '',
    email: '',
    role_id: 0,
    status: 'active'
  })

  // 动态表单验证规则
  const rules = computed<FormRules>(() => {
    const baseRules: FormRules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
      ],
      phone: [
        { required: false, message: '请输入手机号', trigger: 'blur' },
        { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号格式', trigger: 'blur' }
      ],
      email: [
        { required: false, message: '请输入邮箱', trigger: 'blur' },
        {
          pattern: /^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$/,
          message: '请输入正确的邮箱格式',
          trigger: 'blur'
        }
      ],
      role_id: [{ required: true, message: '请选择角色', trigger: 'blur' }]
    }

    // 添加时密码为必填
    if (dialogType.value === 'add') {
      baseRules.password = [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
      ]
    }

    return baseRules
  })

  // 初始化表单数据
  const initFormData = () => {
    const isEdit = props.type === 'edit' && props.userData
    const row = props.userData

    Object.assign(formData, {
      id: isEdit ? row.id || '' : '',
      username: isEdit ? row.username || '' : '',
      nickname: isEdit ? row.nickname || '' : '',
      phone: isEdit ? row.phone || '' : '',
      password: '', // 编辑时不显示密码，添加时为空
      email: isEdit ? row.email || '' : '',
      role_id: isEdit ? row.roleId || '' : '',
      status: isEdit ? row.status || 'active' : 'active'
    })
  }

  // 统一监听对话框状态变化
  watch(
    () => [props.visible, props.type, props.userData],
    ([visible]) => {
      if (visible) {
        initFormData()
        nextTick(() => {
          formRef.value?.clearValidate()
        })
      }
    },
    { immediate: true }
  )

  // 提交表单
  const handleSubmit = async () => {
    if (!formRef.value) return

    try {
      const valid = await formRef.value.validate()
      if (valid) {
        // 准备提交数据
        const submitData = { ...formData }

        if (dialogType.value === 'add') {
          // 创建用户
          await fetchCreateUser({
            username: submitData.username,
            password: submitData.password,
            nickname: submitData.nickname,
            phone: submitData.phone,
            email: submitData.email,
            role_id: Number(submitData.role_id),
            status: submitData.status as Api.Common.EnableStatus
          })
          ElMessage.success('用户创建成功')
        } else {
          // 更新用户 - 编辑时移除密码字段
          await fetchUpdateUser({
            id: Number(props.userData?.id),
            username: submitData.username,
            nickname: submitData.nickname,
            phone: submitData.phone,
            email: submitData.email,
            role_id: Number(submitData.role_id),
            status: submitData.status as Api.Common.EnableStatus
          })
          ElMessage.success('用户更新成功')
        }

        dialogVisible.value = false
        emit('submit')
      }
    } catch (error: any) {
      console.error('操作失败:', error)
    }
  }

  // 获取角色列表
  const getRoleList = async () => {
    try {
      const data = await fetchGetEnabledRoleList()
      console.log('角色列表', data)
      roleList.value = data.list || []
    } catch (error) {
      console.error('获取角色列表失败:', error)
      ElMessage.error('获取角色列表失败')
    }
  }

  // 组件挂载时获取角色列表
  onMounted(() => {
    getRoleList()
  })
</script>
