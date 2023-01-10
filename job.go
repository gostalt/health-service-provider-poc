package health

import "fmt"

type job struct {
	checks map[string]func() error
}

func (j job) ShouldFire() bool {
	return true
}

func (j job) Handle() error {
	var failed []string

	for name, check := range j.checks {
		if err := check(); err != nil {
			failed = append(failed, name)
		}
	}

	if len(failed) > 0 {
		return fmt.Errorf("Some checks failed: %v", failed)
	}
	return nil
}
