package modpacksch

type Pack struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Synopsis     string `json:"synopsis"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Featured     bool   `json:"featured"`
	Installs     int    `json:"installs"`
	Plays        int    `json:"plays"`
	Updated      int64  `json:"updated"`
	Refreshed    int64  `json:"refreshed"`
	Notification string `json:"notification"`

	Art      []*Art             `json:"art"`
	Links    []*Link            `json:"links"`
	Authors  []*Author          `json:"authors"`
	Versions []*PackVersionInfo `json:"versions"`
	Tags     []*Tag             `json:"tags"`
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

type PackVersionInfo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Updated int64  `json:"updated"`

	Specs *Specs `json:"specs"`
}
