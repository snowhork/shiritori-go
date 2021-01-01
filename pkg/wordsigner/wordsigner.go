package wordsigner

import (
	"crypto/md5"
	"fmt"
	"shiritori/values"
)

type WordSigner struct {
	key string
}

func NewWordSigner(key string) *WordSigner {
	return &WordSigner{key: key}
}

func (w *WordSigner) Sign(word values.WordBody, exists bool) values.WordBodyHash {
	hash := md5.Sum([]byte(fmt.Sprintf("%s-%v-%s", word, exists, w.key)))
	return values.WordBodyHash(hash[:])
}
