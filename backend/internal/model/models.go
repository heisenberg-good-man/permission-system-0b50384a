package model

import "time"

type Job struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	SalaryMin   int       `json:"salary_min"`
	SalaryMax   int       `json:"salary_max"`
	Description string    `json:"description"`
	Requirements string   `json:"requirements"`
	Status      string    `json:"status"`
	RecruiterID string    `json:"recruiter_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Candidate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	ResumeSummary string  `json:"resume_summary"`
	Experience  int       `json:"experience"`
	Education   string    `json:"education"`
	Skills      []string  `json:"skills"`
	CreatedAt   time.Time `json:"created_at"`
}

type Application struct {
	ID          string    `json:"id"`
	JobID       string    `json:"job_id"`
	CandidateID string    `json:"candidate_id"`
	Status      string    `json:"status"`
	AppliedAt   time.Time `json:"applied_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Message struct {
	ID          string    `json:"id"`
	ApplicationID string  `json:"application_id"`
	SenderID    string    `json:"sender_id"`
	SenderType  string    `json:"sender_type"`
	Content     string    `json:"content"`
	SentAt      time.Time `json:"sent_at"`
}

type Stats struct {
	OpenJobs         int `json:"open_jobs"`
	TodayApplications int `json:"today_applications"`
	PendingReview    int `json:"pending_review"`
	InCommunication  int `json:"in_communication"`
}