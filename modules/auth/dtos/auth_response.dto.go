package dtos

type AuthResponse struct {
	Token    string `json:"token"`
	Type     string `json:"type"`
	ExpiredAt int64  `json:"expired_at"`
}