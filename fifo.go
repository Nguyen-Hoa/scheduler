package scheduler

import (
	"errors"

	"github.com/Nguyen-Hoa/worker"
)

type FIFO struct {
}

func (scheduler *FIFO) Schedule(workers map[string]*worker.ManagerWorker, jobs []Job) error {
	if len(jobs) == 0 {
		return nil
	}

	for _, w := range workers {
		if (w.LatestPredictedPower < w.PowerThresh) && w.Available {
			if err := w.StartJob(jobs[0].Image, jobs[0].Cmd, jobs[0].Duration); err != nil {
				return err
			} else {
				_, jobs = jobs[0], jobs[1:]
				return nil
			}
		}
	}

	return errors.New("no workers suitable for scheduling")
}
