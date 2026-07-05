import type { Application } from '@/types'

const STATUS_TEXT_MAP: Record<string, string> = {
  pending: '待筛选',
  communicating: '沟通中',
  interviewing: '面试中',
  offered: '已发offer',
  offer_accepted: '已接受offer',
  rejected: '已拒绝'
}

export function getStatusText(status: string): string {
  return STATUS_TEXT_MAP[status] || status
}

export function getStatusBadgeClass(status: string): string {
  return `badge-${status}`
}

export function canSendOffer(status: string): boolean {
  return status === 'pending' || status === 'communicating' || status === 'interviewing'
}

export function canViewMessages(status: string): boolean {
  return status !== 'pending'
}

export function formatSalary(min: number, max: number): string {
  if (min <= 0 && max <= 0) return '面议'
  return `${Math.round(min / 1000)}K - ${Math.round(max / 1000)}K`
}

export function formatDate(dateStr: string): string {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (Number.isNaN(date.getTime())) return '-'
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

export function formatDateTime(dateStr: string): string {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  if (Number.isNaN(date.getTime())) return '-'
  return `${formatDate(dateStr)} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
}

export function formatTime(dateStr: string): string {
  if (!dateStr) return ''
  const date = new Date(dateStr)
  if (Number.isNaN(date.getTime())) return ''
  return `${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}`
}

export function getStatusHistoryText(history: Application['status_history']): string {
  if (!history || history.length === 0) return '暂无记录'
  return history.map(h => `${getStatusText(h.status)} (${formatDateTime(h.changed_at)})`).join(' → ')
}
