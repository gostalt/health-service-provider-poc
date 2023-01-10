package health

import "fmt"

func route() func() string {
	checks := []string{"1", "2"}
	return func() string {
		return fmt.Sprintf("these are the health routes, %v", checks)
	}
}
