<script setup lang="ts">
defineProps<{
  title: string
  value: number
  color: string
}>()
</script>

<template>
  <div class="stat-card" :style="{ borderLeftColor: color }">
    <div class="stat-value">{{ value }}</div>
    <div class="stat-title">{{ title }}</div>
  </div>
</template>

<style scoped>
.stat-card {
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  border-left: 4px solid;
}

.stat-value {
  font-size: 32px;
  font-weight: 700;
  color: #333;
}

.stat-title {
  font-size: 14px;
  color: #999;
  margin-top: 8px;
}
</style>