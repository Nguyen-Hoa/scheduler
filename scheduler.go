package scheduler

import (
	job "github.com/Nguyen-Hoa/job"
	"github.com/Nguyen-Hoa/worker"
)

type Scheduler interface {
	Schedule(map[string]*worker.ManagerWorker, []job.Job) error
}
