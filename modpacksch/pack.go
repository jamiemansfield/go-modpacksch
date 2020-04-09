package modpacksch

import "strconv"

// PackService handles communication with the pack related
// methods of the modpacks.ch API.
//
// modpacks.ch API docs: https://modpacksch.docs.apiary.io/#/reference/0/modpacks
type PackService service

func (s *PackService) GetPack(id int) (*Pack, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/" + strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	var response Pack
	_, err = s.client.Do(request, &response)
	return &response, err
}

func (s *PackService) GetVersion(packId int, versionId int) (*Version, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId), nil)
	if err != nil {
		return nil, err
	}

	var response Version
	_, err = s.client.Do(request, &response)
	return &response, err
}

func (s *PackService) GetVersionChangelog(packId int, versionId int) (*VersionChangelog, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId) + "/changelog", nil)
	if err != nil {
		return nil, err
	}

	var response VersionChangelog
	_, err = s.client.Do(request, &response)
	return &response, err
}

func (s *PackService) Search(limit int, term string) ([]int, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/search/" + strconv.Itoa(limit) + "?term=" + term, nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

type searchResponse struct {
	Packs []int `json:"packs"`
}

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
