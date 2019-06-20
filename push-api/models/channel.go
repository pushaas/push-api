package models

import (
	"time"
)

type (
	Channel struct {
		Id string `json:"id"`
		Ttl time.Duration `json:"ttl"`
	}
)
