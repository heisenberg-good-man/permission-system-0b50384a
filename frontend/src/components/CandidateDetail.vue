<script setup lang="ts">
import type { Candidate } from '@/types'

defineProps<{
  candidate: Candidate | null
}>()

const emit = defineEmits<{
  close: []
}>()
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
</style>