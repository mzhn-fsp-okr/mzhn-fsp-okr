package domain

type ParticipantRequirements struct {
	Gender bool
	MinAge *int32
	MaxAge *int32
}

type DateRange struct {
	From string
	To   string
}

type EventLoadInfo struct {
	EkpId        string
	SportType    string
	SportSubtype string
	Name         string
	Description  string
	Dates        DateRange
	Location     string
	Participants int

	ParticipantRequirements []ParticipantRequirements
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
