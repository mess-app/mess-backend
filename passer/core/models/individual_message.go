package models

type WsEventSendMessageData struct {
	Message      string `mapstructure:"message"`
	ConnectionId string `mapstructure:"connection_id"`
}
