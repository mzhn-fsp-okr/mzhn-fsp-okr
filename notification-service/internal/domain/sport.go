package domain

type SportTypeWithSubtypes struct {
	Id       string          `json:"id"`
	Name     string          `json:"name"`
	Subtypes []SportSubtype2 `json:"subtypes"`
}
