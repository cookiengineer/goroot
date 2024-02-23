package utils

import "github.com/cookiengineer/groot/structs"
import "strings"

func ParseRelease(value string) structs.Release {

	var result structs.Release

	lines := strings.Split(strings.TrimSpace(value), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if strings.Contains(line, "=") {

			key := strings.TrimSpace(line[0:strings.Index(line, "=")])
			val := strings.TrimSpace(line[strings.Index(line, "=")+1:])

			if strings.HasPrefix(val, "\"") {
				val = val[1:]
			}

			if strings.HasSuffix(val, "\"") {
				val = val[0:len(val)-1]
			}

			if key == "ID" {
				result.ID = val
			} else if key == "NAME" {
				result.Name = val
			} else if key == "BUILD_ID" {
				result.Build = val
			} else if result.Build == "" && key == "VERSION_ID" {
				result.Build = val
			}

		}

	}

	return result

}
