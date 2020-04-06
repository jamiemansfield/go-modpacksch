package modpacksch

import "strconv"

type PackService struct {
	client *Client
}

func (s *PackService) GetPack(id int) *Pack {
	request, err := s.client.newRequest("GET", "/public/modpack/" + strconv.Itoa(id), nil)
	if err != nil {
		return nil
	}

	var response Pack
	_, err = s.client.do(request, &response)
	return &response
}

func (s *PackService) Search(limit int, term string) []int {
	request, err := s.client.newRequest("GET", "/public/modpack/search/" + strconv.Itoa(limit) + "?term=" + term, nil)
	if err != nil {
		return nil
	}

	var response searchResponse
	_, err = s.client.do(request, &response)
	return response.Packs
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
	Updated int64 `json:"updated"`

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
}

type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
