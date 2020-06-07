package modpacksch

const (
	StatusError = "error"
)

// Hardware minimum spec and recommended spec information.
// Specifically, the Minimum and Recommended fields are for
// memory in megabytes.
type Specs struct {
	ID int `json:"id"`
	Minimum int `json:"minimum"`
	Recommended int `json:"recommended"`
}
