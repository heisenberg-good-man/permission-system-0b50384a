import axios from 'axios'
import type { Job, Candidate, Application, Message, Stats } from '@/types'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000
})

export async function getJobs(status?: string): Promise<Job[]> {
  const params = status ? { status } : {}
  const response = await api.get('/jobs', { params })
  return response.data.data
}

export async function getJob(id: string): Promise<Job> {
  const response = await api.get(`/jobs/${id}`)
  return response.data.data
}

export async function createJob(job: Omit<Job, 'id' | 'created_at' | 'updated_at'>): Promise<{ id: string }> {
  const response = await api.post('/jobs', job)
  return response.data
}

export async function updateJob(id: string, job: Partial<Job>): Promise<void> {
  await api.put(`/jobs/${id}`, job)
}

export async function deleteJob(id: string): Promise<void> {
  await api.delete(`/jobs/${id}`)
}

export async function getCandidates(): Promise<Candidate[]> {
  const response = await api.get('/candidates')
  return response.data.data
}

export async function getCandidate(id: string): Promise<Candidate> {
  const response = await api.get(`/candidates/${id}`)
  return response.data.data
}

export async function createCandidate(candidate: Omit<Candidate, 'id' | 'created_at'>): Promise<{ id: string }> {
  const response = await api.post('/candidates', candidate)
  return response.data
}

export async function updateCandidate(id: string, candidate: Partial<Candidate>): Promise<void> {
  await api.put(`/candidates/${id}`, candidate)
}

export async function getApplications(jobId?: string, candidateId?: string, status?: string): Promise<Application[]> {
  const params: Record<string, string> = {}
  if (jobId) params.job_id = jobId
  if (candidateId) params.candidate_id = candidateId
  if (status) params.status = status
  const response = await api.get('/applications', { params })
  return response.data.data
}

export async function getApplication(id: string): Promise<Application> {
  const response = await api.get(`/applications/${id}`)
  return response.data.data
}

export async function createApplication(application: { job_id: string; candidate_id: string }): Promise<{ id: string }> {
  const response = await api.post('/applications', application)
  return response.data
}

export async function updateApplicationStatus(id: string, status: string): Promise<void> {
  await api.put(`/applications/${id}/status`, { status })
}

export async function getMessages(applicationId: string): Promise<Message[]> {
  const response = await api.get('/messages', { params: { application_id: applicationId } })
  return response.data.data
}

export async function createMessage(message: {
  application_id: string
  sender_id: string
  sender_type: string
  content: string
}): Promise<{ id: string }> {
  const response = await api.post('/messages', message)
  return response.data
}

export async function getStats(): Promise<Stats> {
  const response = await api.get('/stats')
  return response.data.data
}