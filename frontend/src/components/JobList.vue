<script setup lang="ts">
import type { Job } from '@/types'

defineProps<{
  jobs: Job[]
  showStatus?: boolean
}>()

const emit = defineEmits<{
  viewDetail: [job: Job]
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
    <div v-for="job in jobs" :key="job.id" class="job-card" @click="emit('viewDetail', job)">
      <div class="job-header">
        <div class="job-title">{{ job.title }}</div>
        <div v-if="showStatus" class="badge" :class="`badge-${job.status}`">
          {{ job.status === 'open' ? '开放' : '已关闭' }}
        </div>
      </div>
      <div class="job-company">{{ job.company }}</div>
      <div class="job-meta">
        <span class="meta-item">{{ job.location }}</span>
        <span class="meta-item">{{ formatSalary(job.salary_min, job.salary_max) }}</span>
        <span class="meta-item">{{ formatDate(job.created_at) }}</span>
      </div>
      <div class="job-description">{{ job.description }}</div>
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