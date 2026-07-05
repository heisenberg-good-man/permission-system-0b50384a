export interface Job {
  id: string
  title: string
  company: string
  location: string
  salary_min: number
  salary_max: number
  description: string
  requirements: string
  status: string
  recruiter_id: string
  created_at: string
  updated_at: string
}

export interface Candidate {
  id: string
  name: string
  email: string
  phone: string
  resume_summary: string
  experience: number
  education: string
  skills: string[]
  created_at: string
}

export interface StatusHistory {
  status: string
  changed_at: string
}

export interface OfferDetail {
  offer_salary_min: number
  offer_salary_max: number
  start_date: string
  remarks: string
  sent_at?: string
}

export interface Application {
  id: string
  job_id: string
  candidate_id: string
  status: string
  applied_at: string
  updated_at: string
  status_history?: StatusHistory[]
  offer_status?: string
  offer_detail?: OfferDetail
  last_activity?: string
}

export interface Message {
  id: string
  application_id: string
  sender_id: string
  sender_type: string
  content: string
  sent_at: string
}

export interface Stats {
  open_jobs: number
  today_applications: number
  pending_review: number
  in_communication: number
  interviewing: number
  offered: number
  accepted: number
  rejected: number
}

export type Role = 'candidate' | 'recruiter'