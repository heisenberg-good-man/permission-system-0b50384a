package store

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"recruitment-platform/internal/model"
)

var (
	jobs         = make(map[string]model.Job)
	candidates   = make(map[string]model.Candidate)
	applications = make(map[string]model.Application)
	messages     = make(map[string]model.Message)
	mu           sync.RWMutex
)

func InitMockData() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()

	jobs["job-1"] = model.Job{
		ID:          "job-1",
		Title:       "高级后端开发工程师",
		Company:     "字节跳动",
		Location:    "北京",
		SalaryMin:   30000,
		SalaryMax:   50000,
		Description: "负责公司核心业务系统的后端开发，与产品、前端团队紧密协作，设计和实现高性能、可扩展的服务架构。",
		Requirements: "本科及以上学历，5年以上后端开发经验，熟悉Go语言，有分布式系统开发经验优先。",
		Status:      "open",
		RecruiterID: "recruiter-1",
		CreatedAt:   now.Add(-24 * time.Hour),
		UpdatedAt:   now.Add(-24 * time.Hour),
	}

	jobs["job-2"] = model.Job{
		ID:          "job-2",
		Title:       "前端开发工程师",
		Company:     "阿里巴巴",
		Location:    "杭州",
		SalaryMin:   25000,
		SalaryMax:   40000,
		Description: "负责电商平台前端开发，优化用户体验，参与技术架构设计。",
		Requirements: "本科及以上学历，3年以上前端开发经验，精通Vue/React，有大型项目经验优先。",
		Status:      "open",
		RecruiterID: "recruiter-1",
		CreatedAt:   now.Add(-48 * time.Hour),
		UpdatedAt:   now.Add(-48 * time.Hour),
	}

	jobs["job-3"] = model.Job{
		ID:          "job-3",
		Title:       "产品经理",
		Company:     "腾讯",
		Location:    "深圳",
		SalaryMin:   20000,
		SalaryMax:   35000,
		Description: "负责社交产品的规划和设计，推动产品迭代优化。",
		Requirements: "本科及以上学历，3年以上产品经验，有社交产品经验优先。",
		Status:      "closed",
		RecruiterID: "recruiter-2",
		CreatedAt:   now.Add(-72 * time.Hour),
		UpdatedAt:   now.Add(-12 * time.Hour),
	}

	candidates["candidate-1"] = model.Candidate{
		ID:            "candidate-1",
		Name:          "张三",
		Email:         "zhangsan@example.com",
		Phone:         "13800138001",
		ResumeSummary: "10年后端开发经验，曾任职于多家知名互联网公司，精通Go/Python，有丰富的分布式系统设计经验。",
		Experience:    10,
		Education:     "本科",
		Skills:        []string{"Go", "Python", "Redis", "MySQL"},
		CreatedAt:     now.Add(-24 * time.Hour),
	}

	candidates["candidate-2"] = model.Candidate{
		ID:            "candidate-2",
		Name:          "李四",
		Email:         "lisi@example.com",
		Phone:         "13800138002",
		ResumeSummary: "5年前端开发经验，精通Vue、React，有大型电商项目经验。",
		Experience:    5,
		Education:     "本科",
		Skills:        []string{"Vue", "React", "TypeScript", "Webpack"},
		CreatedAt:     now.Add(-48 * time.Hour),
	}

	applications["app-1"] = model.Application{
		ID:          "app-1",
		JobID:       "job-1",
		CandidateID: "candidate-1",
		Status:      "pending",
		AppliedAt:   now.Add(-12 * time.Hour),
		UpdatedAt:   now.Add(-12 * time.Hour),
	}

	applications["app-2"] = model.Application{
		ID:          "app-2",
		JobID:       "job-2",
		CandidateID: "candidate-2",
		Status:      "communicating",
		AppliedAt:   now.Add(-6 * time.Hour),
		UpdatedAt:   now.Add(-2 * time.Hour),
	}

	messages["msg-1"] = model.Message{
		ID:            "msg-1",
		ApplicationID: "app-2",
		SenderID:      "recruiter-1",
		SenderType:    "recruiter",
		Content:       "您好，感谢您投递前端开发工程师职位，请问您方便下周安排面试吗？",
		SentAt:        now.Add(-4 * time.Hour),
	}

	messages["msg-2"] = model.Message{
		ID:            "msg-2",
		ApplicationID: "app-2",
		SenderID:      "candidate-2",
		SenderType:    "candidate",
		Content:       "您好，我下周都方便，请问具体时间安排是？",
		SentAt:        now.Add(-2 * time.Hour),
	}
}

