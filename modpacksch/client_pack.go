package modpacksch

import (
	"errors"
	"net/http"
	"strconv"
)

// PackService handles communication with the pack related
// methods of the modpacks.ch API.
//
// modpacks.ch API docs: https://modpacksch.docs.apiary.io/#/reference/0/modpacks
type PackService service

func (s *PackService) GetPack(id int) (*Pack, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/" + strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	var response Pack
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == StatusError {
		err = errors.New("client: " + response.Message)
	}
	return &response, err
}

func (s *PackService) GetVersion(packId int, versionId int) (*Version, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId), nil)
	if err != nil {
		return nil, err
	}

	var response Version
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == StatusError {
		err = errors.New("client: " + response.Message)
	}
	return &response, err
}

func (s *PackService) GetVersionChangelog(packId int, versionId int) (*VersionChangelog, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId) + "/changelog", nil)
	if err != nil {
		return nil, err
	}

	var response VersionChangelog
	_, err = s.client.Do(request, &response)
	if err != nil {
		return nil, err
	}

	if response.Status == StatusError {
		err = errors.New("client: " + response.Content)
	}
	return &response, err
}

func (s *PackService) IncrementPlayCount(packId int, versionId int) error {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId) + "/play", nil)
	if err != nil {
		return err
	}

	var response statsResponse
	_, err = s.client.Do(request, &response)
	return err
}

func (s *PackService) IncrementInstallCount(packId int, versionId int) error {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/" + strconv.Itoa(packId) + "/" + strconv.Itoa(versionId) + "/install", nil)
	if err != nil {
		return err
	}

	var response statsResponse
	_, err = s.client.Do(request, &response)
	return err
}

func (s *PackService) All() ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/all/", nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) Search(limit int, term string) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/search/" + strconv.Itoa(limit) + "?term=" + term, nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostPlayed(limit int) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/popular/plays/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostPlayedWithTag(limit int, tag string) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/popular/plays/" + tag + "/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostInstalled(limit int) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/popular/installs/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostInstalledWithTag(limit int, tag string) ([]int, error) {
	request, err := s.client.NewRequest(http.MethodGet, "public/modpack/popular/installs/" + tag + "/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

type statsResponse struct {
}

type searchResponse struct {
	Packs []int `json:"packs"`
}
