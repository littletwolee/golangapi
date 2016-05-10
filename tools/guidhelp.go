package tools

import (
    "crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "io"
)

// @Title Generate md5
// @Description Generate 32-bit string md5
// @Success string
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

// @Title Generate guid
// @Description Generate guid
// @Success string
func GetGuid() string {
    b := make([]byte, 48)

    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
