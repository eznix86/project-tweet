package entity

import (
	"time"
)

// Tweet is a Tweet and a comment
type Tweet struct {
	ID         int
	Text       string     // DATA TYPE --> primitive (builtin) bool string int int8 int16 int32 (rune for unicode [for text]) int64 uint uint8 (byte) uint16 uint32 uint64 uintptr float32 float64 complex64 complex128
	Attachment Attachment // DATA TYPE ---> custom
	TweetedAt  time.Time
	Comments   []Tweet
	PostedBy   User
}
