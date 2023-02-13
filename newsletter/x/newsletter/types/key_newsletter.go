package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewsletterKeyPrefix is the prefix to retrieve all Newsletter
	NewsletterKeyPrefix = "Newsletter/value/"
)

// NewsletterKey returns the store key to retrieve a Newsletter from the index fields
func NewsletterKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
