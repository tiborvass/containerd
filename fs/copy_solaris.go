// +build solaris

package fs

import (
	"io"
	"os"
	"syscall"

	"github.com/pkg/errors"
)

func copyFileInfo(fi os.FileInfo, name string) error {
	st := fi.Sys().(*syscall.Stat_t)
	if err := os.Lchown(name, int(st.Uid), int(st.Gid)); err != nil {
		return errors.Wrapf(err, "failed to chown %s", name)
	}

	if (fi.Mode() & os.ModeSymlink) != os.ModeSymlink {
		if err := os.Chmod(name, fi.Mode()); err != nil {
			return errors.Wrapf(err, "failed to chmod %s", name)
		}
	}

	/*
	if err := syscall.UtimesNano(name, []syscall.Timespec{st.Atimespec, st.Mtimespec}); err != nil {
		return errors.Wrapf(err, "failed to utime %s", name)
	}
	*/

	return nil
}

func copyFileContent(dst, src *os.File) error {
	buf := bufferPool.Get().([]byte)
	_, err := io.CopyBuffer(dst, src, buf)
	bufferPool.Put(buf)

	return err
}

func copyXAttrs(dst, src string) error {
	return nil
}

func copyDevice(dst string, fi os.FileInfo) error {
	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return errors.New("unsupported stat type")
	}
	return syscall.Mknod(dst, uint32(fi.Mode()), int(st.Rdev))
}
