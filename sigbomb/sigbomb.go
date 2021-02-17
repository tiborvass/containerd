package sigbomb

import (
	"os"
	"syscall"
)

func Start() {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			p.Signal(syscall.SIGURG)
		}
	}()
}
