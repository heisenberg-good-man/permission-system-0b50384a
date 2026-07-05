<script setup lang="ts">
import type { Job } from '@/types'

defineProps<{
  jobs: Job[]
  showStatus?: boolean
  showActions?: boolean
}>()

const emit = defineEmits<{
  viewDetail: [job: Job]
  edit: [job: Job]
  delete: [jobId: string]
}>()

function formatSalary(min: number, max: number): string {
  return `${(min / 10000).toFixed(1)}K - ${(max / 10000).toFixed(1)}K`
}

function formatDate(dateStr: string): string {
  const date = new Date(dateStr)
  return `${date.getMonth() + 1}/${date.getDate()}`
}
</script>

<template>
  <div class="job-list">
    <div v-for="job in jobs" :key="job.id" class="job-card">
      <div class="job-header">
        <div class="job-title" @click="emit('viewDetail', job)">{{ job.title }}</div>
        <div class="job-header-right">
          <div v-if="showActions" class="job-actions">
            <button class="btn btn-outline btn-sm" @click.stop="emit('edit', job)">编辑</button>
            <button class="btn btn-danger btn-sm" @click.stop="emit('delete', job.id)">删除</button>
          </div>
          <div v-if="showStatus" class="badge" :class="`badge-${job.status}`">
            {{ job.status === 'open' ? '开放' : '已关闭' }}
          </div>
        </div>
      </div>
      <div class="job-company" @click="emit('viewDetail', job)">{{ job.company }}</div>
      <div class="job-meta">
        <span class="meta-item">{{ job.location }}</span>
        <span class="meta-item">{{ formatSalary(job.salary_min, job.salary_max) }}</span>
        <span class="meta-item">{{ formatDate(job.created_at) }}</span>
      </div>
      <div class="job-description" @click="emit('viewDetail', job)">{{ job.description }}</div>
    </div>
  </div>
</template>

<style scoped>
.job-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.job-card {
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  cursor: pointer;
  transition: box-shadow 0.2s;
}

.job-card:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.12);
}

.job-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.job-header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.job-actions {
  display: flex;
  gap: 8px;
}

.job-title {
  font-size: 18px;
  font-weight: 600;
  color: #333;
}

.job-company {
  font-size: 14px;
  color: #666;
  margin-bottom: 12px;
}

.job-meta {
  display: flex;
  gap: 16px;
  margin-bottom: 12px;
}

.meta-item {
  font-size: 13px;
  color: #999;
}

.job-description {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>