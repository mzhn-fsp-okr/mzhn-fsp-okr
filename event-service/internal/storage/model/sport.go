package model

type SportType struct {
	Id   string
	Name string
}

type Subtype struct {
	Id   string
	Name string
}

type SportTypeWithSubtypes struct {
	SportType
	Subtypes []Subtype
}
