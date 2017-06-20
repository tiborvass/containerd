package console

// TODO: change everything to work on solaris

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

const TODO uintptr = 0

func tcget(fd uintptr, p *unix.Termios) error {
	return ioctl(fd, TODO, uintptr(unsafe.Pointer(p)))
}

func tcset(fd uintptr, p *unix.Termios) error {
	return ioctl(fd, TODO, uintptr(unsafe.Pointer(p)))
}

func ioctl(fd, flag, data uintptr) error {
	/*
	if _, _, err := unix.Syscall(TODO, fd, flag, data); err != 0 {
		return err
	}
	*/
	return nil
}

// unlockpt unlocks the slave pseudoterminal device corresponding to the master pseudoterminal referred to by f.
// unlockpt should be called before opening the slave side of a pty.
// This does not exist on FreeBSD, it does not allocate controlling terminals on open
func unlockpt(f *os.File) error {
	return nil
}

// ptsname retrieves the name of the first available pts for the given master.
func ptsname(f *os.File) (string, error) {
	var n int32
	if err := ioctl(f.Fd(), TODO, uintptr(unsafe.Pointer(&n))); err != nil {
		return "", err
	}
	return fmt.Sprintf("/dev/pts/%d", n), nil
}

func saneTerminal(f *os.File) error {
	// Go doesn't have a wrapper for any of the termios ioctls.
	var termios unix.Termios
	if err := tcget(f.Fd(), &termios); err != nil {
		return err
	}
	// Set -onlcr so we don't have to deal with \r.
	termios.Oflag &^= unix.ONLCR
	return tcset(f.Fd(), &termios)
}

func cfmakeraw(t unix.Termios) unix.Termios {
	t.Iflag = t.Iflag & ^uint32((unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON))
	t.Oflag = t.Oflag & ^uint32(unix.OPOST)
	t.Lflag = t.Lflag & ^(uint32(unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN))
	t.Cflag = t.Cflag & ^(uint32(unix.CSIZE | unix.PARENB))
	t.Cflag = t.Cflag | unix.CS8
	t.Cc[unix.VMIN] = 1
	t.Cc[unix.VTIME] = 0

	return t
}
