<template>
  <ElDialog
    v-model="visible"
    :title="dialogType === 'add' ? '新增角色' : '编辑角色'"
    width="30%"
    align-center
    @close="handleClose"
  >
    <ElForm ref="formRef" :model="form" :rules="rules" label-width="120px">
      <ElFormItem label="角色显示名称" prop="displayName">
        <ElInput v-model="form.displayName" placeholder="请输入角色显示名称" />
      </ElFormItem>
      <ElFormItem label="角色名称" prop="name">
        <ElInput v-model="form.name" placeholder="请输入角色名称（英文唯一标识）" />
      </ElFormItem>
      <ElFormItem label="描述" prop="description">
        <ElInput
          v-model="form.description"
          type="textarea"
          :rows="3"
          placeholder="请输入角色描述"
        />
      </ElFormItem>
      <ElFormItem label="启用">
        <ElSwitch  
          v-model="form.status"  
          active-value="active"
          inactive-value="disabled"
        />
      </ElFormItem>
    </ElForm>
    <template #footer>
      <div class="dialog-footer">
        <ElButton @click="handleClose">取消</ElButton>
        <ElButton type="primary" :loading="loading" @click="handleSubmit">提交</ElButton>
      </div>
    </template>
  </ElDialog>
</template>

<script setup lang="ts">
  import type { FormInstance, FormRules } from 'element-plus'
  import { fetchCreateRole, fetchUpdateRole } from '@/api/system-manage'

  type RoleListItem = Api.SystemManage.RoleListItem

  interface Props {
    modelValue: boolean
    dialogType: 'add' | 'edit'
    roleData?: RoleListItem
  }

  interface Emits {
    (e: 'update:modelValue', value: boolean): void
    (e: 'success'): void
  }

  const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
    dialogType: 'add',
    roleData: undefined
  })

  const emit = defineEmits<Emits>()

  const visible = computed({
    get: () => props.modelValue,
    set: (value) => emit('update:modelValue', value)
  })

  const formRef = ref<FormInstance>()
  const loading = ref(false)

  const rules = reactive<FormRules>({
    displayName: [
      { required: true, message: '请输入角色显示名称', trigger: 'blur' },
      { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
    ],
    name: [
      { required: true, message: '请输入角色名称', trigger: 'blur' },
      { min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur' },
      { pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/, message: '角色名称必须以字母开头，只能包含字母、数字和下划线', trigger: 'blur' }
    ],
    description: [{ required: true, message: '请输入角色描述', trigger: 'blur' }]
  })

  const form = reactive({
    id: 0,
    name: '',
    displayName: '',
    description: '',
    status: 'active' as 'active' | 'disabled'
  })

  // 监听弹窗打开，初始化表单数据
  watch(
    () => props.modelValue,
    (newVal) => {
      if (newVal) {
        initForm()
      }
    },
    { immediate: true }
  )

  // 监听角色数据变化
  watch(
    () => props.roleData,
    (newData) => {
      if (newData && props.modelValue) {
        initForm()
      }
    },
    { deep: true }
  )

  const initForm = () => {
    if (props.dialogType === 'edit' && props.roleData) {
      Object.assign(form, {
        id: props.roleData.id,
        name: props.roleData.name,
        displayName: props.roleData.displayName,
        description: props.roleData.description,
        status: props.roleData.status
      })
    } else {
      Object.assign(form, {
        id: 0,
        name: '',
        displayName: '',
        description: '',
        status: 'active'
      })
    }
  }

  const handleClose = () => {
    visible.value = false
    formRef.value?.resetFields()
    loading.value = false
  }

  const handleSubmit = async () => {
    if (!formRef.value) return

    try {
      await formRef.value.validate()
      loading.value = true

      if (props.dialogType === 'add') {
        // 新增角色
        await fetchCreateRole({
          name: form.name,
          displayName: form.displayName,
          description: form.description,
          status: form.status
        })
        ElMessage.success('新增成功')
      } else {
        // 编辑角色
        await fetchUpdateRole({
          id: form.id,
          name: form.name,
          displayName: form.displayName,
          description: form.description,
          status: form.status
        })
        ElMessage.success('修改成功')
      }
      
      emit('success')
      handleClose()
    } catch (error: any) {
      console.error('操作失败:', error)
      ElMessage.error(error?.message || '操作失败，请重试')
    } finally {
      loading.value = false
    }
  }
</script>

<style lang="scss" scoped>
  .dialog-footer {
    display: flex;
    gap: 12px;
    justify-content: flex-end;
  }
</style>
