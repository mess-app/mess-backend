package models

type WsEvent struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}
