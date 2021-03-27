package modpacksch

import (
	"net/http"
	"strconv"
)

// ModService handles communication with the mods related methods of
// the modpacks.ch API.
//
// modpacks.ch API docs: https://modpacksch.docs.apiary.io/#/reference/0/mods
type ModService service

func (s *ModService) Get(id int) (*Mod, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/mod/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	var response Mod
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *ModService) Search(limit int, term string) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/mod/search/"+strconv.Itoa(limit)+"?term="+term, nil)
	if err != nil {
		return nil, err
	}

	var response modSearchResponse
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return response.Mods, nil
}

type modSearchResponse struct {
	Mods []int `json:"packs"`
}