func CreateJob(job model.Job) string {
	mu.Lock()
	defer mu.Unlock()
	job.ID = uuid.New().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
	jobs[job.ID] = job
	return job.ID
}

func GetJob(id string) (model.Job, bool) {
	mu.RLock()
	defer mu.RUnlock()
	job, ok := jobs[id]
	return job, ok
}

func ListJobs(status string) []model.Job {
	mu.RLock()
	defer mu.RUnlock()
	var result []model.Job
	for _, job := range jobs {
		if status == "" || job.Status == status {
			result = append(result, job)
		}
	}
	return result
}

func UpdateJob(id string, job model.Job) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := jobs[id]; !ok {
		return false
	}
	job.ID = id
	job.UpdatedAt = time.Now()
	jobs[id] = job
	return true
}

func DeleteJob(id string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := jobs[id]; !ok {
		return false
	}
	delete(jobs, id)
	return true
}

func CreateCandidate(candidate model.Candidate) string {
	mu.Lock()
	defer mu.Unlock()
	candidate.ID = uuid.New().String()
	candidate.CreatedAt = time.Now()
	candidates[candidate.ID] = candidate
	return candidate.ID
}

func GetCandidate(id string) (model.Candidate, bool) {
	mu.RLock()
	defer mu.RUnlock()
	candidate, ok := candidates[id]
	return candidate, ok
}

func ListCandidates() []model.Candidate {
	mu.RLock()
	defer mu.RUnlock()
	var result []model.Candidate
	for _, candidate := range candidates {
		result = append(result, candidate)
	}
	return result
}

func UpdateCandidate(id string, candidate model.Candidate) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := candidates[id]; !ok {
		return false
	}
	candidate.ID = id
	candidates[id] = candidate
	return true
}

func CreateApplication(app model.Application) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, existing := range applications {
		if existing.JobID == app.JobID && existing.CandidateID == app.CandidateID {
			return "", ErrDuplicateApplication
		}
	}

	job, ok := jobs[app.JobID]
	if !ok {
		return "", ErrJobNotFound
	}
	if job.Status != "open" {
		return "", ErrJobClosed
	}

	app.ID = uuid.New().String()
	app.Status = "pending"
	app.AppliedAt = time.Now()
	app.UpdatedAt = time.Now()
	applications[app.ID] = app
	return app.ID, nil
}

func GetApplication(id string) (model.Application, bool) {
	mu.RLock()
	defer mu.RUnlock()
	app, ok := applications[id]
	return app, ok
}

func ListApplications(jobID, status string) []model.Application {
	mu.RLock()
	defer mu.RUnlock()
	var result []model.Application
	for _, app := range applications {
		if jobID != "" && app.JobID != jobID {
			continue
		}
		if status != "" && app.Status != status {
			continue
		}
		result = append(result, app)
	}
	return result
}

func UpdateApplicationStatus(id string, status string) bool {
	mu.Lock()
	defer mu.Unlock()
	app, ok := applications[id]
	if !ok {
		return false
	}
	app.Status = status
	app.UpdatedAt = time.Now()
	applications[id] = app
	return true
}

func CreateMessage(msg model.Message) string {
	mu.Lock()
	defer mu.Unlock()
	msg.ID = uuid.New().String()
	msg.SentAt = time.Now()
	messages[msg.ID] = msg
	return msg.ID
}

func ListMessages(applicationID string) []model.Message {
	mu.RLock()
	defer mu.RUnlock()
	var result []model.Message
	for _, msg := range messages {
		if msg.ApplicationID == applicationID {
			result = append(result, msg)
		}
	}
	return result
}

func GetStats() model.Stats {
	mu.RLock()
	defer mu.RUnlock()

	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	stats := model.Stats{}
	for _, job := range jobs {
		if job.Status == "open" {
			stats.OpenJobs++
		}
	}
	for _, app := range applications {
		if app.AppliedAt.After(todayStart) {
			stats.TodayApplications++
		}
		if app.Status == "pending" {
			stats.PendingReview++
		}
		if app.Status == "communicating" {
			stats.InCommunication++
		}
	}
	return stats
}

var (
	ErrDuplicateApplication = &StoreError{"已投递该职位"}
	ErrJobNotFound          = &StoreError{"职位不存在"}
	ErrJobClosed            = &StoreError{"职位已关闭"}
)

type StoreError struct {
	Message string
}

func (e *StoreError) Error() string {
	return e.Message
}