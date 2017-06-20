// +build darwin freebsd solaris

package content

import (
	"os"
	"time"
)

func getStartTime(fi os.FileInfo) time.Time {
	return time.Time{}
}
