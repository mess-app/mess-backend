package models

type SupabaseConnection struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	PioneerId   string `json:"pioneer_id"`
	RecipientId string `json:"recipient_id"`
	Status      string `json:"status"`
	Pioneer     struct {
		UserId string `json:"user_id"`
	} `json:"pioneer"`
	Recipient struct {
		UserId string `json:"user_id"`
	} `json:"recipient"`
}
