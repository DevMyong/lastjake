package mongo

import (
	"time"
)

const (
	DBName = "littlejake"

	User collection = "users"
)

const (
	ContextTimeout = 15 * time.Second
)

type collection string

func (c collection) String() string {
	return string(c)
}
