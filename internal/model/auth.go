// model/auth.go
package model

// AuthResponse содержит токены, которые будут возвращаться при успешной аутентификации.
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
