<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Job } from '@/types'

const props = defineProps<{
  job?: Job | null
}>()

const emit = defineEmits<{
  submit: [job: Omit<Job, 'id' | 'created_at' | 'updated_at'> | Partial<Job>]
  cancel: []
}>()

const form = ref({
  title: '',
  company: '',
  location: '',
  salary_min: 0,
  salary_max: 0,
  description: '',
  requirements: '',
  status: 'open',
  recruiter_id: 'recruiter-1'
})

watch(() => props.job, (newJob) => {
  if (newJob) {
    form.value = {
      title: newJob.title,
      company: newJob.company,
      location: newJob.location,
      salary_min: newJob.salary_min,
      salary_max: newJob.salary_max,
      description: newJob.description,
      requirements: newJob.requirements,
      status: newJob.status,
      recruiter_id: newJob.recruiter_id
    }
  }
}, { immediate: true })

const errors = ref<Record<string, string>>({})

function validate(): boolean {
  errors.value = {}
  if (!form.value.title.trim()) {
    errors.value.title = '职位名称为必填项'
  }
  if (!form.value.company.trim()) {
    errors.value.company = '公司名称为必填项'
  }
  if (!form.value.description.trim()) {
    errors.value.description = '职位描述为必填项'
  }
  if (form.value.salary_min > form.value.salary_max) {
    errors.value.salary = '最低薪资不能大于最高薪资'
  }
  return Object.keys(errors.value).length === 0
}

function handleSubmit() {
  if (validate()) {
    emit('submit', { ...form.value })
  }
}
</script>

<template>
  <div class="job-form">
    <div class="form-group">
      <label class="form-label">职位名称 *</label>
      <input v-model="form.title" type="text" class="form-input" placeholder="请输入职位名称" />
      <span v-if="errors.title" class="error-message">{{ errors.title }}</span>
    </div>
    
    <div class="form-group">
      <label class="form-label">公司名称 *</label>
      <input v-model="form.company" type="text" class="form-input" placeholder="请输入公司名称" />
      <span v-if="errors.company" class="error-message">{{ errors.company }}</span>
    </div>
    
    <div class="form-group">
      <label class="form-label">工作地点</label>
      <input v-model="form.location" type="text" class="form-input" placeholder="请输入工作地点" />
    </div>
    
    <div class="flex gap-4">
      <div class="form-group" style="flex: 1">
        <label class="form-label">最低薪资（元）</label>
        <input v-model.number="form.salary_min" type="number" class="form-input" placeholder="0" />
      </div>
      <div class="form-group" style="flex: 1">
        <label class="form-label">最高薪资（元）</label>
        <input v-model.number="form.salary_max" type="number" class="form-input" placeholder="0" />
      </div>
    </div>
    <span v-if="errors.salary" class="error-message">{{ errors.salary }}</span>
    
    <div class="form-group">
      <label class="form-label">职位描述 *</label>
      <textarea v-model="form.description" class="form-textarea" placeholder="请输入职位描述"></textarea>
      <span v-if="errors.description" class="error-message">{{ errors.description }}</span>
    </div>
    
    <div class="form-group">
      <label class="form-label">任职要求</label>
      <textarea v-model="form.requirements" class="form-textarea" placeholder="请输入任职要求"></textarea>
    </div>
    
    <div class="form-group">
      <label class="form-label">状态</label>
      <select v-model="form.status" class="form-select">
        <option value="open">开放</option>
        <option value="closed">已关闭</option>
      </select>
    </div>
    
    <div class="flex justify-end gap-2 mt-4">
      <button class="btn btn-outline" @click="emit('cancel')">取消</button>
      <button class="btn btn-primary" @click="handleSubmit">保存</button>
    </div>
  </div>
</template>

<style scoped>
.job-form {
  padding: 20px;
}
</style>