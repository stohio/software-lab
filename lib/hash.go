package softwarelab

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// HashFileMd5 takes in a path to a file, calculates the md5 hash of that file,
// and returns the hex encoding as a string
func HashFileMd5(file io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	hash := h.Sum(nil)

	return hex.EncodeToString(hash), nil
}
