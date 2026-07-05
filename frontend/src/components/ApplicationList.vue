<script setup lang="ts">
import type { Application, Candidate, Job } from '@/types'
import {
  formatDate,
  formatDateTime,
  getStatusText,
  getStatusBadgeClass,
  canSendOffer,
  canViewMessages
} from '@/utils/status'

defineProps<{
  applications: Application[]
  candidates: Map<string, Candidate>
  jobs: Map<string, Job>
}>()

const emit = defineEmits<{
  viewCandidate: [candidateId: string]
  viewMessage: [application: Application]
  updateStatus: [applicationId: string, status: string]
  sendOffer: [applicationId: string]
}>()

const STATUS_OPTIONS = [
  { value: 'pending', label: '待筛选' },
  { value: 'communicating', label: '沟通中' },
  { value: 'interviewing', label: '面试中' },
  { value: 'offered', label: '已发offer' },
  { value: 'rejected', label: '已拒绝' }
]
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
              <button class="candidate-name-btn" @click="emit('viewCandidate', app.candidate_id)">
                {{ candidates.get(app.candidate_id)?.name || '-' }}
              </button>
              <div class="candidate-email">{{ candidates.get(app.candidate_id)?.email || '-' }}</div>
            </div>
          </td>
          <td>{{ jobs.get(app.job_id)?.title || '-' }}</td>
          <td>
            <span class="badge" :class="getStatusBadgeClass(app.status)">{{ getStatusText(app.status) }}</span>
            <div v-if="app.offer_detail && (app.status === 'offered' || app.status === 'offer_accepted')" class="offer-mini">
              {{ app.offer_detail.offer_salary_min }}-{{ app.offer_detail.offer_salary_max }} 元
            </div>
          </td>
          <td>
            <div>{{ formatDate(app.applied_at) }}</div>
            <div v-if="app.updated_at" class="updated-hint">更新于 {{ formatDateTime(app.updated_at) }}</div>
          </td>
          <td>
            <div class="action-buttons">
              <button class="btn btn-outline btn-sm" @click="emit('viewCandidate', app.candidate_id)">查看详情</button>
              <button
                v-if="canViewMessages(app.status)"
                class="btn btn-primary btn-sm"
                @click="emit('viewMessage', app)"
              >
                消息
              </button>
              <button
                v-if="canSendOffer(app.status)"
                class="btn btn-success btn-sm"
                @click="emit('sendOffer', app.id)"
              >
                发送Offer
              </button>
              <select
                :value="app.status"
                class="status-select"
                @change="(e) => emit('updateStatus', app.id, (e.target as HTMLSelectElement).value)"
              >
                <option value="" disabled>切换状态</option>
                <option v-for="opt in STATUS_OPTIONS" :key="opt.value" :value="opt.value">{{ opt.label }}</option>
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

.candidate-name-btn {
  background: none;
  border: none;
  padding: 0;
  font-weight: 500;
  color: #409eff;
  cursor: pointer;
  font-size: 14px;
  text-align: left;
}

.candidate-name-btn:hover {
  text-decoration: underline;
}

.candidate-email {
  font-size: 12px;
  color: #999;
}

.offer-mini {
  font-size: 11px;
  color: #67c23a;
  margin-top: 4px;
}

.updated-hint {
  font-size: 11px;
  color: #bbb;
  margin-top: 2px;
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

.btn-success {
  background-color: #67c23a;
  color: white;
  border: none;
}

.btn-success:hover {
  background-color: #5eb838;
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
