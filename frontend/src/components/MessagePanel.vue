<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import type { Message, Application, Job, Candidate } from '@/types'
import { formatTime, getStatusText, getStatusBadgeClass } from '@/utils/status'

const props = defineProps<{
  messages: Message[]
  application: Application | null
  job: Job | null
  candidate: Candidate | null
  disabled?: boolean
}>()

const emit = defineEmits<{
  send: [content: string]
  close: []
}>()

const messageInput = ref('')
const messagesContainer = ref<HTMLElement | null>(null)

watch(() => props.messages, async () => {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}, { deep: true })

function handleSend() {
  if (props.disabled) return
  const content = messageInput.value.trim()
  if (content) {
    emit('send', content)
    messageInput.value = ''
  }
}

function isRecruiterMessage(senderType: string): boolean {
  return senderType === 'recruiter'
}
</script>

<template>
  <div class="message-panel">
    <div class="panel-header">
      <div class="header-info">
        <div class="conversation-title">
          {{ job?.title || '-' }} - {{ candidate?.name || '-' }}
        </div>
        <div class="conversation-status">
          <span v-if="application" class="badge" :class="getStatusBadgeClass(application.status)">
            {{ getStatusText(application.status) }}
          </span>
          <span v-else class="text-gray">未选择投递</span>
        </div>
      </div>
      <button class="modal-close" @click="emit('close')">&times;</button>
    </div>

    <div ref="messagesContainer" class="messages-container">
      <div
        v-for="msg in messages"
        :key="msg.id"
        class="message-item"
        :class="{ 'recruiter-message': isRecruiterMessage(msg.sender_type) }"
      >
        <div class="message-avatar">
          {{ isRecruiterMessage(msg.sender_type) ? 'HR' : candidate?.name?.charAt(0) || '?' }}
        </div>
        <div class="message-content">
          <div class="message-header">
            <span class="sender-name">{{ isRecruiterMessage(msg.sender_type) ? '招聘方' : candidate?.name || '候选人' }}</span>
            <span class="send-time">{{ formatTime(msg.sent_at) }}</span>
          </div>
          <div class="message-text">{{ msg.content }}</div>
        </div>
      </div>
      <div v-if="messages.length === 0" class="empty-messages">
        暂无消息，开始沟通吧！
      </div>
    </div>

    <div class="message-input-area">
      <input
        v-model="messageInput"
        type="text"
        class="form-input message-input"
        placeholder="输入消息..."
        :disabled="disabled"
        @keyup.enter="handleSend"
      />
      <button class="btn btn-primary" :disabled="disabled" @click="handleSend">发送</button>
    </div>
  </div>
</template>

<style scoped>
.message-panel {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
}

.header-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.conversation-title {
  font-size: 16px;
  font-weight: 600;
}

.conversation-status {
  font-size: 13px;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background-color: #fafafa;
}

.message-item {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
}

.message-item.recruiter-message {
  flex-direction: row-reverse;
}

.message-item.recruiter-message .message-content {
  align-items: flex-end;
}

.message-item.recruiter-message .message-text {
  background-color: #409eff;
  color: #fff;
}

.message-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: #e8f5e9;
  color: #2e7d32;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: 600;
  flex-shrink: 0;
}

.message-content {
  display: flex;
  flex-direction: column;
  gap: 4px;
  max-width: 70%;
}

.message-header {
  display: flex;
  gap: 8px;
  font-size: 12px;
}

.sender-name {
  font-weight: 500;
  color: #666;
}

.send-time {
  color: #999;
}

.message-text {
  padding: 10px 14px;
  background-color: #fff;
  border-radius: 8px;
  font-size: 14px;
  line-height: 1.5;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.05);
}

.empty-messages {
  text-align: center;
  padding: 40px;
  color: #999;
}

.message-input-area {
  display: flex;
  gap: 12px;
  padding: 16px 20px;
  border-top: 1px solid #f0f0f0;
}

.message-input {
  flex: 1;
}

.message-input:disabled {
  background-color: #f5f5f5;
  cursor: not-allowed;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.text-gray {
  color: #999;
}
</style>
