package util

import (
	"fmt"
	"time"
)

func TrackTime(pre time.Time, title string) time.Duration {
	elapsed := time.Since(pre)
	fmt.Printf("%s took: %s\n", title, elapsed)

	return elapsed
}
