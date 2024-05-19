package models

type WsEventSendMessageData struct {
	Message      string `mapstructure:"message"`
	CreatedAt    string `mapstructure:"created_at"`
	ConnectionId string `mapstructure:"connection_id"`
	RecipientId  string `mapstructure:"recipient_id"`
}
