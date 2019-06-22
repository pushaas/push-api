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

func (s *persistentChannelService) revivePersistentChannel(channel *models.Channel) {
	s.publicationService.PublishMessage(&models.Message{
		Channels: []string{channel.Id},
		Content: "ping",
	})
}

func (s *persistentChannelService) revivePersistentChannels() {
	channels, result := s.channelService.GetAll()
	if result != ChannelRetrievalSuccess {
		s.logger.Error("failed to retrieve persistent channels to revive")
		return
	}

	s.logger.Debug("reviving persistent channels")

	// TODO evolve to use pipeline and send all commands at once
	// TODO or even better, send a single message with all the channels in the array
	for _, channel := range channels {
		s.revivePersistentChannel(channel)
	}
}

func (s *persistentChannelService) runWorker() {
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
