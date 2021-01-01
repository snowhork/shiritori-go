package shiritoriapi

import (
	"context"
	"shiritori/values"
	"sort"
	"time"

	"github.com/pkg/errors"
)

type BattleEventSubscriber struct {
	ticker   <-chan time.Time
	lastTime int64
	battleID string
	repo     BattleEventRepository
	callback BattleEventSubscribe
}

type BattleEventSubscribe func(event values.BattleEvent) error

func NewBattleEventSubscriber(ticker <-chan time.Time, currentTime int64, battleID string, repo BattleEventRepository, callback BattleEventSubscribe) *BattleEventSubscriber {
	return &BattleEventSubscriber{ticker: ticker, lastTime: currentTime, battleID: battleID, repo: repo, callback: callback}
}

func (sub *BattleEventSubscriber) Subscribe(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-sub.ticker:
			events, err := sub.repo.GetNewer(sub.battleID, sub.lastTime)
			if err != nil {
				return errors.Wrap(err, "Message Get Error")
			}

			sort.Slice(events, func(i, j int) bool {
				return events[i].Timestamp < events[j].Timestamp
			})

			for _, event := range events {
				if sub.lastTime < event.Timestamp {
					if err := sub.callback(event); err != nil {
						return errors.Wrap(err, "Message Send Error")
					}

					sub.lastTime = event.Timestamp
				}
			}
		}
	}
}
