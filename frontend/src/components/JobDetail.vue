<script setup lang="ts">
import { ref } from 'vue'
import type { Job, Candidate } from '@/types'
import { formatSalary } from '@/utils/status'

const props = defineProps<{
  job: Job | null
  currentCandidate: Candidate | null
  isRecruiter?: boolean
}>()

const emit = defineEmits<{
  close: []
  apply: [jobId: string, resumeSummary: string]
  edit: [job: Job]
  delete: [jobId: string]
}>()

const showApplyForm = ref(false)
const resumeSummary = ref('')
const applyError = ref('')

function handleApply() {
  if (!props.job) return
  const summary = resumeSummary.value.trim()
  if (!summary) {
    applyError.value = '请填写简历摘要后再投递'
    return
  }
  applyError.value = ''
  emit('apply', props.job.id, summary)
}

function toggleApplyForm() {
  showApplyForm.value = !showApplyForm.value
  if (!showApplyForm.value) {
    applyError.value = ''
    resumeSummary.value = ''
  }
}
</script>

<template>
  <div v-if="job" class="job-detail">
    <div class="detail-header">
      <div>
        <h2>{{ job.title }}</h2>
        <div class="job-meta">
          <span>{{ job.company }}</span>
          <span>·</span>
          <span>{{ job.location }}</span>
          <span>·</span>
          <span>{{ formatSalary(job.salary_min, job.salary_max) }}</span>
        </div>
      </div>
      <div class="header-actions">
        <div v-if="isRecruiter" class="action-buttons">
          <button class="btn btn-outline btn-sm" @click="emit('edit', job)">编辑</button>
          <button class="btn btn-danger btn-sm" @click="emit('delete', job.id)">删除</button>
        </div>
        <button class="modal-close" @click="emit('close')">&times;</button>
      </div>
    </div>

    <div class="detail-body">
      <div class="detail-section">
        <h3>职位描述</h3>
        <p>{{ job.description }}</p>
      </div>

      <div class="detail-section">
        <h3>任职要求</h3>
        <p>{{ job.requirements }}</p>
      </div>

      <div class="detail-section">
        <h3>职位状态</h3>
        <span class="badge" :class="`badge-${job.status}`">{{ job.status === 'open' ? '开放' : '已关闭' }}</span>
      </div>
    </div>

    <div class="detail-footer" v-if="!isRecruiter">
      <button
        v-if="job.status === 'open'"
        class="btn btn-primary"
        @click="toggleApplyForm"
      >
        {{ showApplyForm ? '收起投递' : '立即投递' }}
      </button>
      <span v-else class="text-gray">该职位已关闭，无法投递</span>
    </div>

    <div v-if="showApplyForm && job.status === 'open' && !isRecruiter" class="apply-form">
      <h3>填写简历摘要</h3>
      <textarea
        v-model="resumeSummary"
        class="form-textarea"
        placeholder="请简要介绍您的工作经历、技能特长等..."
      ></textarea>
      <div v-if="applyError" class="error-message">{{ applyError }}</div>
      <div class="flex justify-end gap-2 mt-4">
        <button class="btn btn-outline" @click="toggleApplyForm">取消</button>
        <button class="btn btn-success" @click="handleApply">确认投递</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.job-detail {
  padding: 20px;
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 24px;
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.action-buttons {
  display: flex;
  gap: 8px;
}

.btn-sm {
  padding: 4px 8px;
  font-size: 12px;
}

.job-meta {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  font-size: 14px;
  color: #666;
}

.detail-body {
  margin-bottom: 24px;
}

.detail-section {
  margin-bottom: 20px;
}

.detail-section h3 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 12px;
}

.detail-section p {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
}

.detail-footer {
  padding-top: 16px;
  border-top: 1px solid #f0f0f0;
}

.apply-form {
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

.apply-form h3 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 12px;
}
</style>
