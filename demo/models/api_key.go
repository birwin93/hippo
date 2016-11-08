package models

type ApiKey struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}
