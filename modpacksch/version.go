package modpacksch

type Version struct {
	ID int
	Parent int
	Name string
	Type string
	Updated int64
	Refreshed int64
}

type VersionChangelog struct {
	Content string `json:"content"`
	Updated int64 `json:"updated"`
}
