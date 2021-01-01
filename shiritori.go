package shiritoriapi

import (
	"context"
	"log"
	"shiritori/gen/shiritori"
	"shiritori/pkg/wordchecker"
	"shiritori/pkg/wordsigner"
	"shiritori/repository/memory"
	"shiritori/values"
)

type shiritorisrvc struct {
	logger      *log.Logger
	wordChecker WordChecker
	wordSigner  WordSigner
	repo        *RepositoryFactory
}

type WordChecker interface {
	Check(ctx context.Context, word string) (bool, error)
}

type WordSigner interface {
	Sign(word string, exists bool) string
}

type RepositoryFactory struct {
	BattleEvent BattleEventRepository
}

type BattleEventRepository interface {
	Insert(event values.BattleEvent) error
	GetNewer(battleId values.BattleID, timestamp values.BattleEventTimestamp) ([]values.BattleEvent, error)
}

// NewShiritori returns the shiritori service implementation.
func NewShiritori(logger *log.Logger) shiritori.Service {
	return &shiritorisrvc{
		logger:      logger,
		wordChecker: wordchecker.NewWordChecker(),
		wordSigner:  wordsigner.NewWordSigner("123456789"),
		repo:        NewMemoryRepositoryFactory(),
	}
}

func NewMemoryRepositoryFactory() *RepositoryFactory {
	return &RepositoryFactory{
		BattleEvent: memory.NewBattleEventRepository(),
	}
}

// Add implements add.
func (s *shiritorisrvc) Add(ctx context.Context, p *shiritori.AddPayload) (res int, err error) {
	s.logger.Print("shiritori.add")
	return
}
