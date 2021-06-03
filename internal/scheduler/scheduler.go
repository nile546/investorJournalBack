package scheduler

type Scheduler struct {
}

type Job struct {
	Name string
	F    func(interface{})
}

func (s *Scheduler) Process(j *Job) {
	j.F()
}
