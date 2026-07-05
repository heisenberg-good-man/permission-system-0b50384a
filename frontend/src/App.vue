<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import type { Job, Candidate, Application, Message, Stats, Role } from '@/types'
import {
  getJobs,
  createJob,
  updateJob,
  deleteJob,
  getCandidates,
  createCandidate,
  updateCandidate,
  getApplications,
  createApplication,
  updateApplicationStatus,
  getMessages,
  createMessage,
  getStats
} from '@/api'

import RoleSwitcher from '@/components/RoleSwitcher.vue'
import StatCard from '@/components/StatCard.vue'
import JobList from '@/components/JobList.vue'
import JobForm from '@/components/JobForm.vue'
import JobDetail from '@/components/JobDetail.vue'
import ApplicationList from '@/components/ApplicationList.vue'
import CandidateDetail from '@/components/CandidateDetail.vue'
import MessagePanel from '@/components/MessagePanel.vue'

const currentRole = ref<Role>('candidate')
const stats = ref<Stats>({
  open_jobs: 0,
  today_applications: 0,
  pending_review: 0,
  in_communication: 0
})

const jobs = ref<Job[]>([])
const candidates = ref<Candidate[]>([])
const applications = ref<Application[]>([])

const candidatesMap = computed(() => {
  const map = new Map<string, Candidate>()
  candidates.value.forEach(c => map.set(c.id, c))
  return map
})

const jobsMap = computed(() => {
  const map = new Map<string, Job>()
  jobs.value.forEach(j => map.set(j.id, j))
  return map
})

const currentCandidate = ref<Candidate>({
  id: '',
  name: '',
  email: '',
  phone: '',
  resume_summary: '',
  experience: 0,
  education: '',
  skills: [],
  created_at: ''
})

const skillsInput = computed({
  get: () => currentCandidate.value.skills.join(', '),
  set: (val: string) => {
    currentCandidate.value.skills = val.split(',').map(s => s.trim()).filter(Boolean)
  }
})

const showJobDetail = ref(false)
const selectedJob = ref<Job | null>(null)

const showJobForm = ref(false)
const editingJob = ref<Job | null>(null)

const showCandidateDetail = ref(false)
const selectedCandidate = ref<Candidate | null>(null)

const showMessagePanel = ref(false)
const currentApplication = ref<Application | null>(null)
const messages = ref<Message[]>([])

const searchKeyword = ref('')
const filterLocation = ref('')

const recruiterFilter = ref({
  jobId: '',
  status: '',
  keyword: ''
})

const toast = ref<{ show: boolean; type: 'success' | 'error'; message: string }>({
  show: false,
  type: 'success',
  message: ''
})

const loading = ref(false)

function showToast(type: 'success' | 'error', message: string) {
  toast.value = { show: true, type, message }
  setTimeout(() => {
    toast.value.show = false
  }, 3000)
}

async function loadData() {
  loading.value = true
  try {
    const [jobsData, candidatesData, applicationsData, statsData] = await Promise.all([
      getJobs(),
      getCandidates(),
      getApplications(),
      getStats()
    ])
    jobs.value = jobsData
    candidates.value = candidatesData
    applications.value = applicationsData
    stats.value = statsData
    if (candidates.value.length > 0) {
      currentCandidate.value = candidates.value[0]
    }
  } catch (error) {
    showToast('error', '加载数据失败')
  } finally {
    loading.value = false
  }
}

async function loadMessages(applicationId: string) {
  try {
    messages.value = await getMessages(applicationId)
  } catch (error) {
    showToast('error', '加载消息失败')
  }
}

const filteredJobs = computed(() => {
  let result = jobs.value
  if (currentRole.value === 'candidate') {
    result = result.filter(j => j.status === 'open')
  }
  if (searchKeyword.value) {
    const keyword = searchKeyword.value.toLowerCase()
    result = result.filter(j =>
      j.title.toLowerCase().includes(keyword) ||
      j.company.toLowerCase().includes(keyword)
    )
  }
  if (filterLocation.value) {
    result = result.filter(j => j.location === filterLocation.value)
  }
  return result
})

const myApplications = computed(() => {
  if (!currentCandidate.value.id) return []
  return applications.value.filter(a => a.candidate_id === currentCandidate.value.id)
})

