package modpacksch

type Mod struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Synopsis string `json:"synopsis"`
	Description string `json:"description"`
	Type string `json:"type"`
	Installs int `json:"installs"`

	Art []*Art `json:"art"`
	Authors []*Author `json:"authors"`
	Versions []*ModVersion `json:"versions"`

	Updated int64 `json:"updated"`
	Refreshed int64 `json:"refreshed"`
}

type ModVersion struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	Path string `json:"path"`
	Version string `json:"version"`
	URL string `json:"url"`
	Sha1 string `json:"sha1"`
	Size int `json:"size"`
	ClientOnly bool `json:"clientonly"`

	Targets []*Target `json:"targets"`
	Dependencies []int `json:"dependencies"`

	Updated int64 `json:"updated"`
}
