package domain

type Integrations struct {
	UserId           string
	TelegramUsername *string
	WannaMail        bool
}

type SetIntegrations struct {
	UserId           string
	TelegramUsername *string
	WannaMail        *bool
}
