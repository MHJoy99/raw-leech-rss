package task

import (
	"github.com/reaitten/flowerss-bot/bot"
	"github.com/reaitten/flowerss-bot/config"
	"github.com/reaitten/flowerss-bot/model"
	"time"
)

func Update() {
	if config.RunMode == config.TestMode {
		return
	}

	for {
		sources := model.GetSubscribedNormalSources()
		for _, source := range sources {
			if !source.NeedUpdate() {
				continue
			}
			c, err := source.GetNewContents()
			if err == nil && len(c) > 0 {
				subs := model.GetSubscriberBySource(&source)
				bot.BroadcastNews(&source, subs, c)
			}
			if source.ErrorCount >= config.ErrorThreshold {
				bot.BroadcastSourceError(&source)
			}
		}
		time.Sleep(time.Duration(config.UpdateInterval) * time.Minute)
	}
}
