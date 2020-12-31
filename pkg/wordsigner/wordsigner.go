package wordsigner

import (
	"crypto/md5"
	"fmt"
)

type WordSigner struct {
	key string
}

func NewWordSigner(key string) *WordSigner {
	return &WordSigner{key: key}
}

func (w *WordSigner) Sign(word string, exists bool) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s-%v-%s", word, exists, w.key)))
	return string(hash[:])
}
