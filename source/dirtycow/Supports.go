package dirtycow

import "github.com/cookiengineer/groot/structs"
import "strings"

func Supports(kernel structs.Version, release structs.Release) bool {

	var result bool = false

	if kernel.Major > 2 || (kernel.Major == 2 && kernel.Minor > 6) || (kernel.Major == 2 && kernel.Minor == 6 && kernel.Patch >= 22) {

		if release.ID == "debian" {

			if release.Build == "6" {
			} else if release.Build == "7" {
				result = true
			} else if release.Build == "8" {

				if kernel.Major == 3 && kernel.Minor == 16 && kernel.Patch < 36 {
					result = true
				} else if kernel.Major == 3 && kernel.Minor <= 15 {
					result = true
				}

			}

		} else if release.ID == "ubuntu" {

			if release.Build == "12.04" {

				if kernel.Major == 3 && kernel.Minor < 2 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			} else if release.Build == "14.04" {

				if kernel.Major == 3 && kernel.Minor < 13 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			} else if release.Build == "16.04" {

				if kernel.Major == 4 && kernel.Minor < 4 {
					result = true
				} else if kernel.Major == 3 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			} else if release.Build == "16.10" {

				if kernel.Major == 4 && kernel.Minor < 8 {
					result = true
				} else if kernel.Major == 3 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			}

		} else if release.ID == "rhel" {

			if release.Build == "5" || strings.HasPrefix(release.Build, "5.") {

				if kernel.Major <= 2 && kernel.Minor <= 6 && kernel.Patch < 18 {
					result = true
				}

			} else if release.Build == "6" || strings.HasPrefix(release.Build, "6.") {

				if kernel.Major <= 2 && kernel.Minor <= 6 && kernel.Patch < 32 {
					result = true
				}

			} else if release.Build == "7" || strings.HasPrefix(release.Build, "7.") {

				if kernel.Major == 3 && kernel.Minor <= 6 && kernel.Patch < 32 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			}

		} else if release.ID == "sled" || release.ID == "sles" || release.ID == "sles_sap" {

			if release.Build == "11" || strings.HasPrefix(release.Build, "11.") {

				if kernel.Major == 3 && kernel.Minor == 0 && kernel.Patch < 101 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			} else if release.Build == "12" || strings.HasPrefix(release.Build, "12.") {

				if kernel.Major == 3 && kernel.Minor <= 12 && kernel.Patch < 60 {
					result = true
				} else if kernel.Major == 2 {
					result = true
				}

			}

		} else if kernel.Major == 2 {

			result = true

		} else if kernel.Major == 3 {

			result = true

		} else if kernel.Major == 4 {

			if kernel.Minor < 4 {
				result = true
			} else if kernel.Minor == 4 && kernel.Patch < 26 {
				result = true
			} else if kernel.Minor < 7 {
				result = true
			} else if kernel.Minor == 7 && kernel.Patch < 9 {
				result = true
			} else if kernel.Minor < 8 {
				result = true
			} else if kernel.Minor == 8 && kernel.Patch < 3 {
				result = true
			}

		}

	}

	return result

}
