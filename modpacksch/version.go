package modpacksch

type PackVersion struct {
	ID        int    `json:"id"`
	Parent    int    `json:"parent"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Installs  int    `json:"installs"`
	Plays     int    `json:"plays"`
	Updated   int64  `json:"updated"`
	Refreshed int64  `json:"refreshed"`

	Specs   *Specs    `json:"specs"`
	Targets []*Target `json:"targets"`
	Files   []*File   `json:"files"`

	// Error Handling
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Target struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Updated int64  `json:"updated"`
}

type File struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	URL        string `json:"url"`
	Sha1       string `json:"sha1"`
	Size       int    `json:"size"`
	ClientOnly bool   `json:"clientonly"`
	ServerOnly bool   `json:"serveronly"`
	Optional   bool   `json:"optional"`
	Updated    int64  `json:"updated"`
}

type VersionChangelog struct {
	Content string `json:"content"`
	Updated int64  `json:"updated"`

	// Error Handling
	Status string `json:"status"`
}
