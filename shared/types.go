package shared

type UserLanguages struct {
	Name   string `json:"name"`
	Weight int    `json:"-"`
}
