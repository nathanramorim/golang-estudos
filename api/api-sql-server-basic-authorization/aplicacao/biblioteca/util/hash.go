package util

import (
	"crypto/sha256"
	"fmt"
)

// GerarHashSha256 -
func GerarHashSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
