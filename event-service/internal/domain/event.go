package domain

type EventLoadInfo struct {
	EkpId        string
	SportType    string
	SportSubtype string
	Name         string
	Description  string
	Dates        DateRange
	Location     string
	Participants int
}

type EventInfo struct {
	Id           string
	EkpId        string
	SportType    string
	SportSubtype string
	Name         string
	Description  string
	Dates        DateRange
	Location     string
	Participants int
}

type DateRange struct {
	From string
	To   string
}
