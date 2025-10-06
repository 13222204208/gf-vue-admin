<template>
  <ArtSearchBar
    ref="searchBarRef"
    v-model="formData"
    :items="formItems"
    :rules="rules"
    @reset="handleReset"
    @search="handleSearch"
  >
  </ArtSearchBar>
</template>

<script setup lang="ts">
  interface Props {
    modelValue: Record<string, any>
  }
  interface Emits {
    (e: 'update:modelValue', value: Record<string, any>): void
    (e: 'search', params: Record<string, any>): void
    (e: 'reset'): void
  }
  const props = defineProps<Props>()
  const emit = defineEmits<Emits>()

  // 表单数据双向绑定
  const searchBarRef = ref()
  const formData = computed({
    get: () => props.modelValue,
    set: (val) => emit('update:modelValue', val)
  })

  // 校验规则
  const rules = {
    // name: [{ required: true, message: '请输入用户名', trigger: 'blur' }]
  }

  // 角色状态选项
  const statusOptions = ref([
    { label: '启用', value: 'active' },
    { label: '禁用', value: 'disabled' }
  ])

  // 表单配置
  const formItems = computed(() => [
    {
      label: '显示名称',
      key: 'displayName',
      type: 'input',
      placeholder: '请输入角色显示名称',
      clearable: true
    },
    {
      label: '角色名称',
      key: 'name',
      type: 'input',
      placeholder: '请输入角色名称',
      clearable: true
    },

    {
      label: '角色状态',
      key: 'status',
      type: 'select',
      props: {
        placeholder: '请选择状态',
        options: statusOptions.value,
        clearable: true
      }
    }
  ])

  // 事件
  function handleReset() {
    console.log('重置表单')
    emit('reset')
  }

  async function handleSearch() {
    await searchBarRef.value.validate()
    emit('search', formData.value)
    console.log('表单数据', formData.value)
  }
</script>
