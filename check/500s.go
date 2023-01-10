package check

import (
	"errors"
	"fmt"
)

func Num500sIn5Minutes(above int) func() error {
	// assume this is worked out somewhere - could write back to something inside
	// the container on a failed request.
	number := 200

	return func() error {
		fmt.Println("I'm being fired")
		if number > above {
			fmt.Println("Im failing")
			return errors.New("number of 500s in 5 minutes is too high")
		}

		return nil
	}
}
