package wordchecker

import "context"

type WordChecker struct{}

func NewWordChecker() *WordChecker {
	return &WordChecker{}

}

func (w *WordChecker) Check(ctx context.Context, word string) (bool, error) {
	return true, nil
}
