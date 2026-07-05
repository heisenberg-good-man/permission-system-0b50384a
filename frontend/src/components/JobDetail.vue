<script setup lang="ts">
import { ref } from 'vue'
import type { Job, Candidate } from '@/types'

defineProps<{
  job: Job | null
  currentCandidate: Candidate | null
}>()

const emit = defineEmits<{
  close: []
  apply: [jobId: string]
}>()

const showApplyForm = ref(false)
const resumeSummary = ref('')

function formatSalary(min: number, max: number): string {
  return `${(min / 10000).toFixed(1)}K - ${(max / 10000).toFixed(1)}K`
}

function handleApply() {
  if (resumeSummary.value.trim()) {
    emit('apply', '')
    showApplyForm.value = false
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
      <button class="modal-close" @click="emit('close')">&times;</button>
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
    
    <div class="detail-footer">
      <button
        v-if="job.status === 'open'"
        class="btn btn-primary"
        @click="showApplyForm = !showApplyForm"
      >
        {{ showApplyForm ? '取消投递' : '立即投递' }}
      </button>
      <span v-else class="text-gray">该职位已关闭，无法投递</span>
    </div>
    
    <div v-if="showApplyForm && job.status === 'open'" class="apply-form">
      <h3>填写简历摘要</h3>
      <textarea
        v-model="resumeSummary"
        class="form-textarea"
        placeholder="请简要介绍您的工作经历、技能特长等..."
      ></textarea>
      <div class="flex justify-end gap-2 mt-4">
        <button class="btn btn-outline" @click="showApplyForm = false">取消</button>
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