package structs

type Release struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Build string `json:"build"`
}

func (release *Release) String() string {

	var result string

	result += release.ID

	if release.Build != "" {
		result += " (" + release.Build + ")"
	}

	return result

}
