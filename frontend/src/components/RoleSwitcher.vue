<script setup lang="ts">
import type { Role } from '@/types'

defineProps<{
  currentRole: Role
}>()

const emit = defineEmits<{
  switch: [role: Role]
}>()
</script>

<template>
  <div class="role-switcher">
    <span class="switch-label">当前角色：</span>
    <div class="switch-buttons">
      <button
        class="switch-btn"
        :class="{ active: currentRole === 'candidate' }"
        @click="emit('switch', 'candidate')"
      >
        应聘方
      </button>
      <button
        class="switch-btn"
        :class="{ active: currentRole === 'recruiter' }"
        @click="emit('switch', 'recruiter')"
      >
        招聘方
      </button>
    </div>
  </div>
</template>

<style scoped>
.role-switcher {
  display: flex;
  align-items: center;
  gap: 12px;
}

.switch-label {
  font-size: 14px;
  color: #666;
}

.switch-buttons {
  display: flex;
  background-color: #f0f0f0;
  border-radius: 4px;
  padding: 2px;
}

.switch-btn {
  padding: 6px 16px;
  border: none;
  border-radius: 3px;
  font-size: 13px;
  cursor: pointer;
  transition: all 0.2s;
  background-color: transparent;
  color: #666;
}

.switch-btn.active {
  background-color: #409eff;
  color: #fff;
}
</style>