const filteredApplications = computed(() => {
  let result = applications.value
  if (recruiterFilter.value.jobId) {
    result = result.filter(a => a.job_id === recruiterFilter.value.jobId)
  }
  if (recruiterFilter.value.status) {
    result = result.filter(a => a.status === recruiterFilter.value.status)
  }
  if (recruiterFilter.value.keyword) {
    const keyword = recruiterFilter.value.keyword.toLowerCase()
    result = result.filter(a => {
      const candidate = candidatesMap.value.get(a.candidate_id)
      const job = jobsMap.value.get(a.job_id)
      return (candidate?.name.toLowerCase().includes(keyword) ||
        candidate?.email.toLowerCase().includes(keyword) ||
        job?.title.toLowerCase().includes(keyword))
    })
  }
  return result
})

function switchRole(role: Role) {
  currentRole.value = role
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

function formatDate(dateStr: string): string {
  const date = new Date(dateStr)
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')}`
}

function viewJobDetail(job: Job) {
  selectedJob.value = job
  showJobDetail.value = true
}

async function handleApply(jobId: string, resumeSummary: string) {
  if (!currentCandidate.value.id) {
    showToast('error', '请先完善简历信息')
    return
  }
  try {
    if (resumeSummary.trim()) {
      currentCandidate.value.resume_summary = resumeSummary
      await updateCandidate(currentCandidate.value.id, { resume_summary: resumeSummary })
    }
    await createApplication({ job_id: jobId, candidate_id: currentCandidate.value.id })
    showToast('success', '投递成功')
    showJobDetail.value = false
    loadData()
  } catch (error: any) {
    const errorMsg = error.response?.data?.error || '投递失败'
    showToast('error', errorMsg)
  }
}

function openJobForm(job?: Job | null) {
  editingJob.value = job || null
  showJobForm.value = true
}

function handleJobSubmit(jobData: Omit<Job, 'id' | 'created_at' | 'updated_at'> | Partial<Job>) {
  if (editingJob.value) {
    updateJob(editingJob.value.id, jobData as Partial<Job>)
      .then(() => {
        showToast('success', '职位更新成功')
        showJobForm.value = false
        loadData()
      })
      .catch(() => showToast('error', '更新失败'))
  } else {
    createJob(jobData as Omit<Job, 'id' | 'created_at' | 'updated_at'>)
      .then(() => {
        showToast('success', '职位创建成功')
        showJobForm.value = false
        loadData()
      })
      .catch(() => showToast('error', '创建失败'))
  }
}

function handleDeleteJob(jobId: string) {
  if (confirm('确定要删除这个职位吗？')) {
    deleteJob(jobId)
      .then(() => {
        showToast('success', '删除成功')
        loadData()
      })
      .catch(() => showToast('error', '删除失败'))
  }
}

function viewCandidate(candidateId: string) {
  const candidate = candidates.value.find(c => c.id === candidateId)
  if (candidate) {
    selectedCandidate.value = candidate
    showCandidateDetail.value = true
  }
}

function viewMessages(app: Application) {
  currentApplication.value = app
  showMessagePanel.value = true
  loadMessages(app.id)
}

function handleUpdateStatus(applicationId: string, status: string) {
  updateApplicationStatus(applicationId, status)
    .then(() => {
      showToast('success', '状态更新成功')
      loadData()
      if (currentApplication.value?.id === applicationId && status !== 'communicating') {
        showMessagePanel.value = false
      }
    })
    .catch(() => showToast('error', '更新失败'))
}

function handleSendMessage(content: string) {
  if (!currentApplication.value) return
  const senderType = currentRole.value
  const senderId = senderType === 'recruiter' ? 'recruiter-1' : currentCandidate.value.id
  createMessage({
    application_id: currentApplication.value.id,
    sender_id: senderId,
    sender_type: senderType,
    content
  })
    .then(() => {
      loadMessages(currentApplication.value!.id)
    })
    .catch(() => showToast('error', '发送失败'))
}

function saveCandidateInfo() {
  if (!currentCandidate.value.name || !currentCandidate.value.email) {
    showToast('error', '姓名和邮箱为必填项')
    return
  }
  if (currentCandidate.value.id) {
    updateCandidate(currentCandidate.value.id, currentCandidate.value)
      .then(() => showToast('success', '信息更新成功'))
      .catch(() => showToast('error', '更新失败'))
  } else {
    createCandidate(currentCandidate.value)
      .then(({ id }) => {
        currentCandidate.value.id = id
        showToast('success', '信息保存成功')
        loadData()
      })
      .catch(() => showToast('error', '保存失败'))
  }
}

const locations = computed(() => {
  const set = new Set(jobs.value.map(j => j.location))
  return Array.from(set)
})

onMounted(() => {
  loadData()
})
</script>

<template>
  <div class="app-container">
    <header class="app-header">
      <div class="header-left">
        <h1>招聘工作台</h1>
      </div>
      <div class="header-right">
        <RoleSwitcher :current-role="currentRole" @switch="switchRole" />
      </div>
    </header>

    <main class="app-main">
      <div class="stats-section">
        <div class="stats-grid">
          <StatCard title="开放职位数" :value="stats.open_jobs" color="#409eff" />
          <StatCard title="今日投递数" :value="stats.today_applications" color="#67c23a" />
          <StatCard title="待筛选人数" :value="stats.pending_review" color="#e6a23c" />
          <StatCard title="沟通中人数" :value="stats.in_communication" color="#909399" />
        </div>
      </div>

      <div v-if="currentRole === 'candidate'" class="candidate-section">
        <div class="card">
          <h2>我的简历</h2>
          <div class="flex gap-4">
            <div class="form-group" style="flex: 1">
              <label class="form-label">姓名 *</label>
              <input v-model="currentCandidate.name" type="text" class="form-input" placeholder="请输入姓名" />
            </div>
            <div class="form-group" style="flex: 1">
              <label class="form-label">邮箱 *</label>
              <input v-model="currentCandidate.email" type="email" class="form-input" placeholder="请输入邮箱" />
            </div>
            <div class="form-group" style="flex: 1">
              <label class="form-label">电话</label>
              <input v-model="currentCandidate.phone" type="tel" class="form-input" placeholder="请输入电话" />
            </div>
          </div>
          <div class="flex gap-4">
            <div class="form-group" style="flex: 1">
              <label class="form-label">学历</label>
              <select v-model="currentCandidate.education" class="form-select">
                <option value="">请选择学历</option>
                <option value="本科">本科</option>
                <option value="硕士">硕士</option>
                <option value="博士">博士</option>
                <option value="大专">大专</option>
                <option value="高中">高中</option>
              </select>
            </div>
            <div class="form-group" style="flex: 1">
              <label class="form-label">工作经验（年）</label>
              <input v-model.number="currentCandidate.experience" type="number" class="form-input" placeholder="0" />
            </div>
          </div>
          <div class="form-group">
            <label class="form-label">技能标签（逗号分隔）</label>
            <input v-model="skillsInput" type="text" class="form-input" placeholder="例如：Go, Python, Vue" />
          </div>
          <div class="form-group">
            <label class="form-label">简历摘要</label>
            <textarea v-model="currentCandidate.resume_summary" class="form-textarea" placeholder="请简要介绍您的工作经历和技能特长"></textarea>
          </div>
          <button class="btn btn-primary" @click="saveCandidateInfo">保存简历</button>
        </div>

        <div class="card">
          <div class="card-header">
            <h2>职位列表</h2>
            <div class="search-bar">
              <input v-model="searchKeyword" type="text" class="form-input search-input" placeholder="搜索职位或公司" />
              <select v-model="filterLocation" class="form-select location-select">
                <option value="">全部地点</option>
                <option v-for="loc in locations" :key="loc" :value="loc">{{ loc }}</option>
              </select>
            </div>
          </div>
          <JobList :jobs="filteredJobs" @view-detail="viewJobDetail" />
        </div>

        <div class="card">
          <div class="card-header">
            <h2>我的投递</h2>
          </div>
          <div class="my-applications">
            <div v-for="app in myApplications" :key="app.id" class="application-item">
              <div class="app-info">
                <div class="app-job-title">{{ jobsMap.get(app.job_id)?.title || '-' }}</div>
                <div class="app-company">{{ jobsMap.get(app.job_id)?.company || '-' }}</div>
              </div>
              <div class="app-status">
                <span class="badge" :class="`badge-${app.status}`">
                  {{ getStatusText(app.status) }}
                </span>
              </div>
              <div class="app-meta">
                <span>{{ formatDate(app.applied_at) }}</span>
              </div>
              <div class="app-actions">
                <button 
                  v-if="app.status === 'communicating'" 
                  class="btn btn-primary btn-sm" 
                  @click="viewMessages(app)"
                >
                  消息
                </button>
              </div>
            </div>
            <div v-if="myApplications.length === 0" class="empty-state">
              暂无投递记录，快去投递心仪的职位吧！
            </div>
          </div>
        </div>
      </div>

      <div v-else class="recruiter-section">
        <div class="card">
          <div class="card-header">
            <h2>职位管理</h2>
            <button class="btn btn-primary" @click="openJobForm()">发布新职位</button>
          </div>
          <JobList 
          :jobs="filteredJobs" 
          :show-status="true" 
          :show-actions="true"
          @view-detail="viewJobDetail"
          @edit="openJobForm"
          @delete="handleDeleteJob"
        />
        </div>

        <div class="card">
          <div class="card-header">
            <h2>投递列表</h2>
          </div>
          <div class="filter-bar">
            <input 
              v-model="recruiterFilter.keyword" 
              type="text" 
              class="form-input" 
              placeholder="搜索候选人姓名、邮箱或职位"
              style="width: 200px;"
            />
            <select 
              v-model="recruiterFilter.jobId" 
              class="form-select"
              style="width: 150px;"
            >
              <option value="">全部职位</option>
              <option v-for="job in jobs" :key="job.id" :value="job.id">{{ job.title }}</option>
            </select>
            <select 
              v-model="recruiterFilter.status" 
              class="form-select"
              style="width: 120px;"
            >
              <option value="">全部状态</option>
              <option value="pending">待筛选</option>
              <option value="communicating">沟通中</option>
              <option value="interviewing">面试中</option>
              <option value="offered">已发offer</option>
              <option value="rejected">已拒绝</option>
            </select>
          </div>
          <ApplicationList
            :applications="filteredApplications"
            :candidates="candidatesMap"
            :jobs="jobsMap"
            @view-candidate="viewCandidate"
            @view-message="viewMessages"
            @update-status="handleUpdateStatus"
          />
        </div>
      </div>
    </main>

    <div v-if="showJobDetail" class="modal-overlay" @click.self="showJobDetail = false">
      <div class="modal">
        <JobDetail
          :job="selectedJob"
          :current-candidate="currentCandidate"
          :is-recruiter="currentRole === 'recruiter'"
          @close="showJobDetail = false"
          @apply="handleApply"
          @edit="(job) => { showJobDetail = false; openJobForm(job) }"
          @delete="(jobId) => { showJobDetail = false; handleDeleteJob(jobId) }"
        />
      </div>
    </div>

    <div v-if="showJobForm" class="modal-overlay" @click.self="showJobForm = false">
      <div class="modal">
        <div class="modal-header">
          <div class="modal-title">{{ editingJob ? '编辑职位' : '发布职位' }}</div>
          <button class="modal-close" @click="showJobForm = false">&times;</button>
        </div>
        <div class="modal-body">
          <JobForm :job="editingJob" @submit="handleJobSubmit" @cancel="showJobForm = false" />
        </div>
      </div>
    </div>

    <div v-if="showCandidateDetail" class="modal-overlay" @click.self="showCandidateDetail = false">
      <div class="modal" style="max-width: 900px;">
        <CandidateDetail 
          :candidate="selectedCandidate" 
          :applications="applications"
          :jobs="jobsMap"
          @close="showCandidateDetail = false"
          @view-message="viewMessages"
          @update-status="handleUpdateStatus"
        />
      </div>
    </div>

    <div v-if="showMessagePanel" class="modal-overlay" @click.self="showMessagePanel = false">
      <div class="modal" style="max-width: 800px; max-height: 70vh;">
        <MessagePanel
          :messages="messages"
          :application="currentApplication"
          :job="currentApplication ? jobsMap.get(currentApplication.job_id) || null : null"
          :candidate="currentApplication ? candidatesMap.get(currentApplication.candidate_id) || null : null"
          @send="handleSendMessage"
          @close="showMessagePanel = false"
        />
      </div>
    </div>

    <div v-if="toast.show" class="toast" :class="`toast-${toast.type}`">
      {{ toast.message }}
    </div>

    <div v-if="loading" class="loading">
      <div class="loading-spinner"></div>
    </div>
  </div>
</template>

<style scoped>
.app-container {
  min-height: 100vh;
  background-color: #f5f7fa;
}

.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  background-color: #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.header-left h1 {
  font-size: 20px;
  font-weight: 600;
  color: #333;
}

.header-right {
  display: flex;
  align-items: center;
}

.app-main {
  padding: 24px;
}

.stats-section {
  margin-bottom: 24px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

@media (max-width: 1024px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

.card {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
  padding: 20px;
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.card-header h2 {
  font-size: 16px;
  font-weight: 600;
}

.search-bar {
  display: flex;
  gap: 12px;
}

.search-input {
  width: 200px;
}

.location-select {
  width: 120px;
}

.candidate-section,
.recruiter-section {
  max-width: 100%;
}

.my-applications {
  margin-top: 16px;
}

.application-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
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

.app-company {
  font-size: 12px;
  color: #999;
  margin-top: 4px;
}

.app-status {
  margin-right: 16px;
}

.app-meta {
  font-size: 12px;
  color: #999;
  margin-right: 16px;
}

.app-actions {
  display: flex;
  gap: 8px;
}

.filter-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid #f0f0f0;
}
</style>