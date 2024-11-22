package domain

type ParticipantRequirements struct {
	Gender bool   `json:"gender"`
	MinAge *int32 `json:"minAge"`
	MaxAge *int32 `json:"maxAge"`
}

type DateRange struct {
	From string `json: "from"`
	To   string `json: "to"`
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
	Id                      string                    `json:"id"`
	EkpId                   string                    `json:"ekpId"`
	SportType               string                    `json:"sportType"`
	SportSubtype            string                    `json:"sportSubtype"`
	Name                    string                    `json:"name"`
	Description             string                    `json:"description"`
	Dates                   DateRange                 `json:"dates"`
	Location                string                    `json:"location"`
	Participants            int                       `json:"participants"`
	ParticipantRequirements []ParticipantRequirements `json:"participantRequirements"`
}
