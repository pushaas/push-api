package services

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/rafaeleyng/push-api/push-api/models"
)

type (
	PersistentChannelService interface {
		DispatchWorker()
	}

	persistentChannelService struct {
		config *viper.Viper
		logger *zap.Logger
		publicationService PublicationService
		channelService ChannelService
	}
)

func (s *persistentChannelService) revivePersistentChannels() {
	channels, channelResult := s.channelService.GetAll()
	if channelResult != ChannelRetrievalSuccess {
		s.logger.Error("failed to retrieve persistent channels to revive")
		return
	}

	s.logger.Debug("reviving persistent channels")
	channelsIds := make([]string, len(channels))

	for i, channel := range channels {
		channelsIds[i] = channel.Id
	}

	messageResult := s.publicationService.PublishMessage(&models.Message{
		Channels: channelsIds,
		Content:  "ping",
	})

	if messageResult == PublishingFailure {
		s.logger.Error("failed to revive persistent channels")
	}

	s.logger.Debug("did revive persistent channels")
}

func (s *persistentChannelService) runWorker() {
	// run once right away
	go s.revivePersistentChannels()

	// thanks https://stackoverflow.com/a/16466581/1717979
	ticker := time.NewTicker(time.Minute)
	// TODO close quit channel on app shutdown
	quit := make(chan struct{})
	for {
		select {
		case <- ticker.C:
			go s.revivePersistentChannels()
		case <- quit:
			ticker.Stop()
			return
		}
	}
}

func (s *persistentChannelService) DispatchWorker() {
	enabled := s.config.GetBool("app.persistent_channels.revive_enabled")

	if enabled {
		go s.runWorker()
	}
}

func NewPersistentChannelService(config *viper.Viper, logger *zap.Logger, publicationService PublicationService, channelService ChannelService) PersistentChannelService {
	return &persistentChannelService{
		config: config,
		logger: logger.Named("persistentChannelService"),
		publicationService: publicationService,
		channelService: channelService,
	}
}
