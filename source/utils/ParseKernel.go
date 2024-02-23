package utils

import "github.com/cookiengineer/groot/structs"
import "strconv"
import "strings"

func ParseKernel(value string) structs.Version {

	var result structs.Version

	if strings.Contains(value, "-") {
		value = value[0:strings.Index(value, "-")]
	}

	if strings.Contains(value, ".") {

		tmp := strings.Split(value, ".")

		if len(tmp) >= 3 {

			num1, err1 := strconv.ParseUint(tmp[0], 10, 64)
			num2, err2 := strconv.ParseUint(tmp[1], 10, 64)
			num3, err3 := strconv.ParseUint(tmp[2], 10, 64)

			if err1 == nil && err2 == nil && err3 == nil {
				result.Major = uint(num1)
				result.Minor = uint(num2)
				result.Patch = uint(num3)
			}

		} else if len(tmp) >= 2 {

			num1, err1 := strconv.ParseUint(tmp[0], 10, 64)
			num2, err2 := strconv.ParseUint(tmp[1], 10, 64)

			if err1 == nil && err2 == nil {
				result.Major = uint(num1)
				result.Minor = uint(num2)
			}

		} else if len(tmp) >= 1 {

			num1, err1 := strconv.ParseUint(tmp[0], 10, 64)

			if err1 == nil {
				result.Major = uint(num1)
			}

		}

	}

	return result

}
