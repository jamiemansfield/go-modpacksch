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

func (s *PackService) MostPlayed(limit int) ([]int, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/popular/plays/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostPlayedWithTag(limit int, tag string) ([]int, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/popular/plays/" + tag + "/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostInstalled(limit int) ([]int, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/popular/installs/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response searchResponse
	_, err = s.client.Do(request, &response)
	return response.Packs, err
}

func (s *PackService) MostInstalledWithTag(limit int, tag string) ([]int, error) {
	request, err := s.client.NewRequest("GET", "public/modpack/popular/installs/" + tag + "/" + strconv.Itoa(limit), nil)
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
