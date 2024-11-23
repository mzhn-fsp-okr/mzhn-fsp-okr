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

type LinkTelegramRequest struct {
	UserId           string
	Token            string
	TelegramUsername string
}
