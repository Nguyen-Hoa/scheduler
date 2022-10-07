package scheduler

import (
	"github.com/Nguyen-Hoa/worker"
)

type Job struct {
	Image    string   `json:"image"`
	Cmd      []string `json:"cmd"`
	Duration int      `json:"duration"`
}

type scheduler interface {
	Schedule(*map[string]worker.ManagerWorker, *[]Job) error
}
