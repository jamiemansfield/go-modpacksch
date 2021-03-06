package modpacksch

import (
	"net/http"
	"strconv"
)

// CurseForgeService handles communication with the CurseForge
// related methods of the modpacks.ch API.
//
// modpacks.ch API docs: https://modpacksch.docs.apiary.io/#/reference/0/curseforge
type CurseForgeService service

func (s *CurseForgeService) GetPack(id int) (*Pack, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/curseforge/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	var response Pack
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *CurseForgeService) GetVersion(packId int, versionId int) (*PackVersion, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/curseforge/"+strconv.Itoa(packId)+"/"+strconv.Itoa(versionId), nil)
	if err != nil {
		return nil, err
	}

	var response PackVersion
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *CurseForgeService) Search(limit int, term string) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/search/"+strconv.Itoa(limit)+"?term="+term, nil)
	if err != nil {
		return nil, err
	}

	var response curseForgeSearchResponse
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return response.CurseForge, nil
}

type curseForgeSearchResponse struct {
	CurseForge []int `json:"curseforge"`
}
