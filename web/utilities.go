package web

import (
	"time"
	"fmt"
)

func pathGenerator(preacher, keyword string, at time.Time) string {
	return fmt.Sprintf("%s/%d-%d-%d - %s - %s.aiff", baseDir,  at.Year(), at.Month(), at.Day(), preacher, keyword)
}
