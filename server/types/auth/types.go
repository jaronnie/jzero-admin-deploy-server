// Code generated by jzero. DO NOT EDIT.
package auth

import (
	"time"
)

var (
	_ = time.Now()
)

type ErrorRequest struct {
}

type ErrorResponse struct {
}

type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
	UserId   string   `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Buttons  []string `json:"buttons"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
}

type RefreshTokenResponse struct {
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
}
