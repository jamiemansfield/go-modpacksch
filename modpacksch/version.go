package modpacksch

type Version struct {
	ID int `json:"id"`
	Parent int `json:"parent"`
	Name string `json:"name"`
	Type string `json:"type"`
	Installs int `json:"installs"`
	Plays int `json:"plays"`
	Updated int64 `json:"updated"`
	Refreshed int64 `json:"refreshed"`

	Specs *Specs `json:"specs"`
}

type VersionChangelog struct {
	Content string `json:"content"`
	Updated int64 `json:"updated"`
}
