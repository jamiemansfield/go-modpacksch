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
	request, err := s.client.NewRequest(http.MethodGet, "public/curseforge/" + strconv.Itoa(id), nil)
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

func (s *CurseForgeService) GetVersion(packId int, versionId int) (*Version, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/curseforge/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId), nil)
	if err != nil {
		return nil, err
	}

	var response Version
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
