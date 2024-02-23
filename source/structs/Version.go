package structs

import "strconv"

type Version struct {
	Major uint `json:"major"`
	Minor uint `json:"minor"`
	Patch uint `json:"patch"`
}

func (version *Version) String() string {

	var result string

	result += strconv.FormatUint(uint64(version.Major), 10)
	result += "."
	result += strconv.FormatUint(uint64(version.Minor), 10)
	result += "."
	result += strconv.FormatUint(uint64(version.Patch), 10)

	return result

}
