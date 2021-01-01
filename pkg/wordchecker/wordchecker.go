package wordchecker

import (
	"context"
	"shiritori/values"
)

type WordChecker struct{}

func NewWordChecker() *WordChecker {
	return &WordChecker{}
}

func (w *WordChecker) Check(ctx context.Context, word values.WordBody) (bool, error) {
	return true, nil
}
