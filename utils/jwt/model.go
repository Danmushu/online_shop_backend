package jwt

import "project/utils/perms"

type AuthToken struct {
	ID       int    `json:"id"`
	Usr      string `json:"usr"`
	IssuedAt int64  `json:"iat"`
	ExpireAt int64  `json:"exp"`
}

type OPAuthToken struct {
	ID   int        `json:"id"`
	Usr  string     `json:"usr"`
	Perm perms.Perm `json:"perm,number"`
	Key  string     `json:"key"`

	IssuedAt int64 `json:"iat"`
	ExpireAt int64 `json:"exp"`
}
