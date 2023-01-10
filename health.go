package health

import (
	"errors"

	"github.com/gostalt/framework/schedule"
	"github.com/gostalt/framework/service"
	"github.com/gostalt/router"
	"github.com/sarulabs/di/v2"
)

func New(checks map[string]func() error) *healthServiceProvider {
	return &healthServiceProvider{
		checks: checks,
	}
}

type healthServiceProvider struct {
	service.BaseProvider
	checks map[string]func() error
}

func (p healthServiceProvider) Register(b *di.Builder) error {
	b.Add(di.Def{
		Name: "health-job",
		Build: func(ctn di.Container) (interface{}, error) {
			return job{checks: p.checks}, nil
		},
	})
	return nil
}

func (p healthServiceProvider) Boot(c di.Container) error {
	if err := p.addToRouter(c); err != nil {
		return err
	}

	if err := p.addToScheduler(c); err != nil {
		return err
	}

	return nil
}

func (p healthServiceProvider) addToRouter(c di.Container) error {
	val, err := c.SafeGet("router")
	if err != nil {
		return err
	}

	rtr, ok := val.(*router.Router)
	if !ok {
		return errors.New("can't get router")
	}

	rtr.Get("health", route())

	return nil
}

func (p healthServiceProvider) addToScheduler(c di.Container) error {
	val, err := c.SafeGet("scheduler")
	if err != nil {
		return err
	}

	scheduler, ok := val.(*schedule.Runner)
	if !ok {
		return errors.New("unable to get scheduler")
	}

	job, err := p.getJob(c)
	if err != nil {
		return err
	}

	scheduler.Add(job)

	return nil
}

func (p healthServiceProvider) getJob(c di.Container) (schedule.Job, error) {
	val, err := c.SafeGet("health-job")
	if err != nil {
		return nil, err
	}

	job, ok := val.(schedule.Job)
	if !ok {
		return nil, errors.New("unable to get job")
	}

	return job, nil
}
