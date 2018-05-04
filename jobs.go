package dkron

import (
	"net/http"
	"time"
)

const jobBasePath = "/v1/jobs"

type Job struct {
	Name                 string                            `json:"name"`
	Schedule             string                            `json:"schedule"`
	Shell                bool                              `json:"shell"`
	Command              string                            `json:"command"`
	EnvironmentVariables []string                          `json:"environment_variables"`
	Payload              []byte                            `json:"payload"`
	Owner                string                            `json:"owner"`
	OwnerEmail           string                            `json:"owner_email"`
	SuccessCount         int                               `json:"success_count"`
	ErrorCount           int                               `json:"error_count"`
	LastSuccess          time.Time                         `json:"last_success"`
	LastError            time.Time                         `json:"last_error"`
	Disabled             bool                              `json:"disabled"`
	Tags                 map[string]string                 `json:"tags"`
	Retries              uint                              `json:"retries"`
	DependentJobs        []string                          `json:"dependent_jobs"`
	ParentJob            string                            `json:"parent_job"`
	Processors           map[string]map[string]interface{} `json:"processors"`
	Concurrency          string                            `json:"concurrency"`
}

type JobsService interface {
	Add(*Job) (*Job, error)
}

type JobsServiceOp struct {
	client *Client
}

func (s *JobsServiceOp) Add(j *Job) (*Job, error) {
	req, err := s.client.NewRequest(http.MethodPost, jobBasePath, j)
	if err != nil {
		return nil, err
	}

	job := &Job{}
	_, err = s.client.Do(req, job)

	return job, nil
}
