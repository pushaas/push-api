package workers

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/services"
)

type (
	PersistentChannelsWorker interface {
		DispatchWorker()
	}

	persistentChannelsWorker struct {
		enabled bool
		logger *zap.Logger
		interval time.Duration
		persistentChannelService services.PersistentChannelService
		quitChan chan struct{}
	}
)

// thanks https://stackoverflow.com/a/16466581/1717979
func (s *persistentChannelsWorker) startWorker() {
	// run once right away
	go s.persistentChannelService.TriggerRevivePersistentChannels()

	ticker := time.NewTicker(s.interval)
	s.quitChan = make(chan struct{})
	for {
		select {
		case <- ticker.C:
			go s.persistentChannelService.TriggerRevivePersistentChannels()
		case <- s.quitChan:
			s.quitChan = nil
			ticker.Stop()
			return
		}
	}
}

func (s *persistentChannelsWorker) stopWorker() {
	if s.quitChan != nil {
		s.quitChan <- struct{}{}
	}
}

func (s *persistentChannelsWorker) DispatchWorker() {
	if s.enabled {
		go s.startWorker()
	}
}

func NewPersistentChannelsWorker(lc fx.Lifecycle, config *viper.Viper, logger *zap.Logger, persistentChannelService services.PersistentChannelService) PersistentChannelsWorker {
	enabled := config.GetBool("workers.persistent_channels.enabled")
	interval := config.GetDuration("workers.persistent_channels.interval")

	s := persistentChannelsWorker{
		enabled: enabled,
		logger: logger.Named("persistentChannelsWorker"),
		interval: interval,
		persistentChannelService: persistentChannelService,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			s.stopWorker()
			return nil
		},
	})
	return &s
}
