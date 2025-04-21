package token

import "time"

type Maker interface {
	// JWT
	CreateToken(id int, permissions string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
