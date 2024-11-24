package model

type SportType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Subtype struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SportTypeWithSubtypes struct {
	SportType
	Subtypes []Subtype `json:"subtypes"`
}

type SportFilter struct {
	Name *string
}
