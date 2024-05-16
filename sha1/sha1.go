package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(content []byte) string {
	h := sha1.New()
	h.Write(content)
	hashedContent := hex.EncodeToString(h.Sum(nil))
	return hashedContent
}
