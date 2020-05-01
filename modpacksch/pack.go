package modpacksch

type Pack struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Synopsis string `json:"synopsis"`
	Description string `json:"description"`
	Type string `json:"type"`
	Featured bool `json:"featured"`
	Installs int `json:"installs"`
	Plays int `json:"plays"`
	Updated int64 `json:"updated"`
	Refreshed int64 `json:"refreshed"`

	Art []*Art `json:"art"`
	Authors []*Author `json:"authors"`
	Versions []*VersionInfo `json:"versions"`
	Tags []*Tag `json:"tags"`
}

func (p *Pack) HasIcon() bool {
	for _, art := range p.Art {
		if art.Type == "square" {
			return true
		}
	}
	return false
}

func (p *Pack) GetIcon() *Art {
	for _, art := range p.Art {
		if art.Type == "square" {
			return art
		}
	}
	return nil
}

type Art struct {
	ID int `json:"id"`
	URL string `json:"url"`
	Type string `json:"type"`
	Width int `json:"width"`
	Height int `json:"height"`
	Compressed bool `json:"compressed"`
	Sha1 string `json:"sha1"`
	Size int `json:"size"`
	Updated int64 `json:"updated"`
}

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Website string `json:"website"`
	Updated int64 `json:"updated"`
}

type VersionInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Updated int64 `json:"updated"`

	Specs *Specs `json:"specs"`
}

type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
