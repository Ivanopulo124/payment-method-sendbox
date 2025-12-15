package auth

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement"`
	Login     string    `bun:",notnull"`
	Password  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
