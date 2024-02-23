package main

import "github.com/cookiengineer/groot/dirtycow"
import "github.com/cookiengineer/groot/utils"
import "os"
import "os/exec"
import "fmt"

func main() {

	cmd1 := exec.Command("uname", "-r")
	buf1, err1 := cmd1.Output()

	buf2, err2 := os.ReadFile("/etc/os-release")

	exploits := make([]string, 0)

	if err1 == nil && err2 == nil {

		kernel := utils.ParseKernel(string(buf1))
		distro := utils.ParseRelease(string(buf2))

		if dirtycow.Supports(kernel, distro) {
			exploits = append(exploits, "dirtycow")
		}

		fmt.Println("Detected \"" + distro.String() + "\" with kernel \"" + kernel.String() + "\"")

		if len(exploits) > 0 {

			fmt.Println(exploits)

		} else {

			fmt.Println("Your operating system is safe, nice work!")
			os.Exit(0)

		}

	}

}
