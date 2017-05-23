package model

type Post struct {
	ID        int    `json:"id,omitempty" db:"id,omitempty"`
	ReceiveID int    `json:"receive_id" db:"receive_id"`
	SendID    int    `json:"send_id" db:"send_id"`
	Message   string `json:"message" db:"message"`
}
