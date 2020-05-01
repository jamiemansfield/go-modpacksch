package modpacksch

import "strconv"

// TagService handles communication with the tags related
// methods of the modpacks.ch API.
//
// modpacks.ch API docs: https://modpacksch.docs.apiary.io/#/reference/0/tags
type TagService service

func (s *PackService) MostUsed(limit int) ([]string, error) {
	request, err := s.client.NewRequest("GET", "public/tag/popular/" + strconv.Itoa(limit), nil)
	if err != nil {
		return nil, err
	}

	var response tagResponse
	_, err = s.client.Do(request, &response)
	return response.Tags, err
}

type tagResponse struct {
	Tags []string `json:"tags"`
}
