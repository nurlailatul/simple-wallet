package database

import (
	"simple-wallet/pkg/contextprop"
	"time"
)

const (
	timeout                        = 3 * time.Second
	Tx      contextprop.ContextKey = "tx"
	Db      contextprop.ContextKey = "db"
)
