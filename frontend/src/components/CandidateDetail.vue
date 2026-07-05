<script setup lang="ts">
import { computed } from 'vue'
import type { Candidate, Application, Job, StatusHistory } from '@/types'

const props = defineProps<{
  candidate: Candidate | null
  applications: Application[]
  jobs: Map<string, Job>
}>()

const emit = defineEmits<{
  close: []
  viewMessage: [application: Application]
  updateStatus: [applicationId: string, status: string]
}>()

function formatDate(dateStr: string): string {
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
}

function getStatusText(status: string): string {
  const map: Record<string, string> = {
    pending: '待筛选',
    communicating: '沟通中',
    interviewing: '面试中',
    offered: '已发offer',
    rejected: '已拒绝'
  }
  return map[status] || status
}

function getStatusHistoryText(history: StatusHistory[]): string {
  if (!history || history.length === 0) return '暂无记录'
  return history.map(h => `${getStatusText(h.status)} (${formatDate(h.changed_at)})`).join(' → ')
}

const candidateApplications = computed(() => {
  if (!props.candidate) return []
  return props.applications.filter(a => a.candidate_id === props.candidate!.id)
})
</script>

<template>
  <div v-if="candidate" class="candidate-detail">
    <div class="detail-header">
      <h3>{{ candidate.name }}</h3>
      <button class="modal-close" @click="emit('close')">&times;</button>
    </div>
    
    <div class="detail-section">
      <div class="section-title">基本信息</div>
      <div class="info-grid">
        <div class="info-item">
          <span class="info-label">邮箱：</span>
          <span>{{ candidate.email }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">电话：</span>
          <span>{{ candidate.phone }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">学历：</span>
          <span>{{ candidate.education }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">经验：</span>
          <span>{{ candidate.experience }}年</span>
        </div>
      </div>
    </div>
    
    <div class="detail-section">
      <div class="section-title">技能标签</div>
      <div class="skills">
        <span v-for="skill in candidate.skills" :key="skill" class="skill-tag">{{ skill }}</span>
      </div>
    </div>
    
    <div class="detail-section">
      <div class="section-title">简历摘要</div>
      <div class="resume-summary">{{ candidate.resume_summary }}</div>
    </div>

    <div class="detail-section">
      <div class="section-title">投递记录</div>
      <div v-if="candidateApplications.length > 0" class="applications-list">
        <div v-for="app in candidateApplications" :key="app.id" class="application-row">
          <div class="app-info">
            <div class="app-job-title">{{ jobs.get(app.job_id)?.title || '-' }}</div>
            <div class="app-meta">投递时间：{{ formatDate(app.applied_at) }}</div>
          </div>
          <div class="app-status">
            <span class="badge" :class="`badge-${app.status}`">{{ getStatusText(app.status) }}</span>
          </div>
          <div class="app-history">
            {{ getStatusHistoryText(app.status_history || []) }}
          </div>
          <div class="app-actions">
            <button 
              v-if="app.status === 'communicating'" 
              class="btn btn-primary btn-sm" 
              @click="emit('viewMessage', app)"
            >
              消息
            </button>
            <select class="status-select" @change="(e) => emit('updateStatus', app.id, (e.target as HTMLSelectElement).value)">
              <option :value="app.status" disabled>{{ getStatusText(app.status) }}</option>
              <option value="pending">待筛选</option>
              <option value="communicating">沟通中</option>
              <option value="interviewing">面试中</option>
              <option value="offered">已发offer</option>
              <option value="rejected">已拒绝</option>
            </select>
          </div>
        </div>
      </div>
      <div v-else class="empty-state">暂无投递记录</div>
    </div>
  </div>
</template>

<style scoped>
.candidate-detail {
  padding: 20px;
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.detail-section {
  margin-bottom: 24px;
}

.section-title {
  font-size: 14px;
  font-weight: 600;
  color: #333;
  margin-bottom: 12px;
}

.info-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.info-item {
  font-size: 14px;
  color: #666;
}

.info-label {
  color: #999;
}

.skills {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.skill-tag {
  padding: 4px 12px;
  background-color: #f0f5ff;
  color: #409eff;
  border-radius: 4px;
  font-size: 13px;
}

.resume-summary {
  font-size: 14px;
  color: #666;
  line-height: 1.6;
  background-color: #fafafa;
  padding: 16px;
  border-radius: 4px;
}

.applications-list {
  margin-top: 8px;
}

.application-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px;
  border: 1px solid #f0f0f0;
  border-radius: 8px;
  margin-bottom: 8px;
}

.app-info {
  flex: 1;
}

.app-job-title {
  font-weight: 600;
  font-size: 14px;
}

.app-meta {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}

.app-status {
  margin-right: 16px;
}

.app-history {
  font-size: 12px;
  color: #999;
  margin-right: 16px;
  max-width: 300px;
}

.app-actions {
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
  padding: 20px;
  color: #999;
}
</style>