package scheduler

import (
	"errors"

	job "github.com/Nguyen-Hoa/job"
	"github.com/Nguyen-Hoa/worker"
)

type FIFO struct {
}

func (scheduler *FIFO) Schedule(workers map[string]*worker.ManagerWorker, jobs job.SharedJobsArray) error {
	if jobs.Length() == 0 {
		return nil
	}

	for _, w := range workers {
		if (w.LatestPredictedPower < w.PowerThresh) && w.Available {
			j := jobs.Pop()
			if err := w.StartJob(j.Image, j.Cmd, j.Duration); err != nil {
				jobs.Append(j)
				return err
			} else {
				return nil
			}
		}
	}

	return errors.New("no workers suitable for scheduling")
}
