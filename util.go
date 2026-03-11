package bosskg

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func DefaultRequestID() string {
	ts := time.Now().Unix()
	var b [2]byte
	if _, err := rand.Read(b[:]); err != nil {
		return fmt.Sprintf("%d0000", ts)
	}
	suffix := int(binary.BigEndian.Uint16(b[:])) % 10000
	return fmt.Sprintf("%d%04d", ts, suffix)
}
