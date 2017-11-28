package goTool

import (
	"crypto/md5"
	"fmt"
)

// IsTool returns whether or not the name is a tool.  Deal with it.
func IsTool(b []byte) bool {

	// md5.Sum returns an array of 16 bytes...convert it to a base16 string.
	md5Sum := fmt.Sprintf("%x", md5.Sum(b))

	// if the last character in the base16 string is int value 8 - F, then return true.
	if md5Sum[len(md5Sum)-1] > '8' {
		return true
	}

	return false
}
