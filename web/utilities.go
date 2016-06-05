package web

import (
	"fmt"
	"time"
)

func pathGenerator(preacher, keyword string, at time.Time) string {
	return fmt.Sprintf("/recordings/%d-%d-%d - %s - %s.aiff", baseDir, at.Year(), at.Month(), at.Day(), preacher, keyword)
}
