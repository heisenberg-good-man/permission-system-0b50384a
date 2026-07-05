<script setup lang="ts">
import type { Application, Candidate, Job } from '@/types'

defineProps<{
  applications: Application[]
  candidates: Map<string, Candidate>
  jobs: Map<string, Job>
}>()

const emit = defineEmits<{
  viewCandidate: [candidateId: string]
  viewMessage: [application: Application]
  updateStatus: [applicationId: string, status: string]
}>()

function formatDate(dateStr: string): string {
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

function getStatusText(status: string): string {
  const map: Record<string, string> = {
    pending: '待筛选',
    communicating: '沟通中',
    interviewing: '面试中',
    offered: '已发offer',
    offer_accepted: '已接受offer',
    rejected: '已拒绝'
  }
  return map[status] || status
}
</script>

<template>
  <div class="application-list">
    <table class="table">
      <thead>
        <tr>
          <th>候选人</th>
          <th>职位</th>
          <th>状态</th>
          <th>投递时间</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="app in applications" :key="app.id">
          <td>
            <div class="candidate-info">
              <div class="candidate-name">{{ candidates.get(app.candidate_id)?.name || '-' }}</div>
              <div class="candidate-email">{{ candidates.get(app.candidate_id)?.email || '-' }}</div>
            </div>
          </td>
          <td>{{ jobs.get(app.job_id)?.title || '-' }}</td>
          <td>
            <span class="badge" :class="`badge-${app.status}`">{{ getStatusText(app.status) }}</span>
          </td>
          <td>{{ formatDate(app.applied_at) }}</td>
          <td>
            <div class="action-buttons">
              <button class="btn btn-outline btn-sm" @click="emit('viewCandidate', app.candidate_id)">查看详情</button>
              <button v-if="app.status === 'communicating'" class="btn btn-primary btn-sm" @click="emit('viewMessage', app)">消息</button>
              <select class="status-select" @change="(e) => emit('updateStatus', app.id, (e.target as HTMLSelectElement).value)">
                <option :value="app.status" disabled>{{ getStatusText(app.status) }}</option>
                <option value="pending">待筛选</option>
                <option value="communicating">沟通中</option>
                <option value="interviewing">面试中</option>
                <option value="offered">已发offer</option>
                <option value="offer_accepted">已接受offer</option>
                <option value="rejected">已拒绝</option>
              </select>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="applications.length === 0" class="empty-state">
      暂无投递记录
    </div>
  </div>
</template>

<style scoped>
.application-list {
  overflow-x: auto;
}

.candidate-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.candidate-name {
  font-weight: 500;
}

.candidate-email {
  font-size: 12px;
  color: #999;
}

.action-buttons {
  display: flex;
  gap: 8px;
  align-items: center;
}

.btn-sm {
  padding: 4px 8px;
  font-size: 12px;
}

.status-select {
  padding: 4px 8px;
  font-size: 12px;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  cursor: pointer;
}

.empty-state {
  text-align: center;
  padding: 40px;
  color: #999;
}
</style>