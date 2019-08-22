package workers

import (
	"context"
	"time"

	"github.com/bsm/redis-lock"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/pushaas/push-api/push-api/services"
)

type (
	PersistentChannelsWorker interface {
		DispatchWorker()
	}

	persistentChannelsWorker struct {
		enabled bool
		interval time.Duration
		lockKey string
		lockTimeout time.Duration
		logger *zap.Logger
		persistentChannelService services.PersistentChannelService
		quitChan chan struct{}
		redisClient redis.UniversalClient
		workersEnabled bool
	}
)

func (w *persistentChannelsWorker) performAction() {
	options := &lock.Options{
		LockTimeout: w.lockTimeout,
	}
	_, err := lock.Obtain(w.redisClient, w.lockKey, options)

	if err == lock.ErrLockNotObtained {
		w.logger.Debug("could not obtain lock")
		return
	}

	if err != nil {
		w.logger.Error("error while trying to obtain lock", zap.Error(err))
		return
	}

	go w.persistentChannelService.TriggerRevivePersistentChannels()
}

// thanks https://stackoverflow.com/a/16466581/1717979
func (w *persistentChannelsWorker) startWorker() {
	w.performAction() // run once right away

	ticker := time.NewTicker(w.interval)
	for {
		select {
		case <- ticker.C:
			w.performAction()
		case <- w.quitChan:
			w.quitChan = nil
			ticker.Stop()
			return
		}
	}
}

func (w *persistentChannelsWorker) stopWorker() {
	if w.quitChan != nil {
		w.quitChan <- struct{}{}
	}
}

func (w *persistentChannelsWorker) DispatchWorker() {
	if w.workersEnabled && w.enabled {
		go w.startWorker()
	}
}

func NewPersistentChannelsWorker(lc fx.Lifecycle, config *viper.Viper, logger *zap.Logger, redisClient redis.UniversalClient, persistentChannelService services.PersistentChannelService) PersistentChannelsWorker {
	enabled := config.GetBool("workers.persistent_channels.enabled")
	interval := config.GetDuration("workers.persistent_channels.interval")
	lockKey := config.GetString("workers.persistent_channels.lock_key")
	lockTimeout := config.GetDuration("workers.persistent_channels.lock_timeout")
	workersEnabled := config.GetBool("workers.enabled")

	s := persistentChannelsWorker{
		enabled: enabled,
		logger: logger.Named("persistentChannelsWorker"),
		interval: interval,
		lockKey: lockKey,
		lockTimeout: lockTimeout,
		persistentChannelService: persistentChannelService,
		quitChan: make(chan struct{}),
		redisClient: redisClient,
		workersEnabled: workersEnabled,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			s.stopWorker()
			return nil
		},
	})

	return &s
}
