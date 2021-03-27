package modpacksch

// Hardware minimum spec and recommended spec information.
// Specifically, the Minimum and Recommended fields are for
// memory in mebibytes.
type Specs struct {
	ID          int `json:"id"`
	Minimum     int `json:"minimum"`
	Recommended int `json:"recommended"`
}

type Art struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Compressed bool   `json:"compressed"`
	Sha1       string `json:"sha1"`
	Size       int    `json:"size"`
	Updated    int64  `json:"updated"`
}

type Author struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Website string `json:"website"`
	Updated int64  `json:"updated"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Link struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	Type string `json:"type"`
}